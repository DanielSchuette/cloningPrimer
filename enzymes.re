/* This file provides information about restriction enzymes --
 * namely, their identifiers, recognition sequences (5' to 3'),  
 * PDB identifiers, non-palindromic cleaveage [(before)(after)],
 * and 8 - 10 common isoschizomers --
 * to a software that is able to interpret/use this information.
 *
 * Special nucleotide codes:
 * B        C or G or T
 * D        A or G or T
 * H        A or C or T
 * K        G or T
 * M        A or C
 * N        A or C or G or T
 * R        A or G 
 * S        C or G
 * V        A or C or G
 * W        A or T
 * Y        C or T
 *
 * All data is derived from the following sources:
 * https://www.rcsb.org/
 * https://www.wwpdb.org/
 * http://rebase.neb.com/rebase/rebase.html
 * https://www.neb.com/products/restriction-endonucleases
 *
 * Author: Daniel Schuette
 * Last update: Oct 11, 2018
 *
 * Legal Disclaimer:
 * This file is provided "as is". The author does not provide any warranty of the item whatsoever, whether express,
 * implied, or statutory, including, but not limited to, any warranty of merchantability or fitness for a particular 
 * purpose or any warranty that the contents of the item will be error-free. In no respect shall the author incur any 
 * liability for any damages, including, but limited to, direct, indirect, special, or consequential damages arising 
 * out of, resulting from, or any way connected to the use of this file; whether or not injury was sustained by persons 
 * or property or otherwise; and whether or not loss was sustained from, or arose out of, or is the results of, this file.
 *
 * Email to d.schuette(at)online.de for more information.
 */
