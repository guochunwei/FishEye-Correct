GOOF----LE-8-2.0ù6      ] U 4 hÕ      ] g  guile¤	 ¤	g  process-use-modules¤	 ¤	 ¤	g  	aisleriot¤	g  	interface¤	 ¤		 ¤	
g  api¤	
 ¤	 ¤	
		 ¤	g  
foundation¤									
		 	¤	g  tableau¤								 ¤	g  reserve¤	g  initialize-playing-area¤	g  set-ace-low¤	g  make-standard-deck¤	g  shuffle-deck¤	g  add-blank-slot¤	g  add-normal-slot¤	g  DECK¤	g  add-carriage-return-slot¤	g  add-extended-slot¤	g  down¤	g  HORIZPOS¤	g  VERTPOS¤	e  0.5¤	 g  
deal-cards¤	!								
										
									
																													 4¤	"g  map¤	#g  flip-top-card¤	$								
									 ¤	%g  new-game¤	&g  empty-slot?¤	'g  is-visible?¤	(g  reverse¤	)g  button-pressed¤	*g  length¤	+g  	get-value¤	,g  ace¤	-g  get-suit¤	.g  get-top-card¤	/g  is-red?¤	0g  	is-black?¤	1g  
droppable?¤	2g  move-n-cards!¤	3g  add-to-score!¤	4g  make-visible-top-card¤	5g  button-released¤	6g  button-clicked¤	7g  move-ace-to-foundation¤	8g  move-card-to-foundation¤	9g  button-double-clicked¤	:g  game-won¤	;g  game-continuable¤	<g  	get-cards¤	=g  strip¤	>g  
strip-size¤	?g  
check-plop¤	@g  	hint-move¤	Ag  check-a-slot-to-foundations¤	Bg  check-uncover¤	Cg  check-a-foundation-for-uncover¤	Dg  check-foundation-for-uncover¤	Eg  check-empty-tslot¤	Fg  check-to-foundations¤	Gg  find-empty-slot¤	Hg  check-simple-foundation¤	Ig  get-min-happy-foundation¤	Jg  min¤	Kg  king¤	Lg  _¤	Mf  Try rearranging the cards¤	Ng  get-hint¤	Og  get-options¤	Pg  apply-options¤	Qg  timeout¤	Rg  set-features¤	Sg  droppable-feature¤	Tg  
set-lambda¤C 5      hH.  _  ] 4	 >  "  G  RRR !"#$  h`  Ý  ] 4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4>   "  G  4>  "  G  4	>   "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  4
>  "  G  
 
 4	>   "  G   		 4>  "  G  4	>   "  G  		 4>  "  G  4	>   "  G  		 4>  "  G  4	>   "  G  		 
 
 4	>   "  G  		 4>   "  G  4>  "  G  4	>   "  G  		 4>   "  G  4>  "  G  4	>   "  G  		 4>   "  G  4>  "  G  4	>   "  G  		 4>   "  G  4>  "  G  4
