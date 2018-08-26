## Main TODO's

0. check spaces in sequence and silently delete them (func ValidateSequence() and func ParseSequenceFromFile())

1. check 3' end of F and R primers (3 or more G or C bases at this position - may stabilize nonspecific annealing; 3' thymidine - more prone to mispriming)

2. check primer length (18 - 30 nucleotides)

3. self-complementarity (primer dimers)

4. GC content (40 - 60%)

5. add more methods for Tm calculation

6. Tm between overlap of target and primer (delta 2 - 4°C and > 60°C)

7. check sequence and recognition sites for restriction sites, start codons, stop codon, ...

8. optional: 1 or 2 additional stop codons 

9. HasRestrictionSite(): check if a certain sequence has a certain restriction enzyme recognition site

10. HasStartInSeq(): check if a certain sequence has more than one start codon

11. HasStopInSeq(): check if a certain sequence has more than one stop codon



## More TODO's
0. increase test coverage

1. Web interface

2. CLI

3. API documentation/web app documentation

4. offer access to NCBI blast API (if available?): Blast()

5. Virtual digest?

6. enable file upload for primer computation (or even sequence upload from database ?)

7. allow for checking custom cloning primers ('check your primer' tab) 

8. Python port

