GOOF----LE-8-2.0Ø      ] B 4    h÷      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-carriage-return-slot¤	g  deal-cards-face-up¤											
 
¤	g  give-status-message¤	g  new-game¤	g  set-statusbar-message¤	g  string-append¤	g  get-stock-no-string¤	g  _¤	f  Stock left:¤	f   ¤	g  number->string¤	g  length¤	 g  	get-cards¤	!g  empty-slot?¤	"g  	get-value¤	#g  get-top-card¤	$g  king¤	%g  button-pressed¤	&g  
droppable?¤	'g  remove-card¤	(g  add-to-score!¤	)g  button-released¤	*g  button-clicked¤	+g  button-double-clicked¤	,g  game-won¤	-g  get-hint¤	.g  game-continuable¤	/g  club¤	0f  Remove the king of clubs.¤	1g  diamond¤	2f  Remove the king of diamonds.¤	3g  heart¤	4f  Remove the king of hearts.¤	5g  spade¤	6f  Remove the king of spades.¤	7g  hint-remove-king¤	8g  check-for-moves¤	9g  
hint-click¤	:g  get-suit¤	;g  	hint-move¤	<g  get-options¤	=g  apply-options¤	>g  timeout¤	?g  set-features¤	@g  droppable-feature¤	Ag  
set-lambda¤C 5    h   Ö   ] 4	 >  "  G         hP  Â  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
	>  "  G  4
>   "  G  		 C       º      g  filenamef  helsinki.scm
	
							#			3			C			U			e			h			m			v			y			~		 		 		 		 		 		  		 ©	 	 ¬	 	 ±	 	 º	"	 Ê	$	 Í	$	 Ò	$	 Û	%	 Þ	%	 ã	%	 ì	&	 ï	&	 ô	&	 ý	'	 	'		'		(		(		(		*	$	*	)	*	2	,	H	.	 ,	I
  g  nameg  new-game CR h   ~   ] 445 56v       g  filenamef  helsinki.scm
	0
		1			1	(		1			1	 		
  g  nameg  give-status-message CR      h    ®   ] 45444
5556 ¦       g  filenamef  helsinki.scm
	3
		4				4			4			4	"		5			5	!		5	)		5	!		5			4	 		
  g  nameg  get-stock-no-string CR!"#$ h@   æ   ] 
$  -4 5$  C45$  44 55CCC       Þ       g  slot-id
		9 g  	card-list		9  g  filenamef  helsinki.scm
	7
		8		
	8			9			8			:	
	 	:		$	8		%	;		(	;		0	;		3	;		4	;	 		9	  g  nameg  button-pressed C%R!"#     h8   û   ]45$  C
$  	454455CC      ó       g  
start-slot
		2 g  	card-list		2 g  end-slot			2  g  filenamef  helsinki.scm
	=
		>			>			?			>			@			@		!	@		"	A		%	A		-	A		.	@		/	@	 		2	  g  nameg  
droppable? C&R&'(!    hp   >  ]4 5$  \45$  O4	5$  B4
5$  "  4
  5$  4
5$  C
 6CCCC 6      g  
start-slot
		o g  	card-list		o g  end-slot			o g  t		-	H g  t		R	g  g  filenamef  helsinki.scm
	C
		D			D			E			D			F		'	D		(	G		-	G		;	H		C	H	!	E	H		L	D		M	I		R	I		e	J	!	g	J	 		o	  g  nameg  button-released C)R!"#$'(     hp     ]	4 5$  C 
$  W44 55$  B4 5$  545$  )4
5$  "  4
  5$  CCCCCC      g  slot-id
		p g  t	E	`  g  filenamef  helsinki.scm
	L
		M			M			N			M			O	
		O		#	O	
	&	O		*	M		+	P		5	M		6	Q		?	M		@	S		E	S		S	T		[	T	&	]	T		d	S	 		p  g  nameg  button-clicked C*R      h   v   ]C    n       g  slot-id
		  g  filenamef  helsinki.scm
	W
 		  g  nameg  button-double-clicked C+R,-     h(   {   ] 4>   "  G  45 $  C6        s       g  filenamef  helsinki.scm
	Z
		[			\			\		!	]	 		!
  g  nameg  game-continuable C.R!    h   ë   ] 45$  n4	5$  a4	5$  T4	5$  G4	5$  :4	5$  -4	5$   4	5$  4		5$  	
6CCCCCCCCC    ã       g  filenamef  helsinki.scm
	_
		`			`			a			`			b		"	`		#	c		-	`		.	d		8	`		9	e		C	`		D	f		N	`		O	g		Y	`		Z	h		d	`		j	i	 		|
  g  nameg  game-won C,R/0123456 
   h@   Î   ] &  6 &  6 &  6 &  	6C    Æ       g  suit
		<  g  filenamef  helsinki.scm
	k
	
	l			l			l			l			m			m		&	l		*	n		,	n		4	l		8	o		:	o	 		<  g  nameg  hint-remove-king C7R!8"#$97:; 
h¨   ¶  ] 	$  C4 5$   	 644 55$   444 5556	$   	 645$  "  	44 554455$  		 6 6   ®      g  slot1
	 ¥ g  slot2	 ¥  g  filenamef  helsinki.scm
	q
		r				r			t				r			u		"	u	&	$	u			%	v		(	v		0	v		3	v			7	r		<	w		?	w	-	B	w	7	J	w	-	L	w		N	w			S	x			W	r		\	y		a	y	&	c	y			d	z		n	z			v	{		y	{	" 	{	 	|	 	|	" 	|	 	{	 	{	 	r	 	}		 £	~	# ¥	~	 '	 ¥	  g  nameg  check-for-moves C8R8 h   ]   ] 	6       U       g  filenamef  helsinki.scm
 
		 	 			
  g  nameg  get-hint C-R      h   W   ] C    O       g  filenamef  helsinki.scm
 
 		
  g  nameg  get-options C<R    h   o   ]C    g       g  options
		  g  filenamef  helsinki.scm
 
 		  g  nameg  apply-options C=R    h   S   ] C    K       g  filenamef  helsinki.scm
 
 		
  g  nameg  timeout C>R4?i@i>  "  G  Aii%i)i*i+i.i,i-i<i=i>i&i6      Î       g  filenamef  helsinki.scm		
U	
ù	0
é	3
)	7
v	=
	A	C

è	L
y	W
6	Z
¶	_
é	k
i	q
è 
Z 
â 
N 
O 
 
 	
   C6 