>  "  G  4>  "  G  		 C     Õ      g  filenamef  king-albert.scm
	
							#			3			C			S			Y			^			g			w			z			|		 		 	 	 	!	 	!	 	!	 ¤	!	 ­	"	 ½	#	 À	#	 Â	#	 Ç	#	 Ð	$	 à	&	 ã	&	 ç	&	 ì	&	 õ	'	 ø	'	 ü	'		'	
	(		(		(		(		)	"	)	&	)	+	)	4	*	7	*	;	*	@	*	I	+	L	+	P	+	U	+	^	,	a	,	e	,	j	,	s	-	v	-	z	-		-		.		.		.		.		0	¢	1	£	3	¶	5	·	5	¹	5	¾	6	À	6	Á	7	Ä	7	Æ	7	Ë	7	Ô	8	è	9	ê	9	ë	:	î	:	ð	:	õ	:	þ	;		<		<		=		=		=		=	(	>	<	?	>	?	A	A	D	B	E	D	Y	F	[	F	\	G	l	H	o	H	q	H	v	H		I		J		J		K	¦	L	©	L	«	L	°	L	¹	M	Í	N	Ï	N	Ð	O	à	P	ã	P	å	P	ê	P	ó	Q		R			R	
	S		T		T		T	$	T	-	V	2	V	7	V	@	Z	F	Z	K	Z	Z	\	 	[
  g  nameg  new-game C%R&'(      h    ±   ]4 5$  C456     ©       g  slot-id
		 g  	card-list		  g  filenamef  king-albert.scm
	^
		_			_			`			`			`	 			  g  nameg  button-pressed C)R*+,&-./(0 
     hè     ] $  C	$  t45$  e45$  45"  $  C45$  C445545$  454455CCC	$  M45$  C44554	455&  44554455CCC      {      g  
start-slot
	 â g  	card-list	 â g  end-slot		 â g  t		>  g  t	  à  g  filenamef  king-albert.scm
	b
		c				c			e				c			f			f		#	f			$	g		)	g	%	+	g		.	g		2	g		3	h		>	g		J	i		T	i		W	j		Z	j	$	b	j		c	k		h	k	$	j	k		k	j		o	i		p	l		u	l	%	w	l		x	m		{	m	* 	m	 	m	 	l	 	n		 	c	 	o	 	o	 ¦	p	 ©	p	* °	p	% ²	p	 ³	q	 ¶	q	' ¾	q	 Â	p	 Ã	r	 Æ	r	/ Í	r	* Ï	r	 Ð	r	 Ñ	s	 Ô	s	% Ü	s	 Ý	r	 7	 â	  g  nameg  
droppable? C1R123&4    h   H  ]4 5$  l4 5$  [ 	$  "  4	ÿ5$  9	$  "  45$  4 5$  C 6CCCC @      g  
start-slot
		 g  	card-list		 g  end-slot			 g  t		%	< g  t		E	[ g  t		f	w  g  filenamef  king-albert.scm
	v
		w			w			x		 	w		%	y		%	y		3	z		@	w		E	{		E	{		S	|		_	w		`	}		f	}		w	~	 			  g  nameg  button-released C5R   h   s   ]C    k       g  slot-id
		  g  filenamef  king-albert.scm
 
 		  g  nameg  button-clicked C6R&7  h    °   ]45$  C 6       ¨       g  slot
		 g  f-slot		  g  filenamef  king-albert.scm
 
	 		 		 	#	 	 			  g  nameg  move-ace-to-foundation C7R&-.+8       hp   \  ]	$  C45$  "  <44 554455$  44 554455"  $  C 6       T      g  slot
		i g  f-slot		i  g  filenamef  king-albert.scm
 
	 			 		 		 			 		! 		) 		* 		- 		5 		6 		: 			; 		> 		F 		G 		J 	!	R 		S 		T 		] 		g 	*	i 	 		i	  g  nameg  move-card-to-foundation C8R&+., 7834 
  h¨     ]	 	$  4 5$  C44 55$  4 4 
5 5"  $  "  #4 
5$  4 4 
5 5"  $  $45$  4 5$  C	 6CCC             g  slot-id
	 ¡ g  t	B	u g  t    g  filenamef  king-albert.scm
 
	 		 		 		 		 		 		$ 		' 		+ 		, 		1 	*	; 	$	= 		B 		P 		[ 		\ 		a 	*	k 	$	m 		y 		z 	  	  	  	  	 	 ¡  g  nameg  button-double-clicked C9R:       h   q   ] 45 C       i       g  filenamef  king-albert.scm
  
	 ¡		 ¡	 			
  g  nameg  game-continuable C;R*<    hP     ] 44
55	$  94455	$  %44	55	$  44	55	CCCC ú       g  filenamef  king-albert.scm
 £
	 ¤	
	 ¤		 ¤	
	 ¤		 ¤		 ¥	
	 ¥		 ¥	
	" ¥		& ¤		' ¦	
	* ¦		2 ¦	
	5 ¦		9 ¤		: §	
	= §		E §	
	H §	 		O
  g  nameg  game-won C:R*'= h8   Ý   ]	4 5$  "  
4 5$   C 6    Õ       g  	card-list
		4 g  t		%  g  filenamef  king-albert.scm
 ©
	 ª		 ª	
	 ª		 «		 «		! «		" «	
	) ª		, ¬		2 ­		4 ­	 		4  g  nameg  strip C=R*'>      h8   û   ]
4 5$  "  
4 5$  C 6 ó       g  	card-list
		7 g  n		7 g  t			%  g  filenamef  king-albert.scm
 ¯
	 °		 °	
	 °		 ±		 ±		! ±		" ±	
	) °		, ²		2 ³		5 ³	"	7 ³	 		7	  g  nameg  
