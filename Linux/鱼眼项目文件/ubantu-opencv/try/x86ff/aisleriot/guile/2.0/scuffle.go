GOOF----LE-8-2.0°      ] I 4     hD      ] g  guile€	 €	g  process-use-modules€	 €	 €	g  	aisleriot€	g  	interface€	 €		 €	
g  api€	
 €	 €	g  initialize-playing-area€	g  set-ace-low€	g  make-deck-list-ace-low€	g  club€	g  DECK€	g  shuffle-deck€	g  add-normal-slot€	g  add-blank-slot€	g  add-carriage-return-slot€	g  	add-card!€	g  make-visible€	g  	make-card€	g  ace€	g  diamond€	g  heart€	g  spade€	g  give-status-message€	g  new-game€	g  set-statusbar-message€	 g  string-append€	!g  get-stock-no-string€	"f     €	#g  get-redeals-string€	$g  _€	%f  Redeals left:€	&f   €	'g  number->string€	(g  FLIP-COUNTER€	)f  Stock left:€	*g  length€	+g  	get-cards€	,g  empty-slot?€	-g  button-pressed€	.g  	get-value€	/g  get-top-card€	0g  
droppable?€	1g  move-n-cards!€	2g  add-to-score!€	3g  button-released€	4g  deal-cards-face-up€	5g  deal-cards-out€	6g  	flip-deck€	7g  button-clicked€	8g  
deal-cards€	9g  check-end-slot?€	:g  button-double-clicked€	;g  game-won€	<g  get-hint€	=g  game-continuable€	>g  movable?€	?g  	hint-move€	@f  Deal another round€	Af  Reshuffle cards€	Bg  	dealable?€	Cg  get-options€	Dg  apply-options€	Eg  timeout€	Fg  set-features€	Gg  droppable-feature€	Hg  
set-lambda€C 5       h  ξ   ] 4	 >  "  G   h°  >  ] 4>   "  G  4>   "  G  4		5 4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4	>   "  G  4>   "  G  4>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
4455>  "  G  4
	4455>  "  G  4
	4455>  "  G  4
	4455>  "  G  4>   "  G  		 C6      g  filenamef  scuffle.scm
	
							#			/			0			@			R			b			e			j			s			v			{		 		 		 		 		 		 		 ¦		 Ά		 Ζ	 	 Φ	!	 Ω	!	 ή	!	 η	"	 κ	"	 ο	"	 ψ	#	 ϋ	#	 	#			$		$		$		&		&	!	&	+	&	0	&	9	'	>	'	A	'	K	'	P	'	Y	(	^	(	a	(	k	(	p	(	y	)	~	)		)		)		)		+	―	-	 :	°
  g  nameg  new-game CR !"# h      ] 445 45 56        g  filenamef  scuffle.scm
	0
		1			1	(		2	(		3	(		1			1	 		
  g  nameg  give-status-message CR $%&'(        h       ] 454	56              g  filenamef  scuffle.scm
	5
		6				6			6			6	$		7			7	!		7			6	 
		
  g  nameg  get-redeals-string C#R $)&'*+       h    ­   ] 45444
5556 ₯       g  filenamef  scuffle.scm
	9
		:				:			:			:	"		;			;	!		;	)		;	!		;			:	 		
  g  nameg  get-stock-no-string C!R,  h      ]4 5$  C 	C          g  slot-id
		 g  	card-list		  g  filenamef  scuffle.scm
	=
		>			>			?	 			  g  nameg  button-pressed C-R./        h0   ϊ   ]	$  !
$  454455CCC  ς       g  
start-slot
		. g  	card-list		. g  end-slot			.  g  filenamef  scuffle.scm
	A
		B			B			C			B			D	
		D			D	
		E			E		'	E		(	E	
	)	D	 		.	  g  nameg  
droppable? C0R012 h0   Η   ]4 5$  4 5$  6CC       Ώ       g  
start-slot
		) g  	card-list		) g  end-slot			)  g  filenamef  scuffle.scm
	H
		I			I			J		 	I		%	K	 		)	  g  nameg  button-released C3R,45    h8   »   ] 		$  $4
5$  C4
  5$   6CC       ³       g  slot
		1  g  filenamef  scuffle.scm
	M
		N			N			O			N			P		 	P	!	"	P		&	P		+	Q		-	Q	 		1  g  nameg  deal-cards-out C5R,5(6    h      ]	 
$  4
5$  "  4	5$  C	$  \ $  P4
	>  "  G  4
	>  "  G  4
	>  "  G  4
	>  "  G  6 CCC    ψ       g  slot-id
	  g  t	    g  filenamef  scuffle.scm
	T
		U		
	U			V			V			W		 	V		0	X		4	X		7	Y	#	9	Y		>	X		?	[		R	\		e	]		x	^	 	_	 	   g  nameg  button-clicked C7R,./829     h`     ]4 5$  "  44 554455$  4  >  "  G  6	$  	 6C      g  slot1
		` g  slot2		`  g  filenamef  scuffle.scm
	a
		b			b			c			c			c			d		"	d		*	d		+	d		,	c		0	b		1	f		:	f		?	f		L	g		Q	h	
	U	h		\	i	!	^	i	
 		`	  g  nameg  check-end-slot? C9R9       h      ] 	$   6C           g  slot-id
		  g  filenamef  scuffle.scm
	l
		m			m			n	 		  g  nameg  button-double-clicked C:R;<      h(   z   ] 4>   "  G  45 $  C6        r       g  filenamef  scuffle.scm
	p
		q			r			r		!	s	 		!
  g  nameg  game-continuable C=R,     h@      ] 4
5$  -4	5$   4	5$  4	5$  	6CCCC            g  filenamef  scuffle.scm
	u
		v			v			w			v			x		"	v		#	y		-	v		3	z	 		;
  g  nameg  game-won C;R,>./?     hh   7  ]
 		$  C	$  "  4 5$   644 554455$  	 6 6   /      g  slot1
		e g  slot2		e g  t			)  g  filenamef  scuffle.scm
	|
		}			}							
	  		-			2 		5 	
	6 		9 		A 		B 		E 	!	M 		N 		O 		S 	
	\ 		c 		e 	 		e	  g  nameg  movable? C>R,$@(A        h@   Ω   ]4
5$  "  
45   $   C	$  
45 CC  Ρ       g  t
		>  g  filenamef  scuffle.scm
 
	 		 		 		 		 		 		 		, 		0 		2 		6 		8 		; 	 		>
  g  nameg  	dealable? CBR>B    h    ~   ]4	5  $   C6        v       g  t
	
	  g  filenamef  scuffle.scm
 
	 		
 		 	 		
  g  nameg  get-hint C<R     h   V   ] C    N       g  filenamef  scuffle.scm
 
 		
  g  nameg  get-options CCR     h   n   ]C    f       g  options
		  g  filenamef  scuffle.scm
 
 		  g  nameg  apply-options CDR     h   R   ] C    J       g  filenamef  scuffle.scm
 
 		
  g  nameg  timeout CER4FiGi>  "  G  Hii-i3i7i:i=i;i<iCiDiEi0i6       ζ       g  filenamef  scuffle.scm		
9	
ψ	0
Χ	5
Θ	9
	=
Υ	A
	β	H

ξ	M
£	T
;	a
ψ	l
΅	p
₯	u
b	|
 
Q 
Α 
I 
΅ 
Ά 
 
 	
   C6 