GOOF----LE-8-2.08      ] f 4        h+      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  stock¤								 ¤	g  
foundation¤	g  waste¤		
									 
¤	g  tableau¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-double-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  add-extended-slot¤	g  right¤	g  down¤	g  deal-cards-face-up¤	g  give-status-message¤	 e  4.5¤	!g  new-game¤	"g  in-tableau?¤	#g  in-foundation?¤	$g  
waste-pile¤	%g  in-tableau-or-waste?¤	&g  
stock-pile¤	'g  start-with-waste¤	(g  start-with-tableau¤	)g  <>¤	*g  set-statusbar-message¤	+g  get-stock-no-string¤	,g  string-append¤	-g  _¤	.f  Stock left:¤	/f   ¤	0g  number->string¤	1g  length¤	2g  	get-cards¤	3g  empty-slot?¤	4g  check-straight-descending-list¤	5g  check-same-suit-list¤	6g  button-pressed¤	7g  foundation-score¤	8g  
set-score!¤	9g  recalculate-score¤	:g  space-score¤	;g  tableau-spaces¤	<g  	get-value¤	=g  ace¤	>g  get-top-card¤	?g  get-suit¤	@g  foundation-droppable?¤	Ag  expt¤	Bg  max¤	Cg  max-move-in-tableau¤	Dg  tableau-droppable?¤	Eg  
droppable?¤	Fg  move-n-cards!¤	Gg  reverse¤	Hg  button-released¤	Ig  try-all-foundations-helper¤	Jg  try-all-foundations¤	Kg  find-tableau-place-helper¤	Lg  find-empty-slot¤	Mg  find-tableau-place¤	Ng  	dealable?¤	Og  button-clicked¤	Pg  do-deal-next-cards¤	Qg  find-any-move-to-foundation¤	Rg  move-to-foundation¤	Sg  find-any-move-in-tableau¤	Tg  delayed-call¤	Ug  autoplay-foundations¤	Vg  button-double-clicked¤	Wg  game-won¤	Xg  get-hint¤	Yg  game-continuable¤	Zf  Deal a card from stock¤	[g  check-for-deal¤	\g  	hint-move¤	]f  $Bug! make-hint called on false move.¤	^g  	make-hint¤	_g  get-options¤	`g  apply-options¤	ag  timeout¤	bg  set-features¤	cg  droppable-feature¤	dg  dealable-feature¤	eg  
set-lambda¤C 5hx-  ç  ] 4	 >  "  G  
RR		RR   hp    ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4
>   "  G  4>  "  G  4
>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4>   "  G  	
 C             g  filenamef  forty-thieves.scm
	
							#			3			C			I			N			W			g	!		j	!		l	!		q	!		z	"		}	"			"	 	"	 	#	 	#	 	#	 	#	  	$	 £	$	 ¥	$	 ª	$	 ³	%	 ¶	%	 ¸	%	 ½	%	 Æ	&	 É	&	 Ë	&	 Ð	&	 Ù	'	 Ü	'	 Þ	'	 ã	'	 ì	(	 ï	(	 ñ	(	 ö	(	 ÿ	*		,		,		,		,	$	-	4	0	7	0	;	0	@	0	I	1	L	1	P	1	U	1	^	2	a	2	e	2	j	2	s	3	v	3	z	3		3		4		4		4		4		5	 	5	¤	5	©	5	²	6	µ	6	¹	6	¾	6	Ç	7	Ê	7	Î	7	Ó	7	Ü	8	ß	8	ã	8	è	8	ñ	9	ô	9	ø	9	ý	9		<		=	,	>	?	?	R	A	e	D	h	D	 _	i
  g  nameg  new-game C!R    h      ] 	
$   	CC     ~       g  slot
		  g  filenamef  forty-thieves.scm
	G
		H			H			H	 		  g  nameg  in-tableau? C"R     h      ] $   	CC             g  slot
		  g  filenamef  forty-thieves.scm
	K
		L		
	L			L	 		  g  nameg  in-foundation? C#R"$    h       ]	4 5$  C C             g  slot
		 g  t			  g  filenamef  forty-thieves.scm
	O
		P				P			P	 		  g  nameg  in-tableau-or-waste? C%R		$R
&R		'R	
(R   h      ] C       y       g  a
			 g  b			  g  filenamef  forty-thieves.scm
	X
		Y			Y	 				  g  nameg  <> C)R*+    h   s   ] 45 6     k       g  filenamef  forty-thieves.scm
	\
		]			]	 		
  g  nameg  give-status-message CR,-./012        h    ³   ] 45444
5556 «       g  filenamef  forty-thieves.scm
	_
		`				`			`			`	"		a			a	!		a	)		a	!		a			`	 		
  g  nameg  get-stock-no-string C+R3%1"45  hP   ú   ]
