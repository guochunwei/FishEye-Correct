GOOF----LE-8-2.0é?      ] e 4 h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  BASE-VAL¤	g  stock¤	g  waste¤					 ¤	g  
foundation¤	g  reserve¤					 ¤	g  tableau¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  HORIZPOS¤	e  0.5¤	g  add-carriage-return-slot¤	g  add-blank-slot¤	g  add-extended-slot¤	 g  down¤	!g  
deal-cards¤	"											 ¤	#g  deal-cards-face-up¤	$						 ¤	%g  	get-value¤	&g  get-top-card¤	'g  add-to-score!¤	(g  give-status-message¤	)g  new-game¤	*g  set-statusbar-message¤	+g  string-append¤	,g  get-stock-no-string¤	-f     ¤	.g  get-reserve-no-string¤	/g  get-base-string¤	0g  _¤	1f  Stock left:¤	2f   ¤	3g  number->string¤	4g  length¤	5g  	get-cards¤	6f  Reserve left:¤	7f  Base Card: Ace¤	8f  Base Card: Jack¤	9f  Base Card: Queen¤	:f  Base Card: King¤	;f   ¤	<f  Base Card: ¤	=g  is-visible?¤	>g  button-pressed¤	?g  empty-slot?¤	@g  ace¤	Ag  reverse¤	Bg  king¤	Cg  get-suit¤	Dg  
droppable?¤	Eg  move-n-cards!¤	Fg  make-visible-top-card¤	Gg  button-released¤	Hg  
flip-stock¤	Ig  button-clicked¤	Jg  remove-card¤	Kg  
place-base¤	Lg  
place-card¤	Mg  button-double-clicked¤	Ng  game-won¤	Og  get-hint¤	Pg  game-continuable¤	Qf  Deal another card¤	Rg  	dealable?¤	Sg  check-a-foundation¤	Tg  check-to-foundations¤	Ug  	hint-move¤	Vg  find-empty-slot¤	Wg  check-a-tableau-list¤	Xg  	find-card¤	Yg  find-tableau-target¤	Zg  check-a-tableau-list-self¤	[g  check-a-tableau-self¤	\g  check-a-tableau¤	]g  check-tableau¤	^g  empty-tableau?¤	_g  get-options¤	`g  apply-options¤	ag  timeout¤	bg  set-features¤	cg  droppable-feature¤	dg  
set-lambda¤C 5   hX6    ] 4	 >  "  G  
R
RRR	RR !"#$%&'(    h¨    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  	
 	4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  4>   "  G  	
 	4>  "  G  4>  "  G  4>  "  G  4
>  "  G  4
>  "  G  44	55 4>  "  G  4>   "  G  		 C            g  filenamef  
kansas.scm
	
							#			3			C	 		I	 		N	 		W	"		Z	"		\	"		a	"		m	$		n	$		p	$		q	&		t	&		v	&		{	&	 	'	 	'	 	'	 	'	 	(	 	(	 	(	 ¡	(	 ª	)	 ­	)	 ¯	)	 ´	)	 ½	+	 Í	-	 Ð	-	 Ò	-	 ×	-	 à	/	 ð	0		2		2		2		4	
	4		4		4		5		5	#	5	(	5	1	6	4	6	8	6	=	6	F	8	K	8	P	8	Y	9	^	9	c	9	l	;	o	;	w	;	y	;	z	=		?	¡	A	 C	¢
  g  nameg  new-game C)R*+,-./    h       ] 445 45 45 56         g  filenamef  
kansas.scm
	C
		D			D	(		E	(		F	(		G	(		H	(		D			D	 
		
  g  nameg  give-status-message C(R+012345       h    ¬   ] 45444
5556 ¤       g  filenamef  
kansas.scm
	J
		K				K			K			K	"		L			L	!		L	)		L	!		L			K	 		
  g  nameg  get-stock-no-string C,R+062345       h    ®   ] 45444	5556¦       g  filenamef  
kansas.scm
	N
		O				O			O			O	$		P			P	!		P	)		P	!		P		 	O	 		 
  g  nameg  get-reserve-no-string C.R0789:;+<3       hp   0  ] "  >$  6	$  6	$  6	$  6C$  	$  4	54
56"ÿÿ"ÿÿ (      g  filenamef  
kansas.scm
	R
	
	V				S			W			W				X				S		!	Y		#	Y			(	Z			,	S		0	[		2	[			7	\			;	S		?	]		A	]			C	^		D	S		H	S		L	S		Q	T		U	S			X	U		\	U		^	U		_	U	)	g	U		 		o
  g  nameg  get-base-string C/R4=     h8   ã   ]
45
$  C45$   $  C 	CC    Û       g  slot-id
		4 g  	card-list		4 g  t		!	2  g  filenamef  
kansas.scm
	`
		a			a			a			b			b			b			a		!	c		!	c		1	d	 		4	  g  nameg  button-pressed C>R?%&@AB4C 
   h@  )  ] $  C"  q	$  f45$   $  CC4455$  4455"  $  C44554455CC$  °	$  £45$  45$  45"  $  C45$  C4	4554	5&  E4455$  45"  $  C445545CC"ÿþÛ"ÿþ×"ÿþÓ  !      g  
