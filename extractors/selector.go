package selector

import (
	"strings"

	"taCaro-backend/extractors/mt"
	extract_info_interface "taCaro-backend/interfaces"
)

// SelectExtractor seleciona e retorna o extractor adequado com base na URL.
func SelectExtractor(url string) extract_info_interface.ExtractorInterface {
	if strings.Contains(url, "sefaz.mt") {
		return &mt.Extractor{}
	}
	// Aqui você pode adicionar outros casos para diferentes URLs.
	// Por exemplo:
	// if strings.Contains(url, "site-abc.com") {
	//     return &abc.Extractor{}
	// }

	// Se nenhum caso casar, retorna um extractor default.
	return &mt.Extractor{}
}
