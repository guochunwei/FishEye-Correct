GOOF----LE-8-2.0`3      ] T 4  h      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  BASE-VAL¤	g  stock¤	g  waste¤					
 ¤	g  
foundation¤							 ¤	g  tableau¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  deal-cards-face-up¤								 ¤	g  add-to-score!¤	g  	get-value¤	 g  get-top-card¤	!g  give-status-message¤	"g  new-game¤	#g  set-statusbar-message¤	$g  string-append¤	%g  get-stock-no-string¤	&f     ¤	'g  get-base-string¤	(g  _¤	)f  Stock left:¤	*f   ¤	+g  number->string¤	,g  length¤	-g  	get-cards¤	.f  Base Card: Ace¤	/f  Base Card: Jack¤	0f  Base Card: Queen¤	1f  Base Card: King¤	2f   ¤	3f  
Base Card:¤	4g  empty-slot?¤	5g  button-pressed¤	6g  get-suit¤	7g  king¤	8g  ace¤	9g  
droppable?¤	:g  move-n-cards!¤	;g  button-released¤	< ¤	=g  button-clicked¤	>g  	add-card!¤	?g  remove-card¤	@g  move-to-foundation¤	Ag  button-double-clicked¤	Bg  game-won¤	Cg  get-hint¤	Dg  game-continuable¤	Ef  Deal a card¤	Fg  	dealable?¤	Gg  check-a-foundation¤	Hg  to-foundations?¤	Ig  	hint-move¤	Jg  find-empty-slot¤	Kg  waste-to-tableau?¤	Lg  check-tslot¤	Mg  tableau-movement?¤	Ng  get-options¤	Og  apply-options¤	Pg  timeout¤	Qg  set-features¤	Rg  droppable-feature¤	Sg  
set-lambda¤C 5     h+  F  ] 4	 >  "  G  
R
RRRR !       hð    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4	>   "  G  4
>  "  G  4>  "  G  4
>  "  G  4>   "  G  4	>   "  G  4	>   "  G  4	>   "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4	>   "  G  4	>   "  G  4	>   "  G  4
>  "  G  4>  "  G  4
>  "  G  4
>  "  G  4>  "  G  44	55 4>   "  G  		 C        g  filenamef  royal-east.scm
	
							#			3			C			I			N			W	 		Z	 		\	 		a	 		j	"		z	%		}	%			%	 	%	 	&	 	&	 	&	 	&	  	'	 £	'	 ¥	'	 ª	'	 ³	(	 Ã	*	 Ó	+	 ã	,	 ó	.	 ö	.	 ø	.	 ý	.		/			/		/		/		0		0		0	#	0	,	1	<	3	L	4	\	5	l	7	o	7	q	7	v	7		8		8		8		8		9		9		9		9	¥	;	ª	;	¯	;	¸	=	É	?	Ì	?	Ô	?	Ö	?	×	A	í	C	 D	î
  g  nameg  new-game C"R#$%&'      h      ] 445 45 56        g  filenamef  royal-east.scm
	F
		G			G	(		H	(		I	(		G			G	 		
  g  nameg  give-status-message C!R$()*+,-   h    °   ] 45444
5556 ¨       g  filenamef  royal-east.scm
	K
		L				L			L			L	"		M			M	!		M	)		M	!		M			L	 		
  g  nameg  get-stock-no-string C%R(./012$3*+   hx   <  ] "  >$  6	$  6	$  6	$  6C$  !	$  4	5
456"ÿÿ"ÿÿ       4      g  filenamef  royal-east.scm
	O
	
	S				P			T			T				U				P		!	V		#	V			(	W			,	P		0	X		2	X			7	Y			;	P		?	Z		A	Z			C	[		D	P		H	P		L	P		Q	Q		U	P			X	R		\	R		^	R		`	R	(	a	R	,	i	R		 		q
  g  nameg  get-base-string C'R4   hh   ;  ]
4 5$  C $  C 	$  C 	$  C 	$  C 	$  C 		C    3      g  slot-id
		d g  	card-list		d g  t			d g  t		#	d g  t		3	d g  t		C	d g  t		S	d  g  filenamef  royal-east.scm
	]
		^			^			_			_		#	`		#	_		3	a		3	_		C	b		C	_		S	c		S	_		c	d	 		d	  g  nameg  button-pressed C5R46 78        hH  Q  ]	$  "  /	$  "  	$  "  	
