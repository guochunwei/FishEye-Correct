GOOF----LE-8-2.0*      ] W 4       hu      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  allow-two-spot-use¤											
	 
¤	g  tableau¤	
 ¤	g  	tmp-spots¤	g  stock¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-blank-slot¤	g  add-normal-slot¤	g  DECK¤	g  add-carriage-return-slot¤	g  map¤	g  add-extended-slot¤	g  down¤	g  set-slot-y-expansion!¤	e  0.25¤	 g  deal-ten-across-cards¤	!g  deal-cards-face-up¤	" ¤	#g  flip-top-card¤	$g  new-game¤	%g  length¤	&g  	list-head¤	'g  
deal-cards¤	(g  	list-tail¤	)g  reverse¤	*				 ¤	+g  member¤	,g  is-visible?¤	-g  button-pressed¤	.g  move-n-cards!¤	/g  empty-slot?¤	0g  make-visible-top-card¤	1g  complete-transaction¤	2g  get-suit¤	3g  	get-value¤	4g  is-ok-to-place¤	5g  king¤	6g  get-top-card¤	7g  
droppable?¤	8g  button-released¤	9g  button-clicked¤	:g  button-double-clicked¤	;g  or-map¤	<g  	get-cards¤	=g  have-empty-slot?¤	>g  king?¤	?g  find-king-move¤	@g  	hint-move¤	Ag  find-placeable-card¤	Bg  find-stack-move¤	Cg  test-stack-move¤	Dg  get-top-cards¤	Eg  append¤	Fg  _¤	Gf  &Move a card to an empty temporary slot¤	Hf  No hint available¤	Ig  get-hint¤	Jg  final-stack-helper¤	Kg  final-stack?¤	Lg  
won-tester¤	Mg  game-won¤	Ng  	game-over¤	Of  Allow temporary spots use¤	Pg  get-options¤	Qg  apply-options¤	Rg  timeout¤	Sg  set-features¤	Tg  droppable-feature¤	Ug  scores-disabled¤	Vg  
set-lambda¤C 5      h8"  X  ] 4	 >  "  G  RRR
R  h   f   ]6       ^       g  ignore
			  g  filenamef  ten-across.scm
	'			'	+			'	 			   C        h   d   ] 6      \       g  slot
		
  g  filenamef  ten-across.scm
	(			*		
	)		 		
   C !"#       hð   î   ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>   "  G  4	
>  "  G  4	>  "  G  4>   "  G  4>  "  G  4>  "  G  	
	 C     æ       g  filenamef  ten-across.scm
	
							#	 		3	!		C	#		S	$		e	%		h	%		m	%		v	&	 	'	 	(	 ®	,	 ¾	.	 Ä	.	 É	.	 Ò	0	 ê	2	 	 ë
  g  nameg  new-game C$R%!&'()      hè     ]M$  e44 5>  "  G  444 5L   5>  "  G  44L  5>  "  G  "  p444L  55>  "  G  4444 5L   55>  "  G  444 55>  "  G  MNC          g  num-now
	 ä  g  filenamef  ten-across.scm
	J			9				;			;	,		;		$	<		)	=		,	>		;	>	7	>	=		C	<		L	?		Q	?	,	Z	?	?	\	?	,	a	?		n	A		s	B	&	v	B	/		C	: 	B	/ 	B	& 	A	 	D	 	E	 	E	' 	F	( «	G	( ®	E	' °	E	 µ	D	 ¾	H	 Ã	H	, Æ	H	5 Ð	H	, Õ	H	 à	I	 â	I	 '	 ä   C*        h    ¯   ]45 H O 6   §       g  deal-len
			 g  	direction		  g  filenamef  ten-across.scm
	4
		5				5			J	8		J	 		
  g  nameg  deal-ten-across-cards C R+%,)        hH   ç   ]
 $  "  4 5$  45"  $  456C     ß       g  slot-id
		C g  	card-list		C g  t			1  g  filenamef  ten-across.scm
	Q
		R			R			S		 	S		!	T		)	T		5	R		8	U		?	U		A	U	 		C	  g  nameg  button-pressed C-R./+0        hP   ×   ]4 >  "  G  4 5$  "  	4 5$  4 >  "  G  "   C   Ï       g  
