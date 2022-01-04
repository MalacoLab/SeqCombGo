package fas_parser

const Nex_tmpl = `#NEXUS
BEGIN DATA;
	DIMENSIONS NTAX={{ .Ntax }} NCHAR={{ .Nchar }};
	FORMAT DATATYPE=DNA GAP=- MISSING=? ; #DATATYPE={}
MATRIX
{{- range $k, $v := .Matrix }}
'{{ $k }}' {{ $v }}{{ end }}
;
END;
BEGIN SETS;
{{- range $_, $i := .Charset }}
	CHARSET {{ $i.Name }} = {{ $i.From }}-{{ $i.To }};
{{- end }}
END;
`

// 最后那个 $i 好像有问题
// {{/* $k| printf "%-40s" */}}