// Package cloningprimer ...
//
// Resources:
// https://www.neb.com/protocols/1/01/01/primer-design-e6901
// https://www.addgene.org/mol-bio-reference/restriction-enzymes/
// TODO: Methylation can block cleavage by some restriction enzymes. In E. coli, Dam methylase affects the sequence GATC and Dcm methylase affects the sequence CCAGG or CCTGG. If these sequences are present, you will need to use a dam-, dcm- E. coli strain to grow your plasmid.
// https://www.embl.de/pepcore/pepcore_services/cloning/pcr_strategy/primer_design/
// TODO: check 3' end of primer (3 or more G or C bases at this position - may stabilize nonspecific annealing; 3' thymidine - more prone to mispriming)
// TODO: primer length (18 - 30 nucleotides)
// TODO: self-complementarity (primer dimers)
// TODO: GC content (40 - 60%)
// TODO: annealing temperature between overlap of target and primer (delta 2 - 4°C and > 60°C, according to: Tm = 2°C * (A + T) + 4°C * (C + G))
// TODO: N-terminal ATG in-frame with ORF
// TODO: C-terminal TAA (preferred because less prone to read-through than TAG and TGA; 2 or 3 stop codons are possible) in-frame with ORF
//
// Code Example:
//
//
package cloningprimer