start-slot
		M g  	card-list		M g  end-slot			M  g  filenamef  ten-across.scm
	W
		X			Y		#	Y		)	Z		5	Y		6	[	 		M	  g  nameg  complete-transaction C1R23      h(   Â   ]4 545$  454 5CCº       g  card1
		( g  card2		(  g  filenamef  ten-across.scm
	^
		_	
	
	`	
		_			_			a	
		b		$	b	
	%	a	 
		(	  g  nameg  is-ok-to-place C4R+/53)46% h   r  ] $  C45$  745$  4455"  445455"  $  C	$  $4
5$  45$  6CCC       j      g  
start-slot
	  g  	card-list	  g  end-slot		  g  t		R   g  filenamef  ten-across.scm
	d
		e			e			f			f			g		%	g		(	h		+	h	,	2	h	'	4	h		5	h		:	i		=	i	)	D	i	$	E	j	$	M	i		R	f		c	k		d	l		p	k		r	m		y	m		}	k	 	n	 	 	  g  nameg  
droppable? C7R71   h    º   ]4 5$  
 6C   ²       g  
start-slot
		 g  	card-list		 g  end-slot			  g  filenamef  ten-across.scm
	p
		q			q			r	 			  g  nameg  button-released C8R h   t   ]C    l       g  
start-slot
		  g  filenamef  ten-across.scm
	t
 		  g  nameg  button-clicked C9R       h   {   ]C    s       g  
start-slot
		  g  filenamef  ten-across.scm
	w
 		  g  nameg  button-double-clicked C:R;%<       h   y   ]
44 55C       q       g  item
		  g  filenamef  ten-across.scm
 	
	 		 	&	 		 	 		   C    h      ] 6      w       g  	slot-list
		
  g  filenamef  ten-across.scm
 
	
 	 		
  g  nameg  have-empty-slot? C=R35      h   x   ]4 5C   p       g  card
		  g  filenamef  ten-across.scm
 
	 		 	 		  g  nameg  king? C>R,35?@%+ 	h   µ  ](  C"  A45$  "  45$   6 645$  &45$  4 5$  C"ÿÿ"ÿÿ"ÿÿ  ­      g  slot
	  g  	card-list	  g  count		  g  to-slot		  g  t			3  g  filenamef  ten-across.scm
 
	 		 ¢		 ¢		 ¢		 ¢		 ¢			% £		* £	 	, £		/ £		0 £		7 		> ¤		A ¤	.	E ¤			O ¦			O 		P 		X 		\ 		] 		b 		d 		g 		k 			l 		x 		 	 	  g  nameg  find-king-move C?R,4A      h8     ] (  C4 5$  4 5$  C 6C         g  	card-list
		5 g  	dest-card		5 g  count			5  g  filenamef  ten-across.scm
 ©
	 ª		 ¬		 ¬		 ¬		 ª		 ®			 ®		  ®			$ ª		, ±		1 ±	8	3 ±		 		5	  g  nameg  find-placeable-card CARB/?<A6@    h   Õ  ]
(  C $  	 645$  &4 4 55$  C 644 5455$   44 54556 6 Í      g  	from-slot
	  g  	slot-list	  g  t		;	O  g  filenamef  ten-across.scm
 ³
	 ´		 ¶		 ¶			 ´		 ·	$	 ·			 ¸			# ¸		% ¸			) ´		* ¹		/ ¹	'	9 ¹	?	; ¹		; ¹			M º	(	O º		P »			S »		Z »	4	_ »	B	a »	4	d »			h ´		m ¼		p ¼	3	w ¼	I	| ¼	W	~ ¼	I  ¼	  ¼	k  ¼		  ¾	$  ¾		 $	 	  g  nameg  find-stack-move CBRBC      h(   Â   ]	 (  C4 5$  C 6  º       g  	slot-list
		& g  t		&  g  filenamef  ten-across.scm
 À
	 Á		 Ã	
	 Ã		 Ã	
	 Ã		$ Ä		& Ä	
 			&  g  nameg  test-stack-move CCR<  h      ]	4 5(  CC        g  slot
		 g  cards			  g  filenamef  ten-across.scm
 È		 É			 É			 Ê		 Ë		 Ì	 		   C h   |   ] 6      t       g  	slot-list
		
  g  filenamef  ten-across.scm
 Ç
	
 È	 		
  g  nameg  get-top-cards CDRCE=FGH 
  hX     ]4455  $   C$  45$  
45 "  "    $   C
4	5 Cü       g  t
		X g  t
	A	X  g  filenamef  ten-across.scm
 Ð
	 Ò		 Ò		 Ò		 Ñ		! Ó		" Ô		, Ó		. Õ		2 Õ		4 Õ		7 Õ		A Ñ		N Ö		R Ö		T Ö		W Ö	 		X
  g  nameg  get-hint CIR,23J     hP   =  ] (  C  45$  +45$  45$  6CCC    5      g  the-list
		L g  num		L g  suit			L g  card			L g  rest			L  g  filenamef  ten-across.scm
 Ù		 Ú		 Ü		 Ý		 Ü		 Þ		 Þ	
	" ß		) ß		- Þ		0 à		7 à		; Þ		B á	'	F á	 		L	  g  nameg  final-stack-helper CJRJ2        h      ] 4 56              g  	card-list
		  g  filenamef  ten-across.scm
 ä
	 å	$	 å	.	 å	$	 å	 		  g  nameg  final-stack? CKR%<KL   hh   ]  ]  	4455$  4455"  $  "  
4455$  &  C6C     U      g  	slot-list
		c g  to-test		c g  to-cont			c g  t		0	N  g  filenamef  ten-across.scm
 è		 é		 ê		 é		 ë		 ë	!	 ë		 ë		 ë		  ì		# ì	!	+ ì		0 ë	
	? í		B í		J í		K í		R ë		S î		Y î	
	a ð	 		c  g  nameg  
won-tester CLRL        h   _   ] 6W       g  filenamef  ten-across.scm
 ó
	 ô	 		
  g  nameg  game-won CMRM        h   i   ] 45 C       a       g  filenamef  ten-across.scm
 ö
	 ÷		 ÷	 			
  g  nameg  	game-over CNRFO  h      ] 45  C      ~       g  filenamef  ten-across.scm
 ù
	 ú		 ú			 ú		 ú		 ú	 		
  g  nameg  get-options CPR h      ]  C     {       g  options
		  g  filenamef  ten-across.scm
 ü
	 ý			 ý	 		  g  nameg  apply-options CQRh   U   ] C    M       g  filenamef  ten-across.scm
 ÿ
 		
  g  nameg  timeout CRR4SiTiUi>  "  G  Vi$i-i8i9i:iNiMiIiPiQiRi7i6P      g  filenamef  ten-across.scm		
		
				"	
	$			'	
	+	
Q	
Ú	4
*	Q
	r	W

u	^
	d
	p
	t
®	w
ú 
 
ø 
Z ©
à ³
å À
? Ç
¿ Ð
h Ø
1 ä
 ç
 ó
 $ ö
 Ù ù
!~ ü
!è ÿ
!é
"8
 #	"8
   C6 