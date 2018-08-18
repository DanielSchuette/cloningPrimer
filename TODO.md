## Main TODO's
1. Methylation can block cleavage by some restriction enzymes. In E. coli, Dam methylase affects the sequence GATC and Dcm methylase affects the sequence CCAGG or CCTGG. If these sequences are present, you will need to use a dam-, dcm- E. coli strain to grow your plasmid.

2. check 3' end of primer (3 or more G or C bases at this position - may stabilize nonspecific annealing; 3' thymidine - more prone to mispriming)

3. primer length (18 - 30 nucleotides)

4. self-complementarity (primer dimers)

5. GC content (40 - 60%)

6. annealing temperature between overlap of target and primer (delta 2 - 4°C and > 60°C, according to: Tm = 2°C * (A + T) + 4°C * (C + G))

7. N-terminal ATG in-frame with ORF

8. C-terminal TAA (preferred because less prone to read-through than TAG and TGA; 2 or 3 stop codons are possible) in-frame with ORF


## More TODO's
1. add more unit tests

2. set up Travis CI

3. Web interface

4. CLI

5. API documentation

6. calculate stats (see above): func CalcStats()

7. check sequence and recognition sites for restriction sites, start codons, stop codon, ...

8. pull plasmid information from file and test for restriction sites: func CheckPlasmid ()

9. offer access to NCBI blast API (if available?): Blast()

10. check spaces in sequence and silently delete them
