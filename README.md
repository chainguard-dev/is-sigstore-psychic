# Sigstore Pyschic Signature Checker

Sigstore is *not* vulnerable to the [psychic signature](https://neilmadden.blog/2022/04/19/psychic-signatures-in-java/) vulnerability in Java.

This repository demonstrates this in two ways:

1. `cmd/uploadbadsig` attempts to upload a "psychic signature" to Rekor and fails.

2. `cmd/csvcheck` takes in a CSV file containing signatures from Rekor, and scans for "psychic signatures." `signatures.csv` contains an example, including a *fake* psychic signature.