start-slot
	> g  	card-list	> g  end-slot		> g  t		Y  g  t	 »2 g  t	0  g  filenamef  
kansas.scm
	f
		g				g			u			g			v		%	u			)	w		-	u			2	{		5	{	%	=	{		@	{		D	{		E	|		H	|	*	O	|	%	Q	|		T	|		Y	{		e	}		h	}	 	p	}		q	~		t	~	*	{	~	%	}	~		~	~			}	 	g	 	h	 	g	 	i	 	h		 	j	 	j	  	h		 ¡	k	 «	k	 ®	l	# ³	l	. µ	l	# ¶	l	 »	k	 Ç	m	 Ñ	m	 Ô	n	 ×	n	& ß	n	 à	o	 å	o	& ç	o	 ë	m	 ì	p	# ï	p	. ÷	p	# ú	p	  þ	p	 ÿ	r	#	r	.	r	#		r	 	p		s	#	s	.%	s	#&	s	'	t	,	t	).	t	/	s	 H	>	  g  nameg  
droppable? CDRD'E?F!(  hè   /  ]4 5$  Õ	$  "  	$  "  45$  4 5$   $  C4	5$  C 	$  4	5"  $  C4 5$  C4	  5$  (4	5$  "  4	5$  6 CCCCC'      g  
start-slot
	 è g  	card-list	 è g  end-slot		 è g  t			A g  t		(	> g  t		X â g  t		j â g  t	  â g  t	  â g  t	 ¿ Ö  
g  filenamef  
kansas.scm
 
	 		 		 		 		( 		( 		6 		E 		F 		T 		X 		X 		d 		j 		z 		~ 		 	  	  	  	  	 © 	 ² 	 ´ 	 ¸ 	 ¹ 	 ¿ 	 Í 	 Ú 	 Þ 	  	 è	  g  nameg  button-released CGRH(      h       ] 
$  4

5$  6 CC         g  slot-id
		  g  filenamef  
kansas.scm
 
	 		
 		 		 		 	 		  g  nameg  button-clicked CIR?E&J'F!(K 
   hð     ]*45$  ×4 4 5 5$  ¾4 5$  ±45$  ¥ $  "  4	5$  "  p 	$  4	5"  $  "  K4 5$  "  34	  5$  "4	5$  "  4	5"  $  6 CCCC	 6         g  slot-id
	 í g  foundation-slot	 í g  t		= Ô g  t		Q Ñ g  t		s Î g  t	  Ë g  t	 ¬ Ã  g  filenamef  
kansas.scm
 
	 			 		 		 	#	 		  		$ 			% 		/ 			0 		9 			= 		= 		K 		Q 		c 		g 		h  		s 	  ¡	  ¡	  	  ¢	  ¢	% ¡ ¢	 ¥ ¢	 ¦ £	 ¬ £	 º ¤	 Ø 		 Ü ¥	 ë §	 í §		 #	 í	  g  nameg  
place-base CKR?C&%@BE'JF!(L    hx  7  ]*	$  C45$  "  r44 554455&  U44 55$  4455"  $  "  44 554455"  $  ×4 4 5 5$  ¾45$  ²4	 5$  ¥ $  "  4	5$  "  p 	$  4
	5"  $  "  K4 5$  "  34	  5$  "4	5$  "  4
	5"  $  6 CCCC 6     /      g  slot-id
	s g  foundation-slot	s g  t		`  g  t	 ÃZ g  t	 ×W g  t	 ùT g  t	Q g  t	2I  g  filenamef  
kansas.scm
 ©
	 ª			 ª		 ¬		 ¬			 ­		! ­		) ­		* ®		- ®		5 ®		9 ¬			: ¯		= ¯	%	E ¯		H ¯		L ¯		M °		P °	%	X °		[ °		` ¯		n ±		q ±	 	y ±		z ²		} ²	%  ²	  ²	  ±	  ª	  ³	  µ	# ¤ µ	 ¦ ³	 ª ³		 « ¶	 ´ ³		 µ ·	 ¿ ³		 Ã ¸	 Ã ¸	 Ñ ¹	 × ¸	 é º	 í º	 î »	 ù ¸	 ¼	 ¼	 ¸	 ½	% ½	%' ½	+ ½	, ¾	2 ¾	@ ¿	^ ³		b À	q Â	s Â		 ?	s	  g  nameg  
place-card CLR?%&KL      hX   ø   ]	 	$  "   $  04 5$  C44 55$   	6 	6C       ð       g  slot-id
		Q g  t		  g  filenamef  
kansas.scm
 Ä
	 Å		 Å		 Æ		 Å		  Ç		* Å		- È		0 È		8 È		; È		? È		G É		O Ë	 		Q  g  nameg  button-double-clicked CMRNO     h   u   ] 45 $  C6        m       g  filenamef  
kansas.scm
 Î
	 Ï		 Ï		 Ð	 		
  g  nameg  game-continuable CPR45        hX   ý   ] 44	55	$  :44	55	$  %44	55	$  44	55	CCCC       õ       g  filenamef  
