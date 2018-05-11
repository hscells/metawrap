# metawrap

metawrap is a wrapper for MetaMap. metawrap can either be used as a library, or can be run as a server to be consumed
via a REST api.

To install the server:

```
go install -u github.com/hscells/metawrap
```

To start the server, issue:

`metawrap --path /path/to/metamap/binary`

Text can be mapped fully or one can just get the matched candidates:

`curl -X GET localhost:4646/mm/map --data "foo..."`

`curl -X GET localhost:4646/mm/candidates --data "bar..."`

Note that the entire mapping is not supported, however an example response looks like:

```json
{
  "AllDocuments": [
    {
      "Document": {
        "Utterances": [
          {
            "UttText": "A commercially available galactomannan sandwich ELISA (Platelia) was the test under evaluation.",
            "Phrases": [
              {
                "PhraseText": "A commercially available galactomannan sandwich ELISA",
                "Mappings": [
                  {
                    "MappingScore": "-740",
                    "MappingCandidates": [
                      {
                        "CandidateScore": "-637",
                        "CandidateCUI": "C0470187",
                        "CandidateMatched": "Available",
                        "CandidatePreferred": "Availability of",
                        "SemTypes": [
                          "ftcn"
                        ],
                        "Sources": [
                          "AOD",
                          "CHV",
                          "NCI",
                          "SNMI",
                          "SNOMEDCT_US"
                        ]
                      },
                      {
                        "CandidateScore": "-637",
                        "CandidateCUI": "C0060961",
                        "CandidateMatched": "Galactomannan",
                        "CandidatePreferred": "galactomannan",
                        "SemTypes": [
                          "orch"
                        ],
                        "Sources": [
                          "CHV",
                          "LNC",
                          "MSH"
                        ]
                      },
                      {
                        "CandidateScore": "-804",
                        "CandidateCUI": "C0014441",
                        "CandidateMatched": "ELISA",
                        "CandidatePreferred": "Enzyme-Linked Immunosorbent Assay",
                        "SemTypes": [
                          "lbpr"
                        ],
                        "Sources": [
                          "AOD",
                          "CHV",
                          "CSP",
                          "HL7V3.0",
                          "LCH",
                          "LCH_NW",
                          "MSH",
                          "NCI",
                          "NCI_CDISC",
                          "NCI_NCI-GLOSS",
                          "NLMSubSyn",
                          "SNMI",
                          "SNOMEDCT_US"
                        ]
                      }
                    ]
                  }
                ]
              },
              {
                "PhraseText": "(Platelia",
              },
              {
                "PhraseText": ")",
              },
              {
                "PhraseText": "was",
              },
              {
                "PhraseText": "the test under evaluation.",
                "Mappings": [
                  {
                    "MappingScore": "-819",
                    "MappingCandidates": [
                      {
                        "CandidateScore": "-819",
                        "CandidateCUI": "C0178628",
                        "CandidateMatched": "test evaluation",
                        "CandidatePreferred": "evaluation/testing",
                        "SemTypes": [
                          "resa"
                        ],
                        "Sources": [
                          "CHV",
                          "CSP",
                          "NLMSubSyn"
                        ]
                      }
                    ]
                  }
                ]
              }
            ]
          }
        ]
      }
    }
  ]
}
```

This can be extended by modifying `metawrap.go`.