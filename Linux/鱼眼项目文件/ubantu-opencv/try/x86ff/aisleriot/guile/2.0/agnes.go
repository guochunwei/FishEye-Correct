GOOF----LE-8-2.087      ] m 4 h!      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	g  ice-9¤	g  format¤	 ¤	 ¤	g  BASE-VAL¤	g  stock¤				 ¤	g  
foundation¤								
	 ¤	g  tableau¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-normal-slot¤	g  DECK¤	g  add-blank-slot¤	g  add-carriage-return-slot¤	g  add-extended-slot¤	 g  down¤	!g  
deal-cards¤	"							
							
						
					
				
		
		 ¤	#g  map¤	$g  flip-top-card¤	%g  deal-cards-face-up¤	& ¤	'g  add-to-score!¤	(g  	get-value¤	)g  get-top-card¤	*g  give-status-message¤	+g  dealable-set-sensitive¤	,g  	dealable?¤	-g  new-game¤	.g  set-statusbar-message¤	/g  string-append¤	0g  get-stock-no-string¤	1f     ¤	2g  get-base-string¤	3g  _¤	4f  Base Card: Ace¤	5f  Base Card: Jack¤	6f  Base Card: Queen¤	7f  Base Card: King¤	8f   ¤	9f  Base Card: ~a¤	:g  number->string¤	;g  length¤	<g  	get-cards¤	=f  Stock left:¤	>f   ¤	?f  Stock left: 0¤	@g  king¤	Ag  ace¤	Bg  check-straight-descending-list¤	Cg  empty-slot?¤	Dg  is-visible?¤	Eg  reverse¤	Fg  check-same-color-list¤	Gg  button-pressed¤	Hg  is-red?¤	Ig  get-suit¤	Jg  
droppable?¤	Kg  move-n-cards!¤	Lg  make-visible-top-card¤	Mg  button-released¤	Ng  check-slot-and-deal¤	Og  do-deal-next-cards¤	Pg  button-clicked¤	Qg  check-dc¤	Rg  or-map¤	Sg  button-double-clicked¤	Tg  delayed-call¤	Ug  autoplay-foundations¤	V	 ¤	W	 ¤	X	 ¤	Yg  game-won¤	Zg  game-continuable¤	[g  	hint-move¤	\g  find-empty-slot¤	]g  check-to-foundation?¤	^g  check-a-tableau¤	_g  strip¤	`g  check-to-tableau?¤	ag  	find-card¤	bf  Deal more cards¤	cg  check-deal?¤	df  Try rearranging the cards¤	eg  get-hint¤	fg  get-options¤	gg  apply-options¤	hg  timeout¤	ig  set-features¤	jg  droppable-feature¤	kg  dealable-feature¤	lg  
set-lambda¤C 5  h(-  ^  ] 4	 >  "  G  
R
RRR !"#$%&'()*+, hð  Ö  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>   "  G  4	>  "  G  4	>  "  G  4	>  "  G  4	>  "  G  4
>   "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4>  "  G  4
>  "  G  4>  "  G  4
>  "  G  4>  "  G  4455 4>   "  G  445 >  "  G  		 C      Î      g  filenamef  	agnes.scm
	
							#			3			C			I			N			W			g	 		w	"		z	"		|	"	 	"	 	#	 	#	 	#	 	#	 	$	  	$	 ¢	$	 §	$	 °	%	 ³	%	 µ	%	 º	%	 Ã	&	 Ó	(	 Ö	(	 Ú	(	 ß	(	 è	)	 ë	)	 ï	)	 ô	)	 ý	*	 	*		*			*		+		+		+		+	'	,	*	,	.	,	3	,	<	-	?	-	C	-	H	-	Q	.	T	.	X	.	]	.	f	0	k	0	p	0	y	3		3		3		5		5		5	 	7	±	8	´	8	»	8	½	8	¾	:	Î	;	Ñ	;	Ú	;	é	=	 K	ê
  g  nameg  new-game C-R./012 h      ] 445 45 56        g  filenamef  	agnes.scm
	?
		@			@	(		A	(		B	(		@			@	 		
  g  nameg  give-status-message C*R3456789:  hp   /  ] "  >$  6	$  6	$  6	$  6C$   	$  4	54
56"ÿÿ"ÿÿ'      g  filenamef  	agnes.scm
	D
	
	H				E			I			I				J				E		!	K		#	K			(	L			,	E		0	M		2	M			7	N			;	E		?	O		A	O			C	P		D	E		H	E		L	E		Q	F		U	E			Y	G		]	G		_	G		`	G	'	h	G		 		p
  g  nameg  get-base-string C2R;</3=>:? 	 h@   ó   ] 44