kansas.scm
 Ò
	 Ó	
	 Ó		 Ó	
	 Ó		 Ó		 Ô	
	 Ô		! Ô	
	$ Ô		( Ó		) Õ	
	, Õ		4 Õ	
	7 Õ		; Ó		< Ö	
	? Ö		G Ö	
	J Ö	 		Q
  g  nameg  game-won CNR?0Q      h       ] 4
5$  C
45 C             g  filenamef  
kansas.scm
 Ø
	 Ù		 Ù		 Ú		 Ú		 Ú		 Ú	 		
  g  nameg  	dealable? CRR?C&%@BS  h     ]
	$  C45$  "  c4 54455&  K4 5$  4455"  $  "  4 54455"  $  C 6      g  card
	  g  foundation-slot	  g  t		V	{  g  filenamef  
kansas.scm
 Ü
	 Ý			 Ý		 ß		 ß			 à		% á		( á		0 á		4 ß			5 â		> â		B â		C ã		F ã	%	N ã		Q ã		V â		d ä		k å		n å	%	v å		w å		x ä	  Ý	  è	"  è		 	 	  g  nameg  check-a-foundation CSRT?%&UVS 
   h   ^  ] 		$  C 	$  	64 5$   644 55$   4564	4 5	5$   4	4 5	56 6   V      g  slot-id
		}  g  filenamef  
kansas.scm
 ê
	 ë			 ë		 í			 ë		 î			 ï			' ë		, ð		. ð			/ ñ		2 ñ		: ñ		= ñ			A ë		G ò		O ò			P ó			S ó		] ó			a ë		g ô		j ô	2	t ô		v ô			{ ö		} ö		 		}  g  nameg  check-to-foundations CTR4W%@BSUX 	  h¨   ù  ]45
$  C"   64545$  "  45$  45"  $  >45$  "  4	5$  45 6"ÿÿr"ÿÿn ñ      g  slot1
	 § g  slot2	 § g  card		 § g  	card-list		 § g  t		3	a g  t		n   g  filenamef  