$  45$  45"  $  C45$  C445545&  B445545$  C4455$  45CCC	$  b $  C45$  C4455$  45"  $  C445545CC  I      g  
start-slot
	F g  	card-list	F g  end-slot		F g  t			F g  t			C g  t		+	@ g  t		e Ù g  t	 « × g  t	 ôD g  t	"D  
g  filenamef  royal-east.scm
	f
		g			g				h			g			+	i		+	g			=	j		J	g		K	k		U	k		V	l		[	l	%	]	l		`	l		e	k		q	m		{	m		~	n	 	n	& 	n	 	o	 	o	& 	o	 	m	 	p	# 	p	. ¡	p	# ¢	p	 £	q	 ¨	q	) ª	q	 «	p	 «	p	 ·	r	# º	r	. Â	r	# Å	r	  É	r	 Ê	s	# Ï	s	. Ñ	s	# Ô	s	  Þ	t		 â	g	 ç	u	 ë	u		 î	v	 ô	v	 	w		w	%	w		w		w		x		x	%	x		x	"	v	.	y	1	y	 9	y	:	z	?	z	%A	z	B	z	C	y	 D	F	  g  nameg  
droppable? C9R9:  hx   9  ]4 5$  e4 >  "  G  	$  "  /	$  "  	$  "  	
$  6CC1      g  
start-slot
		x g  	card-list		x g  end-slot			x g  t		,	k g  t		>	h g  t		P	e  g  filenamef  royal-east.scm
	}
		~			~		 			, 		, 		> 		> 		P 		P 		b 		o 			t 	 		x	  g  nameg  button-released C;R4<  h    ¨   ] 
$  4
5$  C
6C         g  slot-id
		  g  filenamef  royal-east.scm
 
	 		
 		 		 		 		 	 		  g  nameg  button-clicked C=R46 78>?@     hÐ   ;  ]
&  C45$  "  u44 554455&  W445544 55$  "  (4455$  44 55"  "  $  -44 5>  "  G  4	>  "  G  C
 6     3      g  source
	 Ë g  target	 Ë g  t		T   g  filenamef  royal-east.scm
 
	 			 		 		 	 	 		 			 		  		( 		) 		, 		1 	+	3 		5 		9 			: 		= 	%	B 	3	D 	%	F 		G 		H 		K 	 	S 		T 		T 		b 		e 	%	j 	3	l 	%	n 		q 		u 		v 		y 	%  	  	  	  		  	  	! § 		 ° 		 É 	' Ë 	 /	 Ë	  g  nameg  move-to-foundation C@R4 ?:@ 
   h8  J  ])4 5$  "  X 
$  "  D 	$  "  / 	$  "   	$  "   		$  C44 55$  4 54 >  "  G  4	5$  4 	 >  "  G  "  A4	5$  4 	 >  "  G  "  4 	
 >  "  G  4>  "  G  C4 	5$  CC   B      g  slot-id
	5 g  t			q g  t			n g  t		,	k g  t		>	h g  t		P	e g  top-card "  g  filenamef  royal-east.scm
 
	 			 			 		 			, 		, 			>  		> 			P ¡		P 			b ¢		u 		x ¤		{ ¤	  ¤	  ¤		  	  ¥	  ¥		  ¦	 ¦ §	 ° §	 ± ¨	 ¼ ¨	+ Á ¨	 Î ©	 Ø §	 Ù ª	 ä ª	+ é ª	 ö ¬	 ¬	, ¬	 ­	% ¯		+ ¯	%- ¯		1 	 (	5  g  nameg  button-double-clicked CAR!BC h(      ] 4>   "  G  45 $  C6        z       g  filenamef  royal-east.scm
 ³
	 ´		 µ		 µ		! ¶	 		!
  g  nameg  game-continuable CDR,-   hX     ] 44	55	$  :44	55	$  %44	55	$  44	
55	CCCC       ù       g  filenamef  royal-east.scm
 ¸
	 ¹	
	 ¹		 ¹	
	 ¹		 ¹		 º	
	 º		! º	
	$ º		( ¹		) »	
	, »		4 »	
	7 »		; ¹		< ¼	
	? ¼		G ¼	
	J ¼	 		Q
  g  nameg  game-won CBR4(E  h       ] 4
5$  C
45 C             g  filenamef  royal-east.scm
 ¾
	 ¿		 ¿		 À		 À		 À		 À	 		
  g  nameg  	dealable? CFR4G6 87      h°   ¾  ]
	
$  C	$  "  45$   	64 54455&  P4 54455$  "  "4 5$  4455"  $  CC 	6      ¶      g  card
	 ª g  f-slot	 ª g  t			) g  t		d   g  filenamef  royal-east.scm
 Ã
	 Ä			 Ä		 Æ		 Æ			  Ç		- Ä		6 È	"	8 È			9 É		@ Ê		C Ê		K Ê		O Ä		P Ë		W Ì		Z Ì	$	b Ì		c Ì		d Ë		d Ë		r Í		{ Í		 Í	  Î	  Î	$  Î	  Î	  Ë		 ¨ Ñ	% ª Ñ	  	 ª	  g  nameg  check-a-foundation CGR4H IJG 
h°   °  ] 	
$  C4 5$  "  / 	$  "   	$  "   	$   644 55$   4564	4 5	5$   4	4 5	56 6   ¨      g  slot-id
	 ­ g  t		S g  t		&	P g  t		8	M  g  filenamef  royal-east.scm
 Ó
	 Ô			 Ô		 Ö		 Ö			& ×		& Ö			8 Ø		8 Ö			J Ù		W Ô		\ Ú		^ Ú			a Û		d Û	 	l Û		m Û			q Ô		w Ü		 Ü		  Ý		  Ý	  Ý		  Ô	  Þ	  Þ	2 ¤ Þ	 ¦ Þ		 « ß	 ­ ß	 	 ­  g  nameg  to-foundations? CHR4K 78I 	hÈ   ò  ] 	
$  "  45$  C 	$  "   	$   64 5$  "  O4455$  44 55"  $  "  445544 55$  	 6 6      ê      g  slot-id
	 Â g  t		 g  t	(	= g  t	O ® g  t	  «  g  filenamef  royal-east.scm
 á
	 â		 â			 ã		! â		( å		( å			: æ		A â		F ç		H ç			I è		O è			] é		` é	 	g é		j é		n é		o ê		r ê	 	z ê		} ê	  è		  ë	  ë	   ë	  ë	  ì	  ì	 § ì	 ¨ ë	 ² â	 » í		 À î	 Â î	 $	 Â  g  nameg  waste-to-tableau? CKR,G4L7 8I 
    hø     ]45
