# GBFT

Contributed by:
Changsheng Liu (UCSC)
Shuli He (UCSC)

GBFT is inprovement algorithm of PBFT
modify from [PBFT](https://github.com/bigpicturelabs/consensusPBFT)
Algorithm designed from [Faisal Nawab](https://github.com/faisalsyn)

## How To Run

run `make`, then you will have two executable files: "main" and "fakemain". The data center nodes are named from "DC1-1" to "DC1-4", from "DC2-1" to "DC2-4" and from "DC3-1" to "DC3-4". You can either run `./main <nodename>` or `./fake <nodename>` to simulate a normal node or a byzantine failure node.