kansas.scm
 ø
	 ù		 ù			 ù		 	0	"			" ù		# û		* ü		/ ü	%	1 ü		2 ü		3 û		3 û		A ý		J ý		N ý		O þ		T þ	%	V þ		Y þ		e ù		f ÿ		n ÿ		n ÿ		| 	  	&  	  û		 	 	0 	 	 "	 §	  g  nameg  check-a-tableau-list CWR?%&@BY    h   ¦  ]		$  C $  "  V45$  "  F445545$  "  $  	"  $  C 6           g  source-slot
	  g  source-card	  g  target-slot		  g  target-value		>	q g  source-value		>	q g  t		H	k  g  filenamef  
kansas.scm

													&			,		"	/		-	7		"	8
	"	>			G	'	H		H		Z		^		c	1	u	 	9 	 	 	  g  nameg  find-tableau-target CYRYWUXZ hX   E  ](  C4 	5$  04  5$   4 54 	56C 6   =      g  slot
		U g  top-card		U g  	card-list			U  g  filenamef  
kansas.scm

							#								&	7	(		,			1		8	.	:		;	?	B	Y	F	?	H		S	5	U	 		U	  g  nameg  check-a-tableau-list-self CZR?Z&5    h0   °   ] 		$  C4 5$  C 4 54 56  ¨       g  slot
		.  g  filenamef  
kansas.scm

										+	&	?	.	 			.  g  nameg  check-a-tableau-self C[R\[?%&@BUA54W h¸  ü  ]
		$  C	$   	6 $  4 5$  C 6$  "  	$  45$  "  d44 554455$  "  '44 55$  4455"  $  4 5"  $  C 645$  "  ]44 55$  44	4
555"  $  "   44 5544	4
555$  44
55 645$  "  4 4 54
55$   4 54
56 6  ô      g  slot1
	¶ g  slot2	¶ g  t		.	B g  t		F	[ g  t	  À g  t	 Ó ç g  t	$T  g  filenamef  
kansas.scm

						 					!			#"			'		(#		.#			@$	$	B$		F%		F%			X&		_		`'		j'		p)		s)	!	{)		|*		*	& *	 *	 )	 (	 +	 +	& ¢+	 ¥+	 ©+	 ª,	 ­,	& µ,	 ¸,	 Ä'	 Å-	 Ó'		 å.	$ ç.	 è/	 ò/		 ø0	 û0	%0	0	
0	1	1	*1	31	*1	%1	1	$0	22	52	 =2	>3	A3	*D3	3L3	*M3	%O3	P3	Q2	X	]4	`4	"h4	l4		m5	w5		}6	7	$8	$6		:	¥;	­9		´=	 ¶=		 U	¶	  g  nameg  check-a-tableau C\R?\]       h8   ¾   ] 		$  C4 5$  "  4 5$   6 6 ¶       g  slot-id
		7  g  filenamef  
kansas.scm
?
	@			@		B		B			C		)@		0D			5F		7F		 		7  g  nameg  check-tableau C]R?UV   hX   â   ]45$  C4	5  $  "  4	5$  "  4		5 $  456C  Ú       g  t
		C g  t	)	@  g  filenamef  
kansas.scm
H
	I		I		J		J		#K		)J		7L		GI		LM		PM	'	RM		TM	 		V
  g  nameg  empty-tableau? C^RT]^R       h@   ½   ]45  $   C4	5  $   C45   $   C6        µ       g  t
		9 g  t
		9 g  t
	*	9  g  filenamef  
kansas.scm
O
	P		P		Q		P		&R		*P		9S	 			9
  g  nameg  get-hint COR      h   U   ] C    M       g  filenamef  
kansas.scm
U
 		
  g  nameg  get-options C_R      h   m   ]C    e       g  options
		  g  filenamef  
kansas.scm
X
 		  g  nameg  apply-options C`R      h   Q   ] C    I       g  filenamef  
kansas.scm
[
 		
  g  nameg  timeout CaR4bici>  "  G  di)i>iGiIiMiPiNiOi_i`iaiDi6w      g  filenamef  
kansas.scm		
		
	!	
	%	
	'			*	
	/	
	1			4	
¬	
	C
w	J
i	N
	3	R

f	`
ô	f
* 
þ 
& 
 ©
s Ä
 Î
 Ò
L Ø
! Ü
# ê
&L ø
(
*P
+K
1'
29?
3H
4¨O
5U
5 X
6[
6^
6X`
 (	6X
   C6 