4 5$  C4 5$  445$  C4 5$  45$  6CCCò       g  slot-id
		P g  	card-list		P g  t		#	N  g  filenamef  forty-thieves.scm
	j		k			k			l			k			m		#	m		#	m		/	n		9	n		:	o		D	n		J	p	 		P	  g  nameg  button-pressed C6R127 hp   P  ] 	$  .44 55	44 55	$  	<"  
C 44 55	44 55	$  	<"  
6   H      g  slot-id
		m g  
prev-total		m  g  filenamef  forty-thieves.scm
	w
		~			~			z	
		z			z	
		z			{		 	{		(	{		+	{		/	{		7	y		> 		A	z	
	D	z		L	z	
	O	z		P	{		S	{		[	{		^	{		b	{		j	y		m 	 		m	  g  nameg  foundation-score C7R87     h   t   ] 4
56   l       g  filenamef  forty-thieves.scm
 
	 		 	 		
  g  nameg  recalculate-score C9R3: h@   Ý   ] 	$  4 5$  "  
C 4 5$  "  
6     Õ       g  slot-id
		; g  prev		;  g  filenamef  forty-thieves.scm
 
	 		 	 	+	 	'	 		% 	/	( 	+	2 	'	9 		; 	" 		;	  g  nameg  space-score C:R:(        h   h   ] 
6       `       g  filenamef  forty-thieves.scm
 
		 	 			
  g  nameg  tableau-spaces C;R543<=>?   hh   |  ]4 5$  V4 5$  I45$  4 5C4 54455$  44554 5CCCC   t      g  	card-list
		e g  f-slot		e  g  filenamef  forty-thieves.scm
 
	 		 		 		 		 		# 		$ 		) 	"	+ 		. 		0 		5 		7 		8 	3	; 	>	C 	3	D 	0	E 		I 		J 		M 	"	U 		V 	9	[ 	C	] 	9	^ 	 		e	  g  nameg  foundation-droppable? C@RAB;3     h8   ë   ]	4
45 45$  "  
4 5$  "  
56 ã       g  	from-slot
		7 g  to-slot		7  g  filenamef  forty-thieves.scm
 ¥
	 ¦		 ¨		 ¨	,	 ¨	(	! ¨		" ©		, ©		3 §		5 ¦		7 ¦	 		7	  g  nameg  max-move-in-tableau CCR541C3<>? 	     h     ]45$  j45$  ]454 5$  F45$  C45454455$  445545CCCCC             g  s-slot
		y g  	card-list		y g  t-slot			y  g  filenamef  forty-thieves.scm
 ³
	 µ		 ´		 ¶		 ´		 ·		  ·		) ·		- ´		. ¸		8 ¸		; ¹		@ ¹	"	B ¹		C ¹	3	J ¹		K ¹	G	N ¹	R	V ¹	G	W ¹		[ ¸		\ º		_ º	"	g º		h º	9	m º	C	o º	9	p º	 		y	  g  nameg  tableau-droppable? CDR#@"D   hH   û   ] $  C4 5$  C45$  645$  
 6C    ó       g  
start-slot
		D g  	card-list		D g  end-slot			D  g  filenamef  forty-thieves.scm
 Ê
	 Ë	
	 Ë		 Ì	
	 Ë		 Í	
	% Ë		- Í	$	. Î	
	8 Ë		B Î	! 		D	  g  nameg  
droppable? CERE"FG      h8   å   ]4 5$  $45$  
 6 456C Ý       g  
start-slot
		7 g  	card-list		7 g  end-slot			7  g  filenamef  forty-thieves.scm
 Õ
	 Ö		 Ö		 ×		 ×		& Ø		- Ù	0	5 Ù	 			7	  g  nameg  button-released CHR@I        h8     ](   C4 5$  
  C 6           g  	from-slot
		3 g  card		3 g  to-slots			3  g  filenamef  forty-thieves.scm
 ã
	 ä		 å		 æ	
	 æ	!	 æ	-	 æ	
	 æ		$ ç		' ç	
	1 è	5	3 è	
 		3	  g  nameg  try-all-foundations-helper CIR3I     h    ²   ]4 5$   C 6    ª       g  	from-slot
		 g  card		  g  filenamef  forty-thieves.scm
 ê
	 ë		 ë		 í		 ì	 			  g  nameg  try-all-foundations CJR3D)K       hX   L  ](  C45$  "  "4  5$  4 5"  $  
  C 6  D      g  	from-slot
		V g  card		V g  to-slots			V  g  filenamef  forty-thieves.scm
 ó
	 ô		 ÷		 ÷		 ÷		 ö	
	 ø		% ø	*	( ø	6	* ø		. ö	
	/ ù		6 ù		8 ù		A ö		G ú		J ú	
	T û	4	V û	
 		V	  g  nameg  find-tableau-place-helper CKR3KL     hX     ]
4 5$   C4 5$  C45$   45 "  $  C C         g  	from-slot
		U g  card		U g  t			U g  t		E	U  g  filenamef  forty-thieves.scm
 ý
	 þ		 þ				 		 ÿ		)		3		7	:	@	'	E ÿ		T	 		U	  g  nameg  find-tableau-place CMR3&       h   l   ] 45C     d       g  filenamef  forty-thieves.scm

			
	 		
  g  nameg  	dealable? CNR&N$     h(   «   ] $  45 $   6CC     £       g  slot-id
		#  g  filenamef  forty-thieves.scm


										&		 		#  g  nameg  button-clicked CORO&  h   l   ] 6d       g  filenamef  forty-thieves.scm

		 		
  g  nameg  do-deal-next-cards CPRQ$G     h(   Ð   ]45  $   4 5 6C È       g  move
			'  g  filenamef  forty-thieves.scm

									
		2		-		I	 	D	#	>	%	 		'
  g  nameg  move-to-foundation CRR%J>Q h8   ù   ]	4 5$  !4 4 55$  C 6 C     ñ       g  
begin-slot
		3 g  test		.  g  filenamef  forty-thieves.scm

								4					  		$ 		,"	.	."		2%	 		3  g  nameg  find-any-move-to-foundation CQR%M>S        h8   ö   ]	4 5$  !4 4 55$  C 6 C     î       g  
begin-slot
		3 g  test		.  g  filenamef  forty-thieves.scm
+
	,		,		-		-	3	-		-		 .		$.		,0	+	.0		23	 		3  g  nameg  find-any-move-in-tableau CSRRTU     h      ] 45 $  6C     w       g  filenamef  forty-thieves.scm
9
	:		:
	:	 		
  g  nameg  autoplay-foundations CUR#U%J>GM 	 hx   ¤  ]4 5$  6 4 5$  V4 4 55$  45 64 4 55$  45 6CC          g  slot-id
		t g  test	*	r g  jump		T	r  g  filenamef  forty-thieves.scm
?
	@	
	@		@	$	A	
	@		B		"B	5	*B		*B		/C		3C		8D	*	9D	%	:D	A	AD	<	DD	6	FD		GE		LE	9	TE		TE		YF		]F		bG	0	cG	+	dG	G	kG	B	nG	<	pG	 		t  g  nameg  button-double-clicked CVRWX       h(      ] 4>   "  G  45 $  C6        }       g  filenamef  forty-thieves.scm
S
	T		U		U		!V	 		!
  g  nameg  game-continuable CYR9  h   k   ] 45 èC    c       g  filenamef  forty-thieves.scm
\
	]		]	 		
  g  nameg  game-won CWR3&-Z      h       ] 45$  C
45 C            g  filenamef  forty-thieves.scm
b
	c		c		d		d		d		d	 		
  g  nameg  check-for-deal C[R\G]      h(   ¾   ] $   4 56
 C       ¶       g  move
		!  g  filenamef  forty-thieves.scm
j
	k			k		l		l		l	*	l	%	l		m		 m	 		!  g  nameg  	make-hint C^RQ'^3$M>S([       hp     ] 45$  45645$  "  4455$  445564	5$  4	56
6              g  filenamef  forty-thieves.scm
z
	{		
{	
	{		|		|	
	~		$~	
	*		/	3	7		8		<{		?		D	4	L		N	
	O		V	
	Z{		]		e	
	i	 		i
  g  nameg  get-hint CXR    h   \   ] C    T       g  filenamef  forty-thieves.scm

 		
  g  nameg  get-options C_R       h   t   ]C    l       g  options
		  g  filenamef  forty-thieves.scm

 		  g  nameg  apply-options C`R       h   X   ] C    P       g  filenamef  forty-thieves.scm

 		
  g  nameg  timeout CaR4bicidi>  "  G  ei!i6iHiOiViYiWiXi_i`iaiEiNi6 ß      g  filenamef  forty-thieves.scm		
		
				"	
	'	
	)			,	
b	
	G
Ä	K
	O
	S
¡	T
¦	U
«	V
	L	X
	æ	\

Þ	_
E	j	w
· 
è 
{ 
 
¾ ¥
ù ³
V Ê
 Õ
ñ ã
Ý ê
 ó
! ý
·
¦

 /
!C
"
#Ù+
$9
&Ç?
'S
(\
(ðb
)ñj
+ªz
,
,¯
-#
-$
-w
 3	-w
   C6 