package nex_tmpl

const Nex_tmpl = `#NEXUS
BEGIN DATA;
	DIMENSIONS NTAX={} NCHAR={};
	FORMAT DATATYPE=DNA GAP=- MISSING=? ; #DATATYPE={}
MATRIX{{ range $k, $v := . }}
'{{ $k }}' {{ $v }}{{ end }}
;
END;
BEGIN SETS;
	CHARSET 16S = 1-811;
	CHARSET CO1 = 812-1421;
END;
`
