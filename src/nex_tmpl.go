package fas_parser

const Nex_tmpl = `#NEXUS
BEGIN DATA;
	DIMENSIONS NTAX={{ .ntax }} NCHAR={{ .nchar }};
	FORMAT DATATYPE=DNA GAP=- MISSING=? ; #DATATYPE={}`

// MATRIX{{ range $k, $v := .matrix }}
// '{{ $k }}' {{ $v }}{{ end }}
// ;
// END;`
// BEGIN SETS;{{ range $_, $i := .charset }}
// 	CHARSET {{ .name }} = {{ .from }}-{{ .to }};{{ end }}
// END;
// `
