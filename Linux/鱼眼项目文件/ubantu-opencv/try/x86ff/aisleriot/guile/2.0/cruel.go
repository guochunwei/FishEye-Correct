GOOF----LE-8-2.0È%      ] ^ 4        hÈ      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  ice-9¤	g  format¤	 ¤	 ¤	g  stock¤				 ¤	g  foundations¤								
						 ¤	g  	from-list¤											
						 ¤	g  to-list¤	g  original-to-slots¤	g  just-redealt¤	g  call-with-deferred-observers¤	 ¤	 ¤	g  module-export!¤	 ¤	 ¤	 g  current-module¤	!  ¤	"  ¤	# ¤	$g  variable-list¤	%g  initialize-playing-area¤	&g  set-ace-low¤	'g  make-standard-deck¤	(g  shuffle-deck¤	)g  add-normal-slot¤	*g  add-blank-slot¤	+g  add-carriage-return-slot¤	,g  
set-cards!¤	-g  remove-aces¤	.g  DECK¤	/g  
cruel-deal¤	0g  give-status¤	1g  new-game¤	2g  	get-value¤	3g  ace¤	4g  move-n-cards!¤	5g  make-visible¤	6g  empty-slot?¤	7g  deal-cards-face-up¤	8g  button-pressed¤	9g  get-suit¤	:g  get-top-card¤	;g  
droppable?¤	<g  add-to-score!¤	=g  button-released¤	>g  	dealable?¤	?g  do-deal-next-cards¤	@g  button-clicked¤	Ag  for-each¤	Bg  	flip-deck¤	C							
						 ¤	Dg  attempt-foundation¤	Eg  button-double-clicked¤	Fg  remove-card¤	Gg  set-statusbar-message¤	Hg  _¤	If  Cards remaining: ~a¤	Jg  number->string¤	Kg  	get-score¤	Lg  game-won¤	Mg  check-moves¤	Ng  headbanger?¤	Og  game-continuable¤	Pg  count¤	Qg  length¤	Rg  	get-cards¤	Sf  Redeal.¤	Tg  get-hint¤	Ug  check-move-helper¤	Vg  	hint-move¤	Wg  get-options¤	Xg  apply-options¤	Yg  timeout¤	Zg  set-features¤	[g  droppable-feature¤	\g  dealable-feature¤	]g  
set-lambda¤C 5   h   V  ] 4	 >  "  G  
RRRRRR4"#     h   ;   ] 45 6   3       g  filenamef  	cruel.scm
	
 		
   C>  "  G  $i$i%&'()*+,-./0   hà    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4	4
5>  "  G   4
>  "  G  4>   "  G  		 C           g  filenamef  	cruel.scm
	
							#			3			C	 		F	 		K	 		T	!		d	"		g	"		l	"		u	#		x	#		}	#	 	$	 	$	 	$	 	%	 	%	 	%	 ¨	&	 ¸	(	 »	(	 À	(	 É	)	 Ì	)	 Ñ	)	 Ú	*	 Ý	*	 â	*	 ë	+	 î	+	 ó	+	 ü	,	 ÿ	,		,		-		-		-		.	.	0	1	0	6	0	?	1	B	1	G	1	P	2	S	2	X	2	a	3	d	3	i	3	r	4	u	4	z	4		5		5		5		7		7	 	7	2¢	7	§	7	²	8	³	9	Ä	:	Ú	;	 E	Û
  g  nameg  new-game C1R2345-       hX   P  ] &  C4 5$  '44 5 5$   6C  6       H      g  cards
		Q g  foundation-ids		Q g  remaining-cards			Q  g  filenamef  	cruel.scm
	>
		?				?			A			A			A			A	
		A			B		#	B	$	$	B	?	)	B	M	+	B	?	.	B	9	0	B		4	B	
	9	C		<	C	(	@	C		G	D		L	D	8	O	D	2	Q	D	
 		Q	  g  nameg  remove-aces C-R67/ h8   ®   ]45$  C4	 	 >  "  G   6     ¦       g  count
		3  g  filenamef  	cruel.scm
	F
		G			G			I			I	+		I	&		I	 	$	I		1	J		3	J	 		3  g  nameg  
cruel-deal C/R6 h      ]4 5$  C 	C          g  slot-id
		 g  	card-list		  g  filenamef  	cruel.scm
	O
		P			P			Q	 			  g  nameg  button-pressed C8R9:26    h   â  ]	$  =$  C445545$  445545CC $  C45$  C445545$  445545CC      Ú      g  
start-slot
	  g  	card-list	  g  end-slot		   g  filenamef  	cruel.scm
	U
		V			V			W			W			X			X		"	X		#	X	.	(	X	7	*	X	.	+	X			/	W		0	Y		3	Y		;	Y		<	Y		=	Y	5	B	Y	?	D	Y	5	E	Y			M	[		Q	[		T	\		^	[		a	]		d	]		l	]		m	]	.	r	]	7	t	]	.	u	]			y	[		z	^		}	^	 	^	 	^	 	^	5 	^	? 	^	5 	^		 *	 	  g  nameg  
