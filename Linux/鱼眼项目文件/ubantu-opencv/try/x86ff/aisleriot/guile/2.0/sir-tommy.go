GOOF----LE-8-2.0Ú      ] D 4  hó      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  stock¤	g  waste¤	g  add-blank-slot¤	g  
foundation¤	g  add-carriage-return-slot¤	g  add-extended-slot¤	g  down¤	g  reserve¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	g  get-stock-no-string¤	g  string-append¤	 g  _¤	!f  Stock left:¤	"f   ¤	#g  number->string¤	$g  length¤	%g  	get-cards¤	&g  empty-slot?¤	'g  button-pressed¤	(g  	get-value¤	)g  ace¤	*g  get-top-card¤	+g  
droppable?¤	,g  move-n-cards!¤	-g  add-to-score!¤	.g  button-released¤	/g  deal-cards-face-up¤	0 ¤	1g  button-clicked¤	2g  check-top-card¤	3g  
deal-cards¤	4g  button-double-clicked¤	5g  game-won¤	6g  get-hint¤	7g  game-continuable¤	8g  check-to-foundation¤	9g  	hint-move¤	:f  Move waste on to a reserve slot¤	;g  
move-waste¤	<f  Deal another card¤	=g  	dealable?¤	>g  get-options¤	?g  apply-options¤	@g  timeout¤	Ag  set-features¤	Bg  droppable-feature¤	Cg  
set-lambda¤C 5h  è   ] 4	 >  "  G       hx    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4	>   "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  4	>   "  G  4	>   "  G  4	>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  		 C              g  filenamef  sir-tommy.scm
	
							#			3			C			I			N			W			Z			\			a			j			z			}					 		 		 		 		 		  		 £		 ¥		 ª		 ³		 ¶		 ¸		 ½		 Æ	 	 Ö	!	 æ	"	 ö	#		$			$		$		$		%		%	"	%	'	%	0	&	3	&	7	&	<	&	E	'	H	'	L	'	Q	'	Z	)	p	*	 4	q
  g  nameg  new-game CR     h   o   ] 45 6     g       g  filenamef  sir-tommy.scm
	,
		-			-	 		
  g  nameg  give-status-message CR !"#$%    h    ¯   ] 45444
5556 §       g  filenamef  sir-tommy.scm
	/
		0				0			0			0	"		1			1	!		1	)		1	!		1			0	 		
  g  nameg  get-stock-no-string CR&$      h8   Ö   ]
4 5$  C45$   $  C 	CC     Î       g  slot-id
		3 g  	card-list		3 g  t		 	1  g  filenamef  sir-tommy.scm
	3
		4			4			5	
		5			4		 	6		 	6		0	7	 
		3	  g  nameg  button-pressed C'R()&*   hh   [  ]	$   C$  J45$  45"  $  C45$  C454455CC    S      g  
start-slot
		d g  	card-list		d g  end-slot			d g  t		3	b  g  filenamef  sir-tommy.scm
	9
		:				:			;				<				:			=			=	 	 	=		#	=		'	=		(	>		3	=			?	?		I	?		L	@		Q	@	 	S	@		T	A		W	A	%	_	A		`	A		a	@	 		d	  g  nameg  
droppable? C+R+,-        hP   é   ]4 5$  8	$  
 6$  4 >  "  G  6CC     á       g  
start-slot
		K g  	card-list		K g  end-slot			K  g  filenamef  sir-tommy.scm
	D
		E			E			F			F		$	G		(	H		,	F		-	J		G	K	 		K	  g  nameg  button-released C.R&/0  h0   °   ] 
$  4
5$  C45$  
6CC     ¨       g  slot-id
		+  g  filenamef  sir-tommy.scm
	N
		O		
	O			P			O			Q		 	O		%	R		'	R	 
		+  g  nameg  button-clicked C1R&(*2)       h   I  ]	$  C45$  "  44 554455$   C"  	 644 55$  45$   C"ÿÿÏ"ÿÿË      A      g  slot
		z g  f-slot		z  g  filenamef  sir-tommy.scm
	T
		U				U			W			W				X		!	X		)	X		*	Y		-	Y	!	5	Y		6	Y		7	X		;	U		@	Z			L	_	!	N	_		N	U		O	[		R	[		Z	[		]	[		a	U		b	]		l	[			q	^		 		z	  g  nameg  check-top-card C2R&23-        hX   í   ]	4 5$  C $  "   	$  *4 	5$  4 4 	55$  6CCCå       g  slot-id
		X g  t		(  g  filenamef  sir-tommy.scm
	a
		b			b			c			c		%	d		,	b		-	e		9	b		:	f		?	f		I	f		M	b		R	g	 		X  g  nameg  button-double-clicked C4R56      h(   |   ] 4>   "  G  45 $  C6        t       g  filenamef  sir-tommy.scm
	i
		j			k			k		!	l	 		!
  g  nameg  game-continuable C7R$% hX   ì   ] 44	55	$  :44	55	$  %44	55	$  44	55	CCCC       ä       g  filenamef  sir-tommy.scm
	n
		o	
		o			o	
		o			o			p	
		p		!	p	
	$	p		(	o		)	q	
	,	q		4	q	
	7	q		;	o		<	r	
	?	r		G	r	
	J	r	 		Q
  g  nameg  game-won C5R8&29     hX   â   ] 	
$  C 	$  	64 5$  "  	4 	5$   4 	56 6       Ú       g  slot
		Q  g  filenamef  sir-tommy.scm
	t
		u				u			w				u			x				y		'	y			-	z		9	u		?	{	 	H	{		J	{			O	|	!	Q	|	 		Q  g  nameg  check-to-foundation C8R& : h(      ] 45$  C4
5$  C
45 C         g  filenamef  sir-tommy.scm
	~
							 					 		  		" 		% 	 
		&
  g  nameg  
move-waste C;R& <        h       ] 4
5$  C
45 C             g  filenamef  sir-tommy.scm
 
	 		 		 		 		 		 	 		
  g  nameg  	dealable? C=R8;=       h(       ]45  $   C45   $   C6         g  t
		' g  t
		'  g  filenamef  sir-tommy.scm
 
	 		 		 		 		' 	 		'
  g  nameg  get-hint C6R   h   X   ] C    P       g  filenamef  sir-tommy.scm
 
 		
  g  nameg  get-options C>R   h   p   ]C    h       g  options
		  g  filenamef  sir-tommy.scm
 
 		  g  nameg  apply-options C?R   h   T   ] C    L       g  filenamef  sir-tommy.scm
 
 		
  g  nameg  timeout C@R4AiBi>  "  G  Cii'i.i1i4i7i5i6i>i?i@i+i6     à       g  filenamef  sir-tommy.scm		
Ë	
b	,
R	/
y	3
V	9
	¬	D

£	N
	T
ð	a
¯	i
	n
]	t
6	~
ÿ 
ã 
S 
Û 
G 
H 
 
 	
   C6 