strip-size C>R&/0.+?  hp   4  ]	$  C45$  "  14 54455&  4 54455"  $  C45$  C 6    ,      g  card
		l g  t-slot		l  g  filenamef  king-albert.scm
 µ
	 ¶			 ¶		 ¸		 ¸			 ¹		% º		( º		0 º		4 ¸			5 »		< »		= ¼		@ ¼		H ¼		I »		R ¶		V ¾			` ¶		j À		l À	 		l	  g  nameg  
check-plop C?R&?=<@>*'AB h¸   À  ] 	$  C4 5$  "  444 55	5$  & 44 5
5444 55	564 5$  "  644 55$  #44 55$  "  4	 
5"  $  	 
6
 6      ¸      g  t-slot
	 ²  g  filenamef  king-albert.scm
 Â
	 Ã			 Ã		 Å		 Å			 Æ		! Æ		$ Æ	!	, Æ		0 Æ		4 Ã		9 Ç		< Ç	'	E Ç		F Ç	=	I Ç	I	L Ç	P	T Ç	I	X Ç	=	Z Ç			[ È		e È			k É		n É		v É		x É		| È			} Ê	  Ê	&  Ê	   Ê	  È		  Ë	 ¤ Ã	 « Ì		 ° Í	 ² Í	 &	 ²  g  nameg  check-uncover CBR&/0.+?@C 	h     ]	$  C45$  "  H4 54455&  04 54455$  445	5"  "  $  445	56 6 {      g  card
	  g  f-slot	   g  filenamef  king-albert.scm
 Ï
	 Ð			 Ð		 Ò		 Ò			 Ó		% Ô		( Ô		0 Ô		4 Ò			5 Õ		< Õ		= Ö		@ Ö		H Ö		I Õ		M Ò			N ×		Q ×		[ ×		i Ð		o Ø		r Ø	)	| Ø		~ Ø		  Ù	1  Ù	 	 	  g  nameg  check-a-foundation-for-uncover CCR&'(<C=D        hh   O  ] 	$  C4 5$  "  -444 555$  "  444 55
5$  44 55
6 6  G      g  t-slot
		f  g  filenamef  king-albert.scm
 Û
	 Ü			 Ü		 Þ		 Þ			 ß		! ß	%	$ ß	.	, ß	%	- ß	 	/ ß		3 Þ			9 à		< à	.	? à	5	G à	.	J à		N Ü		Q á	)	T á	0	\ á	)	_ á			d â	*	f â	 		f  g  nameg  check-foundation-for-uncover CDR&'(<?E@* 	 h¨   ò  ] 	$  C4 5$  "  /444 555$  444 55	5"  $  J4444 55	55$   6 44 55444 55	56 6       ê      g  t-slot
	 ¡  g  filenamef  king-albert.scm
 ä
	 å			 å		 ç		 ç			 è		! è	 	$ è	)	, è	 	- è		/ è		3 ç			4 é		7 é		: é	(	B é		C é		G é		P å		Q ê		T ê		W ê	+	Z ê	4	b ê	+	c ê	&	g ê		i ê		m ê			r ë	 	t ë		y ì		| ì	'  ì	  ì	;  ì	L  ì	U  ì	L  ì	G  ì	;  ì	  í	 ¡ í	 +	 ¡  g  nameg  check-empty-tslot CER&FA h@   í   ] 	$  C4 5$  	 64 5$   6 6 å       g  slot
		? g  f-slot		?  g  filenamef  king-albert.scm
 ï
	 ð			 ð		 ò			 ð		 ó		! ó			" ô			. ð		6 õ			; ö	"	? ö	 		?	  g  nameg  check-to-foundations CFR+.,@G&-A 
 h   ®  ]	$  C44 55$   45645$  "  <44 554455$  44 554455"  $  	 6	 6¦      g  slot
	  g  f-slot	   g  filenamef  king-albert.scm
 ø
	 ù			 ù		 û		 û		 û		 û			  ù		& ý		. ý			/ þ		9 þ			? ÿ		B ÿ		J ÿ		K 		N 		V 		W ÿ		[ þ			\		_		g		h		k	!	s		t		u		~ ù	 		 	. 	 !	 	  g  nameg  check-a-slot-to-foundations CAR&+.AH hX     ] 	$  C4 5$  "   44 55$  4 
5"  $   
6 6             g  slot
		Q g  happynum		Q  g  filenamef  king-albert.scm

													
		!
		)
		,
		0				1		A		H			M	%	Q	 		Q	  g  nameg  check-simple-foundation CHR&IJ+.  h@   ú   ] 	$  C4 5$  	C 4	44 5556       ò       g  fslot
		9 g  value		9  g  filenamef  king-albert.scm

											!	&	"	&	)	6	,	A	4	6	5	1	7	&	9	 		9	  g  nameg  get-min-happy-foundation CIRHIKBEDFLM 
    hx   /  ]4	4	55  $   C4	5  $   C4	5  $   C4	5  $   C4	
5  $   C
4	5 C      '      g  t
		r g  t
	$	r g  t
	6	r g  t
	H	r g  t
	[	r  g  filenamef  king-albert.scm

				!							$		0		6		B		H		T		[		h		l		n		q	 		r
  g  nameg  get-hint CNR    h   Z   ] C    R       g  filenamef  king-albert.scm

 		
  g  nameg  get-options COR h   r   ]C    j       g  options
		  g  filenamef  king-albert.scm
"
 		  g  nameg  apply-options CPR h   V   ] C    N       g  filenamef  king-albert.scm
%
 		
  g  nameg  timeout CQR4RiSi>  "  G  Ti%i)i5i6i9i;i:iNiOiPiQi1i6   W      g  filenamef  king-albert.scm		
					
	 			#	
	%			(	
 	
		^
	b
	v
 
s 
_ 
³ 
L  
µ £
à ©
. ¯
ï µ
 Â
¶ Ï
! Û
$M ä
% ï
'ñ ø
)p
*Å
,
-
-"
-ù%
-ú(
.E*
 #	.E
   C6 