// - go help doc
// - go doc demonstra a documentação de um package, const, func, type, var, método, etc.
// - go doc aceita zero, um, ou dois argumentos:
//     - zero: demonstra a documentação do package do diretório atual
//     - um: toma argumentos nos padrões abaixo
//         - go doc (pkg)
//         - go doc (sym)[.(method)]
//         - go doc [(pkg).](sym)[.(method)]
//         - go doc [(pkg).][(sym).](method)
//     - dois: o primeiro argumento deve ser o nome do package
//         - go doc (pkg) (sym)[.(method)]