droppable? C;R6;4<     hP   õ   ]45$  C4 5$  +4 5$   $  	$  6CCCC     í       g  
start-slot
		K g  	card-list		K g  end-slot			K  g  filenamef  	cruel.scm
	c
		d			d			e			d			f		-	d		0	g		5	d		:	h		>	h		C	i	
 		K	  g  nameg  button-released C=R  h   Y   ] C  Q       g  filenamef  	cruel.scm
	n
		o	 		
  g  nameg  	dealable? C>R>?  h       ] $  45 $  6 CC           g  slot-id
		  g  filenamef  	cruel.scm
	r
		s			s			t			s			u	 		  g  nameg  button-clicked C@RA6B    h   d   ]4 5$  C 6 \       g  x
		  g  filenamef  	cruel.scm
	{			|			|			}		 		   CC/   h0      ] 4>  "  G  4
>  "  G   C    }       g  filenamef  	cruel.scm
	y
		z			 			z		 		* 	 		,
  g  nameg  do-deal-next-cards C?R6D      h(   ¡   ]4 5$  "   	$   6C          g  slot-id
		%  g  filenamef  	cruel.scm
 
	 		 		 	'	 		# 	 		%  g  nameg  button-double-clicked CER=:FD        h8   å   ](  C4 4 5 5$   6 6       Ý       g  
start-slot
		1 g  	end-slots		1  g  filenamef  	cruel.scm
 
	 		 	
	 	!	 		 		 	
	" 		( 	
	/ 	)	1 	
 		1	  g  nameg  attempt-foundation CDRGHIJK        h    ®   ] 4454	045 556¦       g  filenamef  	cruel.scm
 
	 			 	!	 	#	 	!	 	!	 	7	 	1	 	!	 		  	 		 
  g  nameg  give-status C0R0LMN     hH   º   ]4>   "  G  45 $  C  $  "  	45 $  45 CC  ²       g  t
	 	9  g  filenamef  	cruel.scm
 ¡
	 ¢		 £		 £		  ¤		  ¤		. ¤		= £		> ¥		C ¥	 		F
  g  nameg  game-continuable CORP     h    ¬   ] $   4 5C  C ¤       g  x
		 g  y		  g  filenamef  	cruel.scm
 ©
	 ª		 ª		 «		 «		 «		 «		 ¬	 				  g  nameg  count CPRK;:MP       hp   D  ]H 		045 	K 	/45 $  C$  @	045 	$  -4J 4J 5 J 5$  44	J 55CCCC     <      g  	last-slot
		k  g  filenamef  	cruel.scm
 ³
	
 ´	)	 ´	#	 ´		 ´		 µ		 µ		! µ		- ·		2 ·		5 ·		6 ·		: µ		; ¸		@ ¸	#	I ¸		L ¸	=	N ¸		R µ		S ¹		V ¹		] ¹	"	_ ¹		c ¹		d ¹	 		k
  g  nameg  headbanger? CNRQR hP   ü   ] 	4455$  :	44	55$  %	44	55$  	44	55CCCCô       g  filenamef  	cruel.scm
 ½
	 ¾		 ¾		 ¾		 ¾		 ¾		 ¿		 ¿		" ¿		# ¿		' ¾		* À		- À		5 À		6 À		: ¾		= Á		@ Á		H Á		I Á	 		P
  g  nameg  game-won CLRMHS   h(      ]45  $   C
45 C             g  t
		"  g  filenamef  	cruel.scm
 Å
	 Æ		 Æ		 Ç		 Ç		 Ç		! Ç	 		"
  g  nameg  get-hint CTRU      h      ]  6         g  
from-slots
		 g  to-slots		  g  filenamef  	cruel.scm
 Ë
	 Ì		 Í	 			  g  nameg  check-moves CMRU6;:V        hh   s  ] (  C(  	 64 5$  "  & $  "  4 4 5 5$   6 6k      g  
from-slots
		h g  to-slots		h  g  filenamef  	cruel.scm
 Ð
	 Ñ		 Ó		 Ô		 Ô	
	 Õ		 Õ	%	! Õ		% Õ		- Ö		0 Ö	,	1 Ö		5 Õ		; ×		@ ×		A Ø	%	F Ø	2	H Ø	%	K Ø		N Ù		P ×		T Õ	
	Y Ü		] Ü	-	_ Ü		f Ý	-	h Ý	 		h	  g  nameg  check-move-helper CURh   T   ] C    L       g  filenamef  	cruel.scm
 ã
 		
  g  nameg  get-options CWR       h   l   ]C    d       g  options
		  g  filenamef  	cruel.scm
 æ
 		  g  nameg  apply-options CXR       h   P   ] C    H       g  filenamef  	cruel.scm
 é
 		
  g  nameg  timeout CYR4Zi[i\i>  "  G  ]i1i8i=i@iEiOiLiTiWiXiYi;i>i6 N      g  filenamef  	cruel.scm		
		
	!			$	
	&			)	
	+			.	
	/			2	
	6	
a	
+	>
)	F
î	O
	U
è	c
\	n
'	r
	y
l 
¨ 
 
½ ¡
 ©
w ³
× ½
² Å
p Ë
n Ð
× ã
_ æ
Ë é
Ì ì
 î
 $	
   C6 