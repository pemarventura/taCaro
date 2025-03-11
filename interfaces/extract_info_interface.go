// extractor_interface.go (pode estar em um pacote compartilhado ou no main)
package extract_info_interface

import "taCaro-backend/models"

// ExtractorInterface define o método que todos os extractors devem implementar.
type ExtractorInterface interface {
	ExtractInfo(htmlContent string) []models.Item
}
