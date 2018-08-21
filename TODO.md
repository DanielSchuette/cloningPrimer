## Main TODO's

0. check spaces in sequence and silently delete them (func ValidateSequence() and func ParseSequenceFromFile())

1. check 3' end of F and R primers (3 or more G or C bases at this position - may stabilize nonspecific annealing; 3' thymidine - more prone to mispriming)

2. check primer length (18 - 30 nucleotides)

3. self-complementarity (primer dimers)

4. GC content (40 - 60%)

5. annealing temperature between overlap of target and primer (delta 2 - 4°C and > 60°C, according to: Tm = 2°C * (A + T) + 4°C * (C + G))

6. check sequence and recognition sites for restriction sites, start codons, stop codon, ...

7. optional: 1 or 2 additional stop codons 


## More TODO's
0. increase test coverage + add codecov badge

1. Web interface

2. CLI

3. API documentation/web app documentation

4. offer access to NCBI blast API (if available?): Blast()

5. Virtual digest?

6. 'Assets' tab in web app (*.re, *.seq)

7. enable file upload for primer computation

8. enable custom primer input