enzyme_name recognition_sequence                palindromic_cleaveage   PDB_ID  isoschizomers
AclI        AACGTT                              no
HindIII     AAGCTT                              no 
SspI        AATATT                              no 
MluCI       AATT                                no	
PciI        ACATGT                              no	
AgeI        ACCGGT                              no	
BspMI       ACCTGC                              ()(4/8)
BfuAI       ACCTGC                              ()(4/8)
SexAI       ACCWGGT                             no
MluI        ACGCGT                              no	
BceAI       ACGGC                               ()(12/14)
HpyCH4IV    ACGT                                no
HpyCH4III   ACNGT                               no
BaeI        ACNNNNGTAYC                         (10/15)(12/7)
BsaXI       ACNNNNNCTCC                         (9/12)(10/7)
AflIII      ACRYGT                              no
SpeI        ACTAGT                              no
BsrI        ACTGG                               ()(1/-1)	
BmrI        ACTGGG                              ()(5/4)	
BglII       AGATCT                              no
AfeI        AGCGCT                              no
AluI        AGCT                                no
StuI        AGGCCT                              no
ScaI        AGTACT                              no
BspDI       ATCGAT                              no
ClaI        ATCGAT                              no
PI-SceI     ATCTATGTCGGGTGCGGAGAAAGAGGTAAT      ()(-15/-19)	
NsiI        ATGCAT                              no	
AseI        ATTAAT	                            no
SwaI        ATTTAAAT	                        no
CspCI       CAANNNNNGTGG	                    (11/13)(12/10)
MfeI        CAATTG	                            no 
BssSαI      CACGAG                              ()(-5/-1)	
Nb.BssSI    CACGAG	                            no
BmgBI       CACGTC                              ()(-3/-3)	
PmlI        CACGTG	                            no
DraIII      CACNNNGTG	                        no
AleI        CACNNNNGTG	                        no
EcoP15I     CAGCAG                              ()(25/27)	
PvuII       CAGCTG	                            no
AlwNI       CAGNNNCTG	                        no
BtsIMutI    CAGTG                               ()(2/0)	
NdeI        CATATG	                            no
CviAII      CATG	                            no
FatI        CATG	                            no
NlaIII      CATG                                no
MslI        CAYNNNNRTG	                        no
FspEI       CC                                  ()(12/16)	
XcmI        CCANNNNNNNNNTGG	                    no
BstXI       CCANNNNNNTGG	                    no
PflMI       CCANNNNNTGG	                        no
BccI        CCATC                               ()(4/5)	
NcoI        CCATGG	                            no
BseYI       CCCAGC                              ()(-5/-1)	
FauI        CCCGC                               ()(4/6)	
SmaI        CCCGGG	                            no
XmaI        CCCGGG	                            no
TspMI       CCCGGG	                            no
Nt.CviPII   CCD	                                (0/-1)()
LpnPI       CCDG                                ()(10/14)	
AciI        CCGC                                ()(-3/-1)	
SacII       CCGCGG	                            no
BsrBI       CCGCTC                              ()(-3/-3)	
HpaII       CCGG	                            no
MspI        CCGG                                no
ScrFI       CCNGG	                            no
StyD4I      CCNGG                               no	
BsaJI       CCNNGG                              no	
BslI        CCNNNNNNNGG	                        no
BtgI        CCRYGG	                            no
NciI        CCSGG                               no	
AvrII       CCTAGG                              no	
MnlI        CCTC                                ()(7/6)	
Nb.BbvCI    CCTCAGC	                            no
BbvCI       CCTCAGC                             ()(-5/-2)	
Nt.BbvCI    CCTCAGC                             ()(-5/-7)	
SbfI        CCTGCAGG                            no	
Bpu10I      CCTNAGC                             ()(-5/-2)	
Bsu36I      CCTNAGG	                            no
EcoNI       CCTNNNNNAGG	                        no
HpyAV       CCTTC                               ()(6/5)	
BstNI       CCWGG	                            no
PspGI       CCWGG	                            no
StyI        CCWWGG	                            no
BcgI        CGANNNNNNTGC                        (10/12)(12/10)	
PvuI        CGATCG                              no	
BstUI       CGCG                                no	
EagI        CGGCCG                              no	
RsrII       CGGWCCG                             no	
BsiEI       CGRYCG                              no	
BsiWI       CGTACG                              no	
BsmBI       CGTCTC                              ()(1/5)	
Hpy99I      CGWCG                               no
MspA1I      CMGCKG                              no	
AbaSI       CNNNNNNNNNNNNNNNNNNNNG	            no
MspJI       CNNR                                ()(9/13)	
SgrAI       CRCCGGYG	                        no
BfaI        CTAG	                            no
BspCNI      CTCAG                               ()(9/7)	
XhoI        CTCGAG	                            no
PaeR7I      CTCGAG                              no
EarI        CTCTTC                              ()(1/4)	
AcuI        CTGAAG                              ()(16/14)	
PstI        CTGCAG                              no	
BpmI        CTGGAG                              ()(16/14)	
DdeI        CTNAG	                            no
SfcI        CTRYAG	                            no
AflII       CTTAAG                              no	
BpuEI       CTTGAG                              ()(16/14)	
SmlI        CTYRAG	                            no
AvaI        CYCGRG	                            no
BsoBI       CYCGRG	                            no
MboII       GAAGA                               ()(8/7)	
BbsI        GAAGAC                              ()(2/6)	
XmnI        GAANNNNTTC	                        no
BsmI        GAATGC                              ()(1/-1)	
Nb.BsmI     GAATGC	                            no
EcoRI       GAATTC	                            no
HgaI        GACGC                               ()(5/10)	
ZraI        GACGTC	                            no
AatII       GACGTC	                            no
Tth111I     GACNNNGTC	                        no
PflFI       GACNNNGTC	                        no
PshAI       GACNNNNGTC	                        no
AhdI        GACNNNNNGTC	                        no
DrdI        GACNNNNNNGTC	                    no
Eco53kI     GAGCTC	                            no
SacI        GAGCTC                              no	
BseRI       GAGGAG                              ()(10/8)	
PleI        GAGTC                               ()(4/5)	
MlyI        GAGTC                               ()(5/5)	
Nt.BstNBI   GAGTC                               ()(4/-5)	
HinfI       GANTC	                            no
EcoRV       GATATC	                            no
Sau3AI      GATC	                            no
MboI        GATC	                            no
DpnII       GATC	                            no
DpnI        GATC	                            no
BsaBI       GATNNNNATC	                        no
TfiI        GAWTC	                            no
Nb.BsrDI    GCAATG	                            no
BsrDI       GCAATG                              ()(2/0)	
BbvI        GCAGC                               ()(8/12)	
Nb.BtsI     GCAGTG	                            no
BtsαI       GCAGTG                              ()(2/0)	
BstAPI      GCANNNNNTGC	                        no
SfaNI       GCATC                               ()(5/9)	
SphI        GCATGC	                            no
SrfI        GCCCGGGC	                        no
NmeAIII     GCCGAG                              ()(21/19)	
NgoMIV      GCCGGC	                            no
NaeI        GCCGGC	                            no
BglI        GCCNNNNNGGC	                        no
AsiSI       GCGATCGC	                        no
BtgZI       GCGATG                              ()(10/14)	
HinP1I      GCGC	                            no
HhaI        GCGC	                            no
BssHII      GCGCGC	                            no
NotI        GCGGCCGC	                        no
Fnu4HI      GCNGC	                            no
Cac8I       GCNNGC	                            no
MwoI        GCNNNNNNNGC	                        no
NheI        GCTAGC	                            no
BmtI        GCTAGC	                            no
Nt.BspQI    GCTCTTC                             ()(1/-7)	
BspQI       GCTCTTC                             ()(1/4)	
SapI        GCTCTTC                             ()(1/4)	
BlpI        GCTNAGC	                            no
TseI        GCWGC	                            no
ApeKI       GCWGC	                            no
Bsp1286I    GDGCHC	                            no
AlwI        GGATC                               ()(4/5)	
Nt.AlwI     GGATC                               ()(4/-5)	
BamHI       GGATCC	                            no
BtsCI       GGATG                               ()(2/0)	
FokI        GGATG                               ()(9/13)	
HaeIII      GGCC	                            no
FseI        GGCCGGCC	                        no
SfiI        GGCCNNNNNGGCC	                    no
SfoI        GGCGCC	                            no
PluTI       GGCGCC	                            no
NarI        GGCGCC	                            no
KasI        GGCGCC	                            no
AscI        GGCGCGCC	                        no
EciI        GGCGGA                              ()(11/9)	
BsmFI       GGGAC                               ()(10/14)	
PspOMI      GGGCCC	                            no
ApaI        GGGCCC	                            no
Sau96I      GGNCC	                            no
NlaIV       GGNNCC	                            no
Acc65I      GGTACC	                            no
KpnI        GGTACC	                            no
BsaI        GGTCTC                              ()(1/5)	
HphI        GGTGA                               ()(8/7)	
BstEII      GGTNACC	                            no
AvaII       GGWCC	                            no
BanI        GGYRCC	                            no
BaeGI       GKGCMC                              no	
BsaHI       GRCGYC                              no	
BanII       GRGCYC                              no	
CviQI       GTAC                                no	
RsaI        GTAC                                no	
BstZ17I     GTATAC	                            no
BciVI       GTATCC                              ()(6/5)	
SalI        GTCGAC                              no	
Nt.BsmAI    GTCTC                               ()(1/-5)	
BcoDI       GTCTC                               ()(1/5)	
BsmAI       GTCTC                               ()(1/5)
ApaLI       GTGCAC                              no	
BsgI        GTGCAG                              ()(16/14)	
AccI        GTMKAC	                            no
Hpy166II    GTNNAC	                            no
Tsp45I      GTSAC	                            no
HpaI        GTTAAC	                            no
PmeI        GTTTAAAC	                        no
HincII      GTYRAC	                            no
BsiHKAI     GWGCWC	                            no
TspRI       NNCASTGNN                           no
ApoI        RAATTY	                            no
NspI        RCATGY	                            no
BsrFαI      RCCGGY	                            no
BstYI       RGATCY	                            no
HaeII       RGCGCY	                            no
CviKI-1     RGCY	                            no
EcoO109I    RGGNCCY	                            no
PpuMI       RGGWCCY	                            no
I-CeuI      TAACTATAACGGTCCTAAGGTAGCGAA         ()(-9/-13)	
SnaBI       TACGTA	                            no
I-SceI      TAGGGATAACAGGGTAAT                  ()(-9/-13)	
BspHI       TCATGA	                            no
BspEI       TCCGGA	                            no
MmeI        TCCRAC                              (20/18)	
TaqαI       TCGA	                            no
NruI        TCGCGA	                            no
Hpy188I     TCNGA	                            no
Hpy188III   TCNNGA	                            no
XbaI        TCTAGA	                            no
BclI        TGATCA	                            no
HpyCH4V     TGCA	                            no
FspI        TGCGCA                              no	
PI-PspI     TGGCAAACAGCTATTATGGGTATTATGGGT      ()(-13/-17)	
MscI        TGGCCA	                            no
BsrGI       TGTACA	                            no
MseI        TTAA	                            no
PacI        TTAATTAA	                        no
PsiI        TTATAA	                            no
BstBI       TTCGAA	                            no
DraI        TTTAAA	                            no
PspXI       VCTCGAGB	                        no
BsaWI       WCCGGW	                            no
BsaAI       YACGTR	                            no
EaeI        YGGCCR	                            no