$  "  
4	5$  C	
$  C45$  "  	$  "  	$   6"   645$  4455"  $  "  454455$  4 	5$  		 6"ÿÿ"ÿÿ     x      g  slot1
	 ó g  	card-list	 ó g  slot2		 ó g  t			% g  t		=	g g  t		O	d g  t	 ¨ Î  g  filenamef  royal-east.scm
 ð
	 ñ		 ñ		 ñ			 ò		 ò	!	" ò		) ñ		0 ô			4 ñ		7 ö		= ö			O ÷		O ö			a ø		k ñ		t ù	&	v ù		  	)  	  ñ	  ú	  ú	%  ú	  ú	  ú	  û	  û	%   û	 £ û	 ¨ ú	 ¶ ü	 » ü	% ½ ü	 ¾ ü	 ¿ ý	 Â ý	  Ê ý	 Ë ü	 Ò ñ	 Ó þ	 Ú þ	! Þ þ	 â ú		 ë ÿ		 .	 ó	  g  nameg  check-tslot CLR4ML- hx   4  ] 	
$  C4 5$  "   	$  "   	$   64 4 5	5$   4 5	6 6     ,      g  slot-id
		s g  t		> g  t		&	;  g  filenamef  royal-east.scm

											&		&			8		B		G		I			J				O			Y				]		b
		l
			q		s	 		s  g  nameg  tableau-movement? CMRHKMF     h@   Á   ]45  $   C4	5  $   C4	5  $   C6      ¹       g  t
		; g  t
		; g  t
	,	;  g  filenamef  royal-east.scm

									&		,		;	 			;
  g  nameg  get-hint CCR  h   Y   ] C    Q       g  filenamef  royal-east.scm

 		
  g  nameg  get-options CNR  h   q   ]C    i       g  options
		  g  filenamef  royal-east.scm

 		  g  nameg  apply-options COR  h   U   ] C    M       g  filenamef  royal-east.scm

 		
  g  nameg  timeout CPR4QiRi>  "  G  Si"i5i;i=iAiDiBiCiNiOiPi9i6    >      g  filenamef  royal-east.scm		
		
	!	
	%	
	'			*	
	,			/	
ô	
»	F
«	K
	O

>	]
ü	f
Ä	}
£ 
Ö 
} 
= ³
¬ ¸
p ¾
 Ã
  Ó
#] á
&û ð
(¿
)Ü
*L
*Ô
+@
+A
+
 !	+
   C6 