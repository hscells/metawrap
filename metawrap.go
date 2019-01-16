package metawrap

import (
	"os/exec"
	"fmt"
	"bufio"
	"bytes"
	"encoding/json"
)

type MappingCandidate struct {
	CandidateScore     string
	CandidateCUI       string
	CandidateMatched   string
	CandidatePreferred string
	SemTypes           []string
	Sources            []string
	MatchedWords       []string
}

type Mapping struct {
	MappingScore      string
	MappingCandidates []MappingCandidate
}

type Phrase struct {
	PhraseText string
	Mappings   []Mapping
}

type Utterance struct {
	UttText string
	Phrases []Phrase
}

type Document struct {
	Utterances []Utterance
}

type X struct {
	Document Document
}

type MetaMapping struct {
	AllDocuments []X
}

type MetaMap struct {
	path string
}

// NewMetaMapClient creates a new metawrap client.
func NewMetaMapClient(path string) MetaMap {
	return MetaMap{path: path}
}

// Map uses MetaMap to map text to candidates.
func (m MetaMap) Map(text string) (MetaMapping, error) {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("echo '%v' | %v --JSONn --silent", text, m.path))

	r, err := cmd.StdoutPipe()
	if err != nil {
		return MetaMapping{}, err
	}

	cmd.Start()

	s := bufio.NewScanner(bufio.NewReader(r))
	// Skip the first line.
	s.Scan()
	var buff bytes.Buffer
	for s.Scan() {
		_, err = buff.Write(s.Bytes())
		if err != nil {
			return MetaMapping{}, err
		}
	}

	var d MetaMapping
	if err := json.NewDecoder(bytes.NewReader(buff.Bytes())).Decode(&d); err != nil {
		return MetaMapping{}, err
	}

	if err := cmd.Wait(); err != nil {
		return MetaMapping{}, err
	}

	return d, nil
}

// PreferredCandidates extracts only the MetaMap preferred candidates.
func (m MetaMap) PreferredCandidates(text string) ([]MappingCandidate, error) {
	mapping, err := m.Map(text)
	if err != nil {
		return nil, err
	}

	var candidates []MappingCandidate
	for _, d := range mapping.AllDocuments {
		for _, u := range d.Document.Utterances {
			for _, p := range u.Phrases {
				for _, m := range p.Mappings {
					for _, c := range m.MappingCandidates {
						candidates = append(candidates, c)
					}
				}
			}
		}
	}
	return candidates, nil
}