55$  45444
5556456     ë       g  filenamef  	agnes.scm
	R
		S				S			S				S			S			T			T			T			T	&		U		"	U	%	%	U	-	,	U	%	.	U		0	T		3	V		7	V		9	V		;	V	 		;
  g  nameg  get-stock-no-string C0R;(@AB  h   ¼  ]	4 5	$  C4 5$  /4 5$  $  "  4 5"  "  $  C4 54 5$  4 5$  C 6C     ´      g  	card-list
	  g  t	  g  t	V   g  filenamef  	agnes.scm
	X
		Y				Y			Y			Z			Z			Z		"	Z		&	Z		'	[		,	[		/	[		2	[		6	Z		;	\		?	Z		E	]		J	]	+	L	]		V	Y		b	^		g	^		i	^		j	^	-	o	^	8	r	^	-	s	^	*	t	^		x	^		{	_	 	_	' 	_	 	_	 	^	 	`	+ 	`	 %	   g  nameg  check-straight-descending-list CBRCDEFB   h8   Ó   ]4 5$  C4455$  45$  6CC   Ë       g  slot-id
		5 g  	card-list		5  g  filenamef  	agnes.scm
	b
		c			c			d			d			d			d		 	c		!	e		+	c		1	f	 		5	  g  nameg  button-pressed CGRCH)(E@A;I  hH  [  ] $  C"  ~	$  s45$  C454455&  L44554455$  C4455$  4455CCCC
$  ¬	$  45$  45$  45	"  $  C45$  C4
54
455$  B454455$  C45$  4455CCCC"ÿþÎ"ÿþÊ S      g  
start-slot
	G g  	card-list	G g  end-slot		G g  t		Z  g  t	 È= g  t	;  g  filenamef  	agnes.scm
	h
		i				i			x				i			y		%	y			(	z		-	z		/	z		0	{		3	{		;	{		?	y			@	|		C	|	%	J	|	 	L	|		M	}		P	}	#	X	}		Y	}		Z	|		Z	|		f	~		i	~	*	p	~	%	r	~		u	~		y	~		z			}		% 		 		 	i	 	k	 	i	 	l	  	k		 ¡	m	 ©	m	 ­	m		 ®	n	 ¸	n	 ¹	o	 ¾	o	% À	o	 Ã	o	 È	n	 Ô	q	 Þ	q	 á	r	 æ	r	$ è	r	 é	s	 ì	s	$ ô	s	 õ	r	 ù	q	 ú	t	 ÿ	t	)	t		u	#	u	.	u	#	u		t		t		v	# 	v	."	v	#%	v	 )	v	*	w	#-	w	.5	w	#8	w	  N	G	  g  nameg  
droppable? CJRJK'CL    h   R  ]4 5$  l4 5$  [ 	$  "  4	ÿ5$  9	$  "  45$  4 5$  C 6CCCC J      g  
start-slot
		 g  	card-list		 g  end-slot			 g  t		%	< g  t		E	[ g  t		f	w  g  filenamef  	agnes.scm
 
	 		 		 		  		% 		% 		3 		@ 		E 		E 		S 		_ 		` 		f 		w 	 			  g  nameg  button-released CMRC%N h8   É   ]4
5$  "   	$  4
  5$   6CC    Á       g  slot
		4  g  filenamef  	agnes.scm
 
	 		 		 		 		 		# 	!	% 		) 		. 	 	0 	 		4  g  nameg  check-slot-and-deal CNR,N    h   v   ] 45 $  	6C     n       g  filenamef  	agnes.scm
 
	 		 		 	 		
  g  nameg  do-deal-next-cards CORO h      ] 
$  6 C       g  slot-id
		  g  filenamef  	agnes.scm
 
	 		
 		 	 		  g  nameg  button-clicked CPRC       h   d   ] 4
5C      \       g  filenamef  	agnes.scm
 
	 			 	 		

  g  nameg  	dealable? C,RCI)(A@!'LQ h  z  ]	$  C45$  "  á44 554455$  Ã44 554455$  "  '44 55$  4455"  $  i$  "  $  $  "  "  C4  5$  245$  #4 5$  "  4	 5"  "  "  "  $  C
 6    r      g  slot
	 g  f-slot	 g  just-checking?		 g  t		T  g  t	 Ò ê g  t	 þ  g  filenamef  	agnes.scm
 
	 			 		  		  			 ¡		! ¡		) ¡		* ¢		- ¢		5 ¢		6 ¡		:  			; £		> £	 	F £		G ¤		J ¤	%	R ¤		S ¤		T £		T £		b ¥		e ¥	%	m ¥		p ¥		t ¥		u ¦		x ¦	%  ¦	  ¦	   		  §	   §	 ¦ §	 ² ©	 » ©	( ½ ©	 Á ©	 Â ª	 Ë ©	 Ì «	 Ò «	 á ¬	 þ 	 ®	 ®		 0		  g  nameg  check-dc CQRRST      h      ] 45$  L 6C        g  filenamef  	agnes.scm
 ±		 ²			 ²	&	 ²		 ²		 ³	 		
  g  nameg  autoplay-foundations-tail CRST     h8   Ö   ]O   Q  45$  45$   6CC     Î       g  autoplay-foundations-tail
	
	3  g  filenamef  	agnes.scm
 °
	 µ		 µ	$	 µ		 µ		 ²		# ²	&	% ²		) ²		/ ³	 		3
  g  nameg  autoplay-foundations CURCU()!&VWX'LQ h  T  ]4 5$  	 	"  $  "   
$  C 	$  6 44 55$  ­45$  4 5"  $  "  [4	5$  4 5"  $  "  24	5$  4 	5"  $  "  	4 
5$  $45$  4 5$  C 6CC 6  L      g  slot-id
	 g  t		+ g  t	i Ô g  t	  Ñ g  t	 µ Î g  t é ú  g  filenamef  	agnes.scm
 ¹
	 º		 º		 »		 º			( ¼		/ º		6 ¾			: º		> ¿			? À		B À		J À		M À			Q º		R Á		[ Á		\ Â		b Â	+	d Â		i Á		w Ã	  Ã	  Ä	  Ä	+  Ä	  Á	  Å	 § Å	 ¨ Æ	 ® Æ	+ ° Æ	 µ Á	 Ã Ç	 É Ç	& Ë Ç	 Ø Á		 Ù È	 â Á		 ã É	 é É	 ú Ê	 Ì		 ,	  g  nameg  button-double-clicked CSR*+,Y     h0      ] 4>   "  G  445 >  "  G  45 C         g  filenamef  	agnes.scm
 Î
	 Ï		 Ð		 Ð		 Ð		( Ñ		- Ñ	 		.
  g  nameg  game-continuable CZR;<      hP   ü   ] 	4455$  :	44	55$  %	44	55$  	44	55CCCCô       g  filenamef  	agnes.scm
 Ó
	 Ô		 Ô		 Ô		 Ô		 Ô		 Õ		 Õ		" Õ		# Õ		' Ô		* Ö		- Ö		5 Ö		6 Ö		: Ô		= ×		@ ×		H ×		I ×	 		P
  g  nameg  game-won CYRC()[\Q] 
  hx     ] 	$  C4 5$  "  44 55$   4564 5$  "  	4 5$   4 56	 6            g  slot
		r  g  filenamef  	agnes.scm
 Ù
	 Ú			 Ú		 Ü		 Ü			 Ý		! Ý		) Ý		, Ý		0 Ú		6 ß		> ß			? à		I à			O á		[ Ú		a â		k â			p ã	"	r ã	 		r  g  nameg  check-to-foundation? C]RCH)(@A        h   y  ]
45$  C4 54455&  X4455$  C4 5$  4455"  $  C4 54455CCq      g  card
	  g  slot	  g  t		]	~  g  filenamef  	agnes.scm
 å
	 æ		 æ		 ç		 ç		 ç	$	" ç		& æ		' è		* è		2 è		5 è		9 æ		< é		E é		I é		J ê		M ê		U ê		X ê		] é		i ë		p ë		q ì		t ì		| ì		} ë	 	 	  g  nameg  check-a-tableau C^R;DEFB_    hp   R  ]4 5	$   C44 55$  "   4 5$  "  4 5$  44 5564 5CJ      g  	card-list
		p g  t	"	R g  t		7	O  g  filenamef  	agnes.scm
 î
	 ï		 ï			 ï		 ð			 ñ		 ñ	$	 ñ		! ñ		" ñ		" ñ			0 ô		7 ô		7 ñ			E õ		L õ		V ï		Y ö		\ ö		c ö		e ö		g ö			h ÷		o ÷	 		p  g  nameg  strip C_RC`^_<[a h   h  ]
 	$  C	$  "  4 5$  	 	6 $  "  444 555$   4 44 5556 6 `      g  slot1
		 g  slot2		 g  t			)  g  filenamef  	agnes.scm
 ù
	 ú			 ú		 ü		 ü			  ý		- ú		2 þ		6 þ			; ÿ		? ÿ			E 		H 		K 	&	S 		W 		[ ú		`		e	+	h	2	p	+	r		v			}	%		 			  g  nameg  check-to-tableau? C`R,3b   h       ] 45 $  
45 CC              g  filenamef  	agnes.scm

												 		
  g  nameg  check-deal? CcR]`c3d     hH   ×   ]4	5  $   C4		5  $   C45   $   C
45 C    Ï       g  t
			D g  t
		D g  t
	-	D  g  filenamef  	agnes.scm
	
	
			
				
		)		-
		:		>		@		C	 		D
  g  nameg  get-hint CeR    h   T   ] C    L       g  filenamef  	agnes.scm

 		
  g  nameg  get-options CfR       h   l   ]C    d       g  options
		  g  filenamef  	agnes.scm

 		  g  nameg  apply-options CgR       h   P   ] C    H       g  filenamef  	agnes.scm

 		
  g  nameg  timeout ChR4iijiki>  "  G  li-iGiMiPiSiZiYieifigihiJi,i6 V      g  filenamef  	agnes.scm		
		
	#	
	%			(	
	*			-	
9	
ö	?
º	D
		R
	X
¦	b
n	h
] 
t 
 
Ã 
O 
 
ñ °
w ¹
 R Î
!· Ó
#n Ù
% å
'm î
)s ù
*5
+r	
+ß
,g
,Ó
,Ô
-'
 $	-'
   C6 