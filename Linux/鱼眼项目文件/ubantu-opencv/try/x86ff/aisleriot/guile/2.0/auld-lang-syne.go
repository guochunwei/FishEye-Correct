GOOF----LE-8-2.0t      ] E 4 hω      ] g  guile€	 €	g  process-use-modules€	 €	 €	g  	aisleriot€	g  	interface€	 €		 €	
g  api€	
 €	 €	g  initialize-playing-area€	g  set-ace-low€	g  make-deck-list-ace-low€	g  club€	g  DECK€	g  shuffle-deck€	g  add-normal-slot€	g  add-blank-slot€	g  add-carriage-return-slot€	g  	add-card!€	g  make-visible€	g  	make-card€	g  ace€	g  diamond€	g  heart€	g  spade€	g  give-status-message€	g  new-game€	g  set-statusbar-message€	 g  get-stock-no-string€	!g  string-append€	"g  _€	#f  Stock left:€	$f   €	%g  number->string€	&g  length€	'g  	get-cards€	(g  empty-slot?€	)g  button-pressed€	*g  	get-value€	+g  get-top-card€	,g  
droppable?€	-g  move-n-cards!€	.g  add-to-score!€	/g  button-released€	0g  	dealable?€	1g  deal-cards-face-up€	2				 €	3g  do-deal-next-cards€	4g  button-clicked€	5g  
deal-cards€	6g  check-end-slot?€	7g  button-double-clicked€	8g  game-won€	9g  get-hint€	:g  game-continuable€	;g  movable?€	<g  	hint-move€	=f  Deal another round€	>g  get-options€	?g  apply-options€	@g  timeout€	Ag  set-features€	Bg  droppable-feature€	Cg  dealable-feature€	Dg  
set-lambda€C 5  h   ς   ] 4	 >  "  G   h°  E  ] 4>   "  G  4>   "  G  4		5 4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4	>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
4455>  "  G  4
	4455>  "  G  4
	4455>  "  G  4
	4455>  "  G  4>   "  G  		 C=      g  filenamef  auld-lang-syne.scm
	
							#			/			0			@			R			b			e			j			s			v			{		 		 		 		 		 		 		 ¦		 Ά		 Ζ	 	 Φ	!	 Ω	!	 ή	!	 η	"	 κ	"	 ο	"	 ψ	#	 ϋ	#	 	#			$		$		$		&		&	!	&	+	&	0	&	9	'	>	'	A	'	K	'	P	'	Y	(	^	(	a	(	k	(	p	(	y	)	~	)		)		)		)		+	―	-	 :	°
  g  nameg  new-game CR         h   t   ] 45 6     l       g  filenamef  auld-lang-syne.scm
	0
		1			1	 		
  g  nameg  give-status-message CR!"#$%&'       h    ΄   ] 45444
5556 ¬       g  filenamef  auld-lang-syne.scm
	3
		4				4			4			4	"		5			5	!		5	)		5	!		5			4	 		
  g  nameg  get-stock-no-string C R(   h   €   ]4 5$  C 	C          g  slot-id
		 g  	card-list		  g  filenamef  auld-lang-syne.scm
	7
		8			8			9	 			  g  nameg  button-pressed C)R*+ h0     ]	$  !
$  454455CCC  ω       g  
start-slot
		. g  	card-list		. g  end-slot			.  g  filenamef  auld-lang-syne.scm
	;
		<			<			=			<			>	
		>			>	
		?			?		'	?		(	?	
	)	>	 		.	  g  nameg  
droppable? C,R,-.  h0   Ξ   ]4 5$  4 5$  6CC       Ζ       g  
start-slot
		) g  	card-list		) g  end-slot			)  g  filenamef  auld-lang-syne.scm
	A
		B			B			C		 	B		%	D	 		)	  g  nameg  button-released C/R( h   j   ] 4
5C      b       g  filenamef  auld-lang-syne.scm
	F
		G				G	 		

  g  nameg  	dealable? C0R12   h   s   ] 
6       k       g  filenamef  auld-lang-syne.scm
	I
		J				J	 			
  g  nameg  do-deal-next-cards C3R(12        h    ₯   ] 
$  4
5$  C
6C        g  slot-id
		  g  filenamef  auld-lang-syne.scm
	L
		M		
	M			N			M			O			O	 		  g  nameg  button-clicked C4R(*+5.6        h`     ]4 5$  "  44 554455$  4  >  "  G  6	$  	 6C      g  slot1
		` g  slot2		`  g  filenamef  auld-lang-syne.scm
	Q
		R			R			S			S			S			T		"	T		*	T		+	T		,	S		0	R		1	V		:	V		?	V		L	W		Q	X	
	U	X		\	Y	!	^	Y	
 		`	  g  nameg  check-end-slot? C6R6        h      ] 	$   6C           g  slot-id
		  g  filenamef  auld-lang-syne.scm
	\
		]			]			^	 		  g  nameg  button-double-clicked C7R89 h   y   ] 45 $  C6        q       g  filenamef  auld-lang-syne.scm
	`
		a			a			b	 		
  g  nameg  game-continuable C:R(    hP   ©   ] 4>   "  G  4
5$  -4	5$   4	5$  4	5$  	6CCCC     ‘       g  filenamef  auld-lang-syne.scm
	d
		e			f			f			g		'	f		(	h		2	f		3	i		=	f		C	j	 		K
  g  nameg  game-won C8R(;*+<      hh   /  ]
 		$  C	$  "  4 5$   644 554455$  	 6 6   '      g  slot1
		e g  slot2		e g  t			)  g  filenamef  auld-lang-syne.scm
	l
		m			m			o			o	
	 	p		-	o		2	q		5	q	
	6	r		9	r		A	r		B	s		E	s	!	M	s		N	s		O	r		S	r	
	\	t		c	u		e	u	 		e	  g  nameg  movable? C;R("=    h       ] 4
5$  C
45 C             g  filenamef  auld-lang-syne.scm
	w
		x			x			y			y			y			y	 		
  g  nameg  	dealable? C0R;0   h       ]4	5  $   C6        y       g  t
	
	  g  filenamef  auld-lang-syne.scm
	{
		|		
	|			}	 		
  g  nameg  get-hint C9R  h   \   ] C    T       g  filenamef  auld-lang-syne.scm
	
 		
  g  nameg  get-options C>R       h   u   ]C    m       g  options
		  g  filenamef  auld-lang-syne.scm
 
 		  g  nameg  apply-options C?R      h   Y   ] C    Q       g  filenamef  auld-lang-syne.scm
 
 		
  g  nameg  timeout C@R4AiBiCi>  "  G  Dii)i/i4i7i:i8i9i>i?i@i,i0i6κ       g  filenamef  auld-lang-syne.scm		
@	
ί	0
Χ	3
§	7
μ	;
		A
		F

&	I
	L
ͺ	Q
o	\
	`
$	d
Ϊ	l
	w
T	{
Η	
X 
Μ 
Ν 
  
 	 
   C6 