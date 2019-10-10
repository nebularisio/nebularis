// Code generated from /Users/ozbenevren/dev/go/src/nebularis.io/nebularis/Nebularis.g4 by ANTLR 4.7.1. DO NOT EDIT.

package parser // Nebularis

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 70, 527,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 3, 2, 5, 2, 70, 10, 2, 3, 2, 7, 2, 73, 10, 2, 12, 2, 14, 2, 76,
	11, 2, 3, 2, 7, 2, 79, 10, 2, 12, 2, 14, 2, 82, 11, 2, 3, 2, 3, 2, 3, 3,
	7, 3, 87, 10, 3, 12, 3, 14, 3, 90, 11, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4,
	7, 4, 97, 10, 4, 12, 4, 14, 4, 100, 11, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4,
	106, 10, 4, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 5, 6, 114, 10, 6, 3, 7,
	7, 7, 117, 10, 7, 12, 7, 14, 7, 120, 11, 7, 3, 7, 5, 7, 123, 10, 7, 3,
	7, 5, 7, 126, 10, 7, 3, 7, 5, 7, 129, 10, 7, 3, 7, 3, 7, 5, 7, 133, 10,
	7, 3, 7, 3, 7, 3, 7, 3, 7, 5, 7, 139, 10, 7, 3, 7, 3, 7, 3, 7, 5, 7, 144,
	10, 7, 3, 7, 7, 7, 147, 10, 7, 12, 7, 14, 7, 150, 11, 7, 3, 7, 3, 7, 3,
	8, 3, 8, 3, 8, 3, 8, 7, 8, 158, 10, 8, 12, 8, 14, 8, 161, 11, 8, 5, 8,
	163, 10, 8, 3, 8, 3, 8, 3, 9, 5, 9, 168, 10, 9, 3, 9, 3, 9, 3, 10, 3, 10,
	7, 10, 174, 10, 10, 12, 10, 14, 10, 177, 11, 10, 3, 10, 3, 10, 3, 11, 3,
	11, 3, 11, 3, 11, 5, 11, 185, 10, 11, 3, 11, 3, 11, 5, 11, 189, 10, 11,
	3, 11, 3, 11, 3, 11, 5, 11, 194, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 5, 11, 202, 10, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 5, 11, 212, 10, 11, 3, 11, 3, 11, 3, 11, 5, 11, 217, 10,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 7, 11, 225, 10, 11, 12, 11,
	14, 11, 228, 11, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3,
	11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5, 11, 245, 10, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 7, 12, 251, 10, 12, 12, 12, 14, 12, 254, 11,
	12, 3, 12, 3, 12, 6, 12, 258, 10, 12, 13, 12, 14, 12, 259, 3, 12, 5, 12,
	263, 10, 12, 3, 13, 3, 13, 3, 13, 6, 13, 268, 10, 13, 13, 13, 14, 13, 269,
	3, 13, 5, 13, 273, 10, 13, 3, 14, 7, 14, 276, 10, 14, 12, 14, 14, 14, 279,
	11, 14, 3, 14, 3, 14, 3, 14, 5, 14, 284, 10, 14, 3, 14, 3, 14, 7, 14, 288,
	10, 14, 12, 14, 14, 14, 291, 11, 14, 3, 14, 5, 14, 294, 10, 14, 3, 15,
	3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 306,
	10, 15, 3, 15, 3, 15, 3, 15, 7, 15, 311, 10, 15, 12, 15, 14, 15, 314, 11,
	15, 3, 15, 3, 15, 3, 15, 5, 15, 319, 10, 15, 3, 15, 3, 15, 7, 15, 323,
	10, 15, 12, 15, 14, 15, 326, 11, 15, 3, 15, 5, 15, 329, 10, 15, 3, 16,
	3, 16, 3, 16, 3, 16, 3, 16, 7, 16, 336, 10, 16, 12, 16, 14, 16, 339, 11,
	16, 3, 16, 3, 16, 5, 16, 343, 10, 16, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17,
	349, 10, 17, 12, 17, 14, 17, 352, 11, 17, 5, 17, 354, 10, 17, 3, 18, 7,
	18, 357, 10, 18, 12, 18, 14, 18, 360, 11, 18, 3, 18, 3, 18, 3, 18, 5, 18,
	365, 10, 18, 3, 18, 7, 18, 368, 10, 18, 12, 18, 14, 18, 371, 11, 18, 3,
	18, 5, 18, 374, 10, 18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 20, 3, 20,
	7, 20, 383, 10, 20, 12, 20, 14, 20, 386, 11, 20, 3, 20, 3, 20, 3, 21, 3,
	21, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 397, 10, 22, 3, 23, 7, 23,
	400, 10, 23, 12, 23, 14, 23, 403, 11, 23, 3, 23, 3, 23, 3, 23, 7, 23, 408,
	10, 23, 12, 23, 14, 23, 411, 11, 23, 3, 23, 5, 23, 414, 10, 23, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 430, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 7, 24, 450, 10, 24, 12, 24, 14, 24, 453, 11, 24, 5, 24,
	455, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24, 7, 24, 463, 10,
	24, 12, 24, 14, 24, 466, 11, 24, 3, 25, 3, 25, 5, 25, 470, 10, 25, 3, 25,
	3, 25, 5, 25, 474, 10, 25, 3, 25, 7, 25, 477, 10, 25, 12, 25, 14, 25, 480,
	11, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27,
	3, 27, 7, 27, 492, 10, 27, 12, 27, 14, 27, 495, 11, 27, 3, 27, 5, 27, 498,
	10, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33, 3, 33, 7, 33, 517, 10,
	33, 12, 33, 14, 33, 520, 11, 33, 3, 34, 3, 34, 3, 34, 5, 34, 525, 10, 34,
	3, 34, 2, 3, 46, 35, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28,
	30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64,
	66, 2, 5, 4, 2, 55, 55, 66, 66, 4, 2, 53, 53, 60, 60, 8, 2, 35, 36, 38,
	38, 46, 52, 54, 54, 56, 57, 63, 63, 2, 590, 2, 69, 3, 2, 2, 2, 4, 88, 3,
	2, 2, 2, 6, 98, 3, 2, 2, 2, 8, 109, 3, 2, 2, 2, 10, 113, 3, 2, 2, 2, 12,
	118, 3, 2, 2, 2, 14, 153, 3, 2, 2, 2, 16, 167, 3, 2, 2, 2, 18, 171, 3,
	2, 2, 2, 20, 244, 3, 2, 2, 2, 22, 246, 3, 2, 2, 2, 24, 264, 3, 2, 2, 2,
	26, 277, 3, 2, 2, 2, 28, 328, 3, 2, 2, 2, 30, 330, 3, 2, 2, 2, 32, 344,
	3, 2, 2, 2, 34, 358, 3, 2, 2, 2, 36, 375, 3, 2, 2, 2, 38, 378, 3, 2, 2,
	2, 40, 389, 3, 2, 2, 2, 42, 396, 3, 2, 2, 2, 44, 401, 3, 2, 2, 2, 46, 429,
	3, 2, 2, 2, 48, 467, 3, 2, 2, 2, 50, 483, 3, 2, 2, 2, 52, 497, 3, 2, 2,
	2, 54, 499, 3, 2, 2, 2, 56, 503, 3, 2, 2, 2, 58, 505, 3, 2, 2, 2, 60, 507,
	3, 2, 2, 2, 62, 509, 3, 2, 2, 2, 64, 513, 3, 2, 2, 2, 66, 524, 3, 2, 2,
	2, 68, 70, 5, 4, 3, 2, 69, 68, 3, 2, 2, 2, 69, 70, 3, 2, 2, 2, 70, 74,
	3, 2, 2, 2, 71, 73, 5, 6, 4, 2, 72, 71, 3, 2, 2, 2, 73, 76, 3, 2, 2, 2,
	74, 72, 3, 2, 2, 2, 74, 75, 3, 2, 2, 2, 75, 80, 3, 2, 2, 2, 76, 74, 3,
	2, 2, 2, 77, 79, 5, 10, 6, 2, 78, 77, 3, 2, 2, 2, 79, 82, 3, 2, 2, 2, 80,
	78, 3, 2, 2, 2, 80, 81, 3, 2, 2, 2, 81, 83, 3, 2, 2, 2, 82, 80, 3, 2, 2,
	2, 83, 84, 7, 2, 2, 3, 84, 3, 3, 2, 2, 2, 85, 87, 5, 62, 32, 2, 86, 85,
	3, 2, 2, 2, 87, 90, 3, 2, 2, 2, 88, 86, 3, 2, 2, 2, 88, 89, 3, 2, 2, 2,
	89, 91, 3, 2, 2, 2, 90, 88, 3, 2, 2, 2, 91, 92, 7, 22, 2, 2, 92, 93, 5,
	64, 33, 2, 93, 94, 7, 62, 2, 2, 94, 5, 3, 2, 2, 2, 95, 97, 5, 62, 32, 2,
	96, 95, 3, 2, 2, 2, 97, 100, 3, 2, 2, 2, 98, 96, 3, 2, 2, 2, 98, 99, 3,
	2, 2, 2, 99, 101, 3, 2, 2, 2, 100, 98, 3, 2, 2, 2, 101, 102, 7, 19, 2,
	2, 102, 105, 5, 8, 5, 2, 103, 104, 7, 6, 2, 2, 104, 106, 7, 67, 2, 2, 105,
	103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 108,
	7, 62, 2, 2, 108, 7, 3, 2, 2, 2, 109, 110, 5, 64, 33, 2, 110, 9, 3, 2,
	2, 2, 111, 114, 5, 26, 14, 2, 112, 114, 5, 12, 7, 2, 113, 111, 3, 2, 2,
	2, 113, 112, 3, 2, 2, 2, 114, 11, 3, 2, 2, 2, 115, 117, 5, 62, 32, 2, 116,
	115, 3, 2, 2, 2, 117, 120, 3, 2, 2, 2, 118, 116, 3, 2, 2, 2, 118, 119,
	3, 2, 2, 2, 119, 122, 3, 2, 2, 2, 120, 118, 3, 2, 2, 2, 121, 123, 7, 15,
	2, 2, 122, 121, 3, 2, 2, 2, 122, 123, 3, 2, 2, 2, 123, 125, 3, 2, 2, 2,
	124, 126, 7, 23, 2, 2, 125, 124, 3, 2, 2, 2, 125, 126, 3, 2, 2, 2, 126,
	128, 3, 2, 2, 2, 127, 129, 7, 24, 2, 2, 128, 127, 3, 2, 2, 2, 128, 129,
	3, 2, 2, 2, 129, 130, 3, 2, 2, 2, 130, 132, 7, 17, 2, 2, 131, 133, 5, 38,
	20, 2, 132, 131, 3, 2, 2, 2, 132, 133, 3, 2, 2, 2, 133, 138, 3, 2, 2, 2,
	134, 135, 7, 58, 2, 2, 135, 136, 5, 30, 16, 2, 136, 137, 7, 59, 2, 2, 137,
	139, 3, 2, 2, 2, 138, 134, 3, 2, 2, 2, 138, 139, 3, 2, 2, 2, 139, 140,
	3, 2, 2, 2, 140, 141, 7, 67, 2, 2, 141, 143, 5, 14, 8, 2, 142, 144, 5,
	28, 15, 2, 143, 142, 3, 2, 2, 2, 143, 144, 3, 2, 2, 2, 144, 148, 3, 2,
	2, 2, 145, 147, 5, 36, 19, 2, 146, 145, 3, 2, 2, 2, 147, 150, 3, 2, 2,
	2, 148, 146, 3, 2, 2, 2, 148, 149, 3, 2, 2, 2, 149, 151, 3, 2, 2, 2, 150,
	148, 3, 2, 2, 2, 151, 152, 5, 18, 10, 2, 152, 13, 3, 2, 2, 2, 153, 162,
	7, 58, 2, 2, 154, 159, 5, 16, 9, 2, 155, 156, 7, 40, 2, 2, 156, 158, 5,
	16, 9, 2, 157, 155, 3, 2, 2, 2, 158, 161, 3, 2, 2, 2, 159, 157, 3, 2, 2,
	2, 159, 160, 3, 2, 2, 2, 160, 163, 3, 2, 2, 2, 161, 159, 3, 2, 2, 2, 162,
	154, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 164, 3, 2, 2, 2, 164, 165,
	7, 59, 2, 2, 165, 15, 3, 2, 2, 2, 166, 168, 7, 67, 2, 2, 167, 166, 3, 2,
	2, 2, 167, 168, 3, 2, 2, 2, 168, 169, 3, 2, 2, 2, 169, 170, 5, 28, 15,
	2, 170, 17, 3, 2, 2, 2, 171, 175, 7, 41, 2, 2, 172, 174, 5, 20, 11, 2,
	173, 172, 3, 2, 2, 2, 174, 177, 3, 2, 2, 2, 175, 173, 3, 2, 2, 2, 175,
	176, 3, 2, 2, 2, 176, 178, 3, 2, 2, 2, 177, 175, 3, 2, 2, 2, 178, 179,
	7, 42, 2, 2, 179, 19, 3, 2, 2, 2, 180, 181, 7, 32, 2, 2, 181, 184, 7, 67,
	2, 2, 182, 183, 7, 39, 2, 2, 183, 185, 5, 28, 15, 2, 184, 182, 3, 2, 2,
	2, 184, 185, 3, 2, 2, 2, 185, 188, 3, 2, 2, 2, 186, 187, 7, 45, 2, 2, 187,
	189, 5, 46, 24, 2, 188, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 190,
	3, 2, 2, 2, 190, 245, 7, 62, 2, 2, 191, 193, 7, 25, 2, 2, 192, 194, 5,
	46, 24, 2, 193, 192, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195, 3, 2,
	2, 2, 195, 245, 7, 62, 2, 2, 196, 197, 7, 18, 2, 2, 197, 198, 5, 46, 24,
	2, 198, 201, 5, 18, 10, 2, 199, 200, 7, 14, 2, 2, 200, 202, 5, 18, 10,
	2, 201, 199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 245, 3, 2, 2, 2, 203,
	204, 7, 34, 2, 2, 204, 205, 5, 46, 24, 2, 205, 206, 5, 18, 10, 2, 206,
	245, 3, 2, 2, 2, 207, 208, 7, 67, 2, 2, 208, 245, 7, 39, 2, 2, 209, 211,
	7, 12, 2, 2, 210, 212, 7, 67, 2, 2, 211, 210, 3, 2, 2, 2, 211, 212, 3,
	2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 245, 7, 62, 2, 2, 214, 216, 7, 8, 2,
	2, 215, 217, 7, 67, 2, 2, 216, 215, 3, 2, 2, 2, 216, 217, 3, 2, 2, 2, 217,
	218, 3, 2, 2, 2, 218, 245, 7, 62, 2, 2, 219, 220, 7, 28, 2, 2, 220, 221,
	5, 46, 24, 2, 221, 226, 7, 41, 2, 2, 222, 225, 5, 22, 12, 2, 223, 225,
	5, 24, 13, 2, 224, 222, 3, 2, 2, 2, 224, 223, 3, 2, 2, 2, 225, 228, 3,
	2, 2, 2, 226, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2, 2, 227, 229, 3, 2, 2,
	2, 228, 226, 3, 2, 2, 2, 229, 230, 7, 42, 2, 2, 230, 245, 3, 2, 2, 2, 231,
	232, 5, 46, 24, 2, 232, 233, 5, 58, 30, 2, 233, 234, 7, 62, 2, 2, 234,
	245, 3, 2, 2, 2, 235, 236, 5, 46, 24, 2, 236, 237, 7, 45, 2, 2, 237, 238,
	5, 46, 24, 2, 238, 239, 7, 62, 2, 2, 239, 245, 3, 2, 2, 2, 240, 245, 7,
	44, 2, 2, 241, 242, 5, 46, 24, 2, 242, 243, 7, 62, 2, 2, 243, 245, 3, 2,
	2, 2, 244, 180, 3, 2, 2, 2, 244, 191, 3, 2, 2, 2, 244, 196, 3, 2, 2, 2,
	244, 203, 3, 2, 2, 2, 244, 207, 3, 2, 2, 2, 244, 209, 3, 2, 2, 2, 244,
	214, 3, 2, 2, 2, 244, 219, 3, 2, 2, 2, 244, 231, 3, 2, 2, 2, 244, 235,
	3, 2, 2, 2, 244, 240, 3, 2, 2, 2, 244, 241, 3, 2, 2, 2, 245, 21, 3, 2,
	2, 2, 246, 247, 7, 10, 2, 2, 247, 252, 5, 46, 24, 2, 248, 249, 7, 40, 2,
	2, 249, 251, 5, 46, 24, 2, 250, 248, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2,
	252, 250, 3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 255, 3, 2, 2, 2, 254,
	252, 3, 2, 2, 2, 255, 262, 7, 39, 2, 2, 256, 258, 5, 20, 11, 2, 257, 256,
	3, 2, 2, 2, 258, 259, 3, 2, 2, 2, 259, 257, 3, 2, 2, 2, 259, 260, 3, 2,
	2, 2, 260, 263, 3, 2, 2, 2, 261, 263, 5, 18, 10, 2, 262, 257, 3, 2, 2,
	2, 262, 261, 3, 2, 2, 2, 262, 263, 3, 2, 2, 2, 263, 23, 3, 2, 2, 2, 264,
	265, 7, 13, 2, 2, 265, 272, 7, 39, 2, 2, 266, 268, 5, 20, 11, 2, 267, 266,
	3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269, 270, 3, 2,
	2, 2, 270, 273, 3, 2, 2, 2, 271, 273, 5, 18, 10, 2, 272, 267, 3, 2, 2,
	2, 272, 271, 3, 2, 2, 2, 272, 273, 3, 2, 2, 2, 273, 25, 3, 2, 2, 2, 274,
	276, 5, 62, 32, 2, 275, 274, 3, 2, 2, 2, 276, 279, 3, 2, 2, 2, 277, 275,
	3, 2, 2, 2, 277, 278, 3, 2, 2, 2, 278, 280, 3, 2, 2, 2, 279, 277, 3, 2,
	2, 2, 280, 281, 7, 30, 2, 2, 281, 283, 7, 67, 2, 2, 282, 284, 5, 38, 20,
	2, 283, 282, 3, 2, 2, 2, 283, 284, 3, 2, 2, 2, 284, 285, 3, 2, 2, 2, 285,
	289, 5, 28, 15, 2, 286, 288, 5, 36, 19, 2, 287, 286, 3, 2, 2, 2, 288, 291,
	3, 2, 2, 2, 289, 287, 3, 2, 2, 2, 289, 290, 3, 2, 2, 2, 290, 293, 3, 2,
	2, 2, 291, 289, 3, 2, 2, 2, 292, 294, 7, 62, 2, 2, 293, 292, 3, 2, 2, 2,
	293, 294, 3, 2, 2, 2, 294, 27, 3, 2, 2, 2, 295, 296, 7, 61, 2, 2, 296,
	329, 5, 28, 15, 2, 297, 298, 7, 64, 2, 2, 298, 299, 7, 65, 2, 2, 299, 329,
	5, 28, 15, 2, 300, 329, 5, 30, 16, 2, 301, 329, 5, 42, 22, 2, 302, 303,
	7, 17, 2, 2, 303, 305, 5, 14, 8, 2, 304, 306, 5, 28, 15, 2, 305, 304, 3,
	2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 329, 3, 2, 2, 2, 307, 308, 7, 27, 2,
	2, 308, 312, 7, 41, 2, 2, 309, 311, 5, 44, 23, 2, 310, 309, 3, 2, 2, 2,
	311, 314, 3, 2, 2, 2, 312, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313,
	315, 3, 2, 2, 2, 314, 312, 3, 2, 2, 2, 315, 329, 7, 42, 2, 2, 316, 318,
	7, 20, 2, 2, 317, 319, 5, 32, 17, 2, 318, 317, 3, 2, 2, 2, 318, 319, 3,
	2, 2, 2, 319, 320, 3, 2, 2, 2, 320, 324, 7, 41, 2, 2, 321, 323, 5, 34,
	18, 2, 322, 321, 3, 2, 2, 2, 323, 326, 3, 2, 2, 2, 324, 322, 3, 2, 2, 2,
	324, 325, 3, 2, 2, 2, 325, 327, 3, 2, 2, 2, 326, 324, 3, 2, 2, 2, 327,
	329, 7, 42, 2, 2, 328, 295, 3, 2, 2, 2, 328, 297, 3, 2, 2, 2, 328, 300,
	3, 2, 2, 2, 328, 301, 3, 2, 2, 2, 328, 302, 3, 2, 2, 2, 328, 307, 3, 2,
	2, 2, 328, 316, 3, 2, 2, 2, 329, 29, 3, 2, 2, 2, 330, 342, 5, 64, 33, 2,
	331, 332, 7, 49, 2, 2, 332, 337, 5, 28, 15, 2, 333, 334, 7, 40, 2, 2, 334,
	336, 5, 28, 15, 2, 335, 333, 3, 2, 2, 2, 336, 339, 3, 2, 2, 2, 337, 335,
	3, 2, 2, 2, 337, 338, 3, 2, 2, 2, 338, 340, 3, 2, 2, 2, 339, 337, 3, 2,
	2, 2, 340, 341, 7, 47, 2, 2, 341, 343, 3, 2, 2, 2, 342, 331, 3, 2, 2, 2,
	342, 343, 3, 2, 2, 2, 343, 31, 3, 2, 2, 2, 344, 353, 7, 39, 2, 2, 345,
	350, 5, 30, 16, 2, 346, 347, 7, 40, 2, 2, 347, 349, 5, 30, 16, 2, 348,
	346, 3, 2, 2, 2, 349, 352, 3, 2, 2, 2, 350, 348, 3, 2, 2, 2, 350, 351,
	3, 2, 2, 2, 351, 354, 3, 2, 2, 2, 352, 350, 3, 2, 2, 2, 353, 345, 3, 2,
	2, 2, 353, 354, 3, 2, 2, 2, 354, 33, 3, 2, 2, 2, 355, 357, 5, 62, 32, 2,
	356, 355, 3, 2, 2, 2, 357, 360, 3, 2, 2, 2, 358, 356, 3, 2, 2, 2, 358,
	359, 3, 2, 2, 2, 359, 361, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 361, 362,
	7, 67, 2, 2, 362, 364, 5, 14, 8, 2, 363, 365, 5, 28, 15, 2, 364, 363, 3,
	2, 2, 2, 364, 365, 3, 2, 2, 2, 365, 369, 3, 2, 2, 2, 366, 368, 5, 36, 19,
	2, 367, 366, 3, 2, 2, 2, 368, 371, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2, 369,
	370, 3, 2, 2, 2, 370, 373, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 372, 374,
	7, 62, 2, 2, 373, 372, 3, 2, 2, 2, 373, 374, 3, 2, 2, 2, 374, 35, 3, 2,
	2, 2, 375, 376, 7, 33, 2, 2, 376, 377, 5, 46, 24, 2, 377, 37, 3, 2, 2,
	2, 378, 379, 7, 49, 2, 2, 379, 384, 5, 40, 21, 2, 380, 381, 7, 40, 2, 2,
	381, 383, 5, 40, 21, 2, 382, 380, 3, 2, 2, 2, 383, 386, 3, 2, 2, 2, 384,
	382, 3, 2, 2, 2, 384, 385, 3, 2, 2, 2, 385, 387, 3, 2, 2, 2, 386, 384,
	3, 2, 2, 2, 387, 388, 7, 47, 2, 2, 388, 39, 3, 2, 2, 2, 389, 390, 7, 67,
	2, 2, 390, 41, 3, 2, 2, 2, 391, 397, 7, 21, 2, 2, 392, 397, 7, 7, 2, 2,
	393, 397, 7, 26, 2, 2, 394, 397, 7, 9, 2, 2, 395, 397, 7, 11, 2, 2, 396,
	391, 3, 2, 2, 2, 396, 392, 3, 2, 2, 2, 396, 393, 3, 2, 2, 2, 396, 394,
	3, 2, 2, 2, 396, 395, 3, 2, 2, 2, 397, 43, 3, 2, 2, 2, 398, 400, 5, 62,
	32, 2, 399, 398, 3, 2, 2, 2, 400, 403, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2,
	401, 402, 3, 2, 2, 2, 402, 404, 3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 404,
	405, 7, 67, 2, 2, 405, 409, 5, 28, 15, 2, 406, 408, 5, 36, 19, 2, 407,
	406, 3, 2, 2, 2, 408, 411, 3, 2, 2, 2, 409, 407, 3, 2, 2, 2, 409, 410,
	3, 2, 2, 2, 410, 413, 3, 2, 2, 2, 411, 409, 3, 2, 2, 2, 412, 414, 7, 62,
	2, 2, 413, 412, 3, 2, 2, 2, 413, 414, 3, 2, 2, 2, 414, 45, 3, 2, 2, 2,
	415, 416, 8, 24, 1, 2, 416, 430, 7, 44, 2, 2, 417, 430, 7, 67, 2, 2, 418,
	430, 5, 66, 34, 2, 419, 420, 7, 58, 2, 2, 420, 421, 5, 46, 24, 2, 421,
	422, 7, 59, 2, 2, 422, 430, 3, 2, 2, 2, 423, 430, 7, 31, 2, 2, 424, 430,
	5, 50, 26, 2, 425, 430, 5, 48, 25, 2, 426, 427, 5, 56, 29, 2, 427, 428,
	5, 46, 24, 4, 428, 430, 3, 2, 2, 2, 429, 415, 3, 2, 2, 2, 429, 417, 3,
	2, 2, 2, 429, 418, 3, 2, 2, 2, 429, 419, 3, 2, 2, 2, 429, 423, 3, 2, 2,
	2, 429, 424, 3, 2, 2, 2, 429, 425, 3, 2, 2, 2, 429, 426, 3, 2, 2, 2, 430,
	464, 3, 2, 2, 2, 431, 432, 12, 10, 2, 2, 432, 433, 7, 61, 2, 2, 433, 434,
	5, 46, 24, 2, 434, 435, 7, 39, 2, 2, 435, 436, 5, 46, 24, 11, 436, 463,
	3, 2, 2, 2, 437, 438, 12, 3, 2, 2, 438, 439, 5, 60, 31, 2, 439, 440, 5,
	46, 24, 4, 440, 463, 3, 2, 2, 2, 441, 442, 12, 13, 2, 2, 442, 443, 7, 43,
	2, 2, 443, 463, 7, 67, 2, 2, 444, 445, 12, 12, 2, 2, 445, 454, 7, 58, 2,
	2, 446, 451, 5, 46, 24, 2, 447, 448, 7, 40, 2, 2, 448, 450, 5, 46, 24,
	2, 449, 447, 3, 2, 2, 2, 450, 453, 3, 2, 2, 2, 451, 449, 3, 2, 2, 2, 451,
	452, 3, 2, 2, 2, 452, 455, 3, 2, 2, 2, 453, 451, 3, 2, 2, 2, 454, 446,
	3, 2, 2, 2, 454, 455, 3, 2, 2, 2, 455, 456, 3, 2, 2, 2, 456, 463, 7, 59,
	2, 2, 457, 458, 12, 11, 2, 2, 458, 459, 7, 64, 2, 2, 459, 460, 5, 46, 24,
	2, 460, 461, 7, 65, 2, 2, 461, 463, 3, 2, 2, 2, 462, 431, 3, 2, 2, 2, 462,
	437, 3, 2, 2, 2, 462, 441, 3, 2, 2, 2, 462, 444, 3, 2, 2, 2, 462, 457,
	3, 2, 2, 2, 463, 466, 3, 2, 2, 2, 464, 462, 3, 2, 2, 2, 464, 465, 3, 2,
	2, 2, 465, 47, 3, 2, 2, 2, 466, 464, 3, 2, 2, 2, 467, 469, 7, 17, 2, 2,
	468, 470, 5, 38, 20, 2, 469, 468, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470,
	471, 3, 2, 2, 2, 471, 473, 5, 14, 8, 2, 472, 474, 5, 28, 15, 2, 473, 472,
	3, 2, 2, 2, 473, 474, 3, 2, 2, 2, 474, 478, 3, 2, 2, 2, 475, 477, 5, 36,
	19, 2, 476, 475, 3, 2, 2, 2, 477, 480, 3, 2, 2, 2, 478, 476, 3, 2, 2, 2,
	478, 479, 3, 2, 2, 2, 479, 481, 3, 2, 2, 2, 480, 478, 3, 2, 2, 2, 481,
	482, 5, 18, 10, 2, 482, 49, 3, 2, 2, 2, 483, 484, 5, 30, 16, 2, 484, 485,
	7, 41, 2, 2, 485, 486, 5, 52, 27, 2, 486, 487, 7, 42, 2, 2, 487, 51, 3,
	2, 2, 2, 488, 489, 5, 54, 28, 2, 489, 490, 7, 40, 2, 2, 490, 492, 3, 2,
	2, 2, 491, 488, 3, 2, 2, 2, 492, 495, 3, 2, 2, 2, 493, 491, 3, 2, 2, 2,
	493, 494, 3, 2, 2, 2, 494, 496, 3, 2, 2, 2, 495, 493, 3, 2, 2, 2, 496,
	498, 5, 54, 28, 2, 497, 493, 3, 2, 2, 2, 497, 498, 3, 2, 2, 2, 498, 53,
	3, 2, 2, 2, 499, 500, 7, 67, 2, 2, 500, 501, 7, 39, 2, 2, 501, 502, 5,
	46, 24, 2, 502, 55, 3, 2, 2, 2, 503, 504, 9, 2, 2, 2, 504, 57, 3, 2, 2,
	2, 505, 506, 9, 3, 2, 2, 506, 59, 3, 2, 2, 2, 507, 508, 9, 4, 2, 2, 508,
	61, 3, 2, 2, 2, 509, 510, 7, 64, 2, 2, 510, 511, 5, 50, 26, 2, 511, 512,
	7, 65, 2, 2, 512, 63, 3, 2, 2, 2, 513, 518, 7, 67, 2, 2, 514, 515, 7, 43,
	2, 2, 515, 517, 7, 67, 2, 2, 516, 514, 3, 2, 2, 2, 517, 520, 3, 2, 2, 2,
	518, 516, 3, 2, 2, 2, 518, 519, 3, 2, 2, 2, 519, 65, 3, 2, 2, 2, 520, 518,
	3, 2, 2, 2, 521, 525, 7, 3, 2, 2, 522, 525, 7, 4, 2, 2, 523, 525, 7, 5,
	2, 2, 524, 521, 3, 2, 2, 2, 524, 522, 3, 2, 2, 2, 524, 523, 3, 2, 2, 2,
	525, 67, 3, 2, 2, 2, 69, 69, 74, 80, 88, 98, 105, 113, 118, 122, 125, 128,
	132, 138, 143, 148, 159, 162, 167, 175, 184, 188, 193, 201, 211, 216, 224,
	226, 244, 252, 259, 262, 269, 272, 277, 283, 289, 293, 305, 312, 318, 324,
	328, 337, 342, 350, 353, 358, 364, 369, 373, 384, 396, 401, 409, 413, 429,
	451, 454, 462, 464, 469, 473, 478, 493, 497, 518, 524,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "'as'", "'bool'", "'break'", "'byte'", "'case'", "'char'",
	"'continue'", "'default'", "'else'", "'extern'", "'false'", "'fn'", "'if'",
	"'import'", "'interface'", "'int32'", "'module'", "'public'", "'pure'",
	"'return'", "'string'", "'struct'", "'switch'", "'true'", "'type'", "'value'",
	"'var'", "'where'", "'while'", "'&&'", "'*'", "'@'", "'^'", "':'", "','",
	"'{'", "'}'", "'.'", "'...'", "'='", "'=='", "'>'", "'>='", "'<'", "'<='",
	"'+'", "'-'", "'--'", "'%'", "'!'", "'!='", "'||'", "'('", "')'", "'++'",
	"'?'", "';'", "'/'", "'['", "']'", "'~'",
}
var symbolicNames = []string{
	"", "IntegerLiteral", "BoolLiteral", "StringLiteral", "As", "Bool", "Break",
	"Byte", "Case", "Char", "Continue", "Default", "Else", "Extern", "False",
	"Fn", "If", "Import", "Interface", "Int32", "Module", "Public", "Pure",
	"Return", "Str", "Struct", "Switch", "True", "Type", "Value", "Var", "Where",
	"While", "AndAnd", "Asterisk", "At", "Caret", "Colon", "Comma", "CurlyOpen",
	"CurlyClose", "Dot", "Ellipsis", "Equals", "EqualsEquals", "GreaterThan",
	"GreaterThanOrEqualTo", "LessThan", "LessThanOrEqualTo", "Plus", "Minus",
	"MinusMinus", "Modulus", "Not", "NotEquals", "OrOr", "ParenOpen", "ParenClose",
	"PlusPlus", "Question", "SemiColon", "Slash", "SquareOpen", "SquareClose",
	"Tilda", "Identifier", "LineComment", "BlockComment", "WHITESPACE",
}

var ruleNames = []string{
	"compilationUnit", "moduleDeclaration", "importStatement", "moduleReference",
	"declaration", "functionDeclaration", "functionParameters", "functionParameter",
	"codeBlock", "statement", "caseStatement", "defaultStatement", "typeDeclaration",
	"typeSpec", "referencedType", "extendsClause", "method", "whereClause",
	"typeParameters", "typeParameter", "primitiveType", "field", "expression",
	"lambdaExpr", "structExpr", "fieldAssignments", "fieldAssignment", "prefixOp",
	"postfixOp", "binaryOp", "attribute", "qualifiedName", "literal",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type NebularisParser struct {
	*antlr.BaseParser
}

func NewNebularisParser(input antlr.TokenStream) *NebularisParser {
	this := new(NebularisParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Nebularis.g4"

	return this
}

// NebularisParser tokens.
const (
	NebularisParserEOF                  = antlr.TokenEOF
	NebularisParserIntegerLiteral       = 1
	NebularisParserBoolLiteral          = 2
	NebularisParserStringLiteral        = 3
	NebularisParserAs                   = 4
	NebularisParserBool                 = 5
	NebularisParserBreak                = 6
	NebularisParserByte                 = 7
	NebularisParserCase                 = 8
	NebularisParserChar                 = 9
	NebularisParserContinue             = 10
	NebularisParserDefault              = 11
	NebularisParserElse                 = 12
	NebularisParserExtern               = 13
	NebularisParserFalse                = 14
	NebularisParserFn                   = 15
	NebularisParserIf                   = 16
	NebularisParserImport               = 17
	NebularisParserInterface            = 18
	NebularisParserInt32                = 19
	NebularisParserModule               = 20
	NebularisParserPublic               = 21
	NebularisParserPure                 = 22
	NebularisParserReturn               = 23
	NebularisParserStr                  = 24
	NebularisParserStruct               = 25
	NebularisParserSwitch               = 26
	NebularisParserTrue                 = 27
	NebularisParserType                 = 28
	NebularisParserValue                = 29
	NebularisParserVar                  = 30
	NebularisParserWhere                = 31
	NebularisParserWhile                = 32
	NebularisParserAndAnd               = 33
	NebularisParserAsterisk             = 34
	NebularisParserAt                   = 35
	NebularisParserCaret                = 36
	NebularisParserColon                = 37
	NebularisParserComma                = 38
	NebularisParserCurlyOpen            = 39
	NebularisParserCurlyClose           = 40
	NebularisParserDot                  = 41
	NebularisParserEllipsis             = 42
	NebularisParserEquals               = 43
	NebularisParserEqualsEquals         = 44
	NebularisParserGreaterThan          = 45
	NebularisParserGreaterThanOrEqualTo = 46
	NebularisParserLessThan             = 47
	NebularisParserLessThanOrEqualTo    = 48
	NebularisParserPlus                 = 49
	NebularisParserMinus                = 50
	NebularisParserMinusMinus           = 51
	NebularisParserModulus              = 52
	NebularisParserNot                  = 53
	NebularisParserNotEquals            = 54
	NebularisParserOrOr                 = 55
	NebularisParserParenOpen            = 56
	NebularisParserParenClose           = 57
	NebularisParserPlusPlus             = 58
	NebularisParserQuestion             = 59
	NebularisParserSemiColon            = 60
	NebularisParserSlash                = 61
	NebularisParserSquareOpen           = 62
	NebularisParserSquareClose          = 63
	NebularisParserTilda                = 64
	NebularisParserIdentifier           = 65
	NebularisParserLineComment          = 66
	NebularisParserBlockComment         = 67
	NebularisParserWHITESPACE           = 68
)

// NebularisParser rules.
const (
	NebularisParserRULE_compilationUnit     = 0
	NebularisParserRULE_moduleDeclaration   = 1
	NebularisParserRULE_importStatement     = 2
	NebularisParserRULE_moduleReference     = 3
	NebularisParserRULE_declaration         = 4
	NebularisParserRULE_functionDeclaration = 5
	NebularisParserRULE_functionParameters  = 6
	NebularisParserRULE_functionParameter   = 7
	NebularisParserRULE_codeBlock           = 8
	NebularisParserRULE_statement           = 9
	NebularisParserRULE_caseStatement       = 10
	NebularisParserRULE_defaultStatement    = 11
	NebularisParserRULE_typeDeclaration     = 12
	NebularisParserRULE_typeSpec            = 13
	NebularisParserRULE_referencedType      = 14
	NebularisParserRULE_extendsClause       = 15
	NebularisParserRULE_method              = 16
	NebularisParserRULE_whereClause         = 17
	NebularisParserRULE_typeParameters      = 18
	NebularisParserRULE_typeParameter       = 19
	NebularisParserRULE_primitiveType       = 20
	NebularisParserRULE_field               = 21
	NebularisParserRULE_expression          = 22
	NebularisParserRULE_lambdaExpr          = 23
	NebularisParserRULE_structExpr          = 24
	NebularisParserRULE_fieldAssignments    = 25
	NebularisParserRULE_fieldAssignment     = 26
	NebularisParserRULE_prefixOp            = 27
	NebularisParserRULE_postfixOp           = 28
	NebularisParserRULE_binaryOp            = 29
	NebularisParserRULE_attribute           = 30
	NebularisParserRULE_qualifiedName       = 31
	NebularisParserRULE_literal             = 32
)

// ICompilationUnitContext is an interface to support dynamic dispatch.
type ICompilationUnitContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetModule returns the module rule contexts.
	GetModule() IModuleDeclarationContext

	// Get_importStatement returns the _importStatement rule contexts.
	Get_importStatement() IImportStatementContext

	// Get_declaration returns the _declaration rule contexts.
	Get_declaration() IDeclarationContext

	// SetModule sets the module rule contexts.
	SetModule(IModuleDeclarationContext)

	// Set_importStatement sets the _importStatement rule contexts.
	Set_importStatement(IImportStatementContext)

	// Set_declaration sets the _declaration rule contexts.
	Set_declaration(IDeclarationContext)

	// GetImports returns the imports rule context list.
	GetImports() []IImportStatementContext

	// GetDeclarations returns the declarations rule context list.
	GetDeclarations() []IDeclarationContext

	// SetImports sets the imports rule context list.
	SetImports([]IImportStatementContext)

	// SetDeclarations sets the declarations rule context list.
	SetDeclarations([]IDeclarationContext)

	// IsCompilationUnitContext differentiates from other interfaces.
	IsCompilationUnitContext()
}

type CompilationUnitContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	module           IModuleDeclarationContext
	_importStatement IImportStatementContext
	imports          []IImportStatementContext
	_declaration     IDeclarationContext
	declarations     []IDeclarationContext
}

func NewEmptyCompilationUnitContext() *CompilationUnitContext {
	var p = new(CompilationUnitContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_compilationUnit
	return p
}

func (*CompilationUnitContext) IsCompilationUnitContext() {}

func NewCompilationUnitContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CompilationUnitContext {
	var p = new(CompilationUnitContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_compilationUnit

	return p
}

func (s *CompilationUnitContext) GetParser() antlr.Parser { return s.parser }

func (s *CompilationUnitContext) GetModule() IModuleDeclarationContext { return s.module }

func (s *CompilationUnitContext) Get_importStatement() IImportStatementContext {
	return s._importStatement
}

func (s *CompilationUnitContext) Get_declaration() IDeclarationContext { return s._declaration }

func (s *CompilationUnitContext) SetModule(v IModuleDeclarationContext) { s.module = v }

func (s *CompilationUnitContext) Set_importStatement(v IImportStatementContext) {
	s._importStatement = v
}

func (s *CompilationUnitContext) Set_declaration(v IDeclarationContext) { s._declaration = v }

func (s *CompilationUnitContext) GetImports() []IImportStatementContext { return s.imports }

func (s *CompilationUnitContext) GetDeclarations() []IDeclarationContext { return s.declarations }

func (s *CompilationUnitContext) SetImports(v []IImportStatementContext) { s.imports = v }

func (s *CompilationUnitContext) SetDeclarations(v []IDeclarationContext) { s.declarations = v }

func (s *CompilationUnitContext) EOF() antlr.TerminalNode {
	return s.GetToken(NebularisParserEOF, 0)
}

func (s *CompilationUnitContext) ModuleDeclaration() IModuleDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleDeclarationContext)
}

func (s *CompilationUnitContext) AllImportStatement() []IImportStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IImportStatementContext)(nil)).Elem())
	var tst = make([]IImportStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IImportStatementContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) ImportStatement(i int) IImportStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImportStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IImportStatementContext)
}

func (s *CompilationUnitContext) AllDeclaration() []IDeclarationContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDeclarationContext)(nil)).Elem())
	var tst = make([]IDeclarationContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDeclarationContext)
		}
	}

	return tst
}

func (s *CompilationUnitContext) Declaration(i int) IDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclarationContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDeclarationContext)
}

func (s *CompilationUnitContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CompilationUnitContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CompilationUnitContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitCompilationUnit(s)
	}
}

func (s *CompilationUnitContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitCompilationUnit(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) CompilationUnit() (localctx ICompilationUnitContext) {
	localctx = NewCompilationUnitContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, NebularisParserRULE_compilationUnit)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(67)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 0, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(66)

			var _x = p.ModuleDeclaration()

			localctx.(*CompilationUnitContext).module = _x
		}

	}
	p.SetState(72)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(69)

				var _x = p.ImportStatement()

				localctx.(*CompilationUnitContext)._importStatement = _x
			}
			localctx.(*CompilationUnitContext).imports = append(localctx.(*CompilationUnitContext).imports, localctx.(*CompilationUnitContext)._importStatement)

		}
		p.SetState(74)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext())
	}
	p.SetState(78)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NebularisParserExtern)|(1<<NebularisParserFn)|(1<<NebularisParserPublic)|(1<<NebularisParserPure)|(1<<NebularisParserType))) != 0) || _la == NebularisParserSquareOpen {
		{
			p.SetState(75)

			var _x = p.Declaration()

			localctx.(*CompilationUnitContext)._declaration = _x
		}
		localctx.(*CompilationUnitContext).declarations = append(localctx.(*CompilationUnitContext).declarations, localctx.(*CompilationUnitContext)._declaration)

		p.SetState(80)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(81)
		p.Match(NebularisParserEOF)
	}

	return localctx
}

// IModuleDeclarationContext is an interface to support dynamic dispatch.
type IModuleDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_attribute returns the _attribute rule contexts.
	Get_attribute() IAttributeContext

	// GetName returns the name rule contexts.
	GetName() IQualifiedNameContext

	// Set_attribute sets the _attribute rule contexts.
	Set_attribute(IAttributeContext)

	// SetName sets the name rule contexts.
	SetName(IQualifiedNameContext)

	// GetAttributes returns the attributes rule context list.
	GetAttributes() []IAttributeContext

	// SetAttributes sets the attributes rule context list.
	SetAttributes([]IAttributeContext)

	// IsModuleDeclarationContext differentiates from other interfaces.
	IsModuleDeclarationContext()
}

type ModuleDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_attribute IAttributeContext
	attributes []IAttributeContext
	name       IQualifiedNameContext
}

func NewEmptyModuleDeclarationContext() *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_moduleDeclaration
	return p
}

func (*ModuleDeclarationContext) IsModuleDeclarationContext() {}

func NewModuleDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleDeclarationContext {
	var p = new(ModuleDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_moduleDeclaration

	return p
}

func (s *ModuleDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleDeclarationContext) Get_attribute() IAttributeContext { return s._attribute }

func (s *ModuleDeclarationContext) GetName() IQualifiedNameContext { return s.name }

func (s *ModuleDeclarationContext) Set_attribute(v IAttributeContext) { s._attribute = v }

func (s *ModuleDeclarationContext) SetName(v IQualifiedNameContext) { s.name = v }

func (s *ModuleDeclarationContext) GetAttributes() []IAttributeContext { return s.attributes }

func (s *ModuleDeclarationContext) SetAttributes(v []IAttributeContext) { s.attributes = v }

func (s *ModuleDeclarationContext) Module() antlr.TerminalNode {
	return s.GetToken(NebularisParserModule, 0)
}

func (s *ModuleDeclarationContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *ModuleDeclarationContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *ModuleDeclarationContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *ModuleDeclarationContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ModuleDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitModuleDeclaration(s)
	}
}

func (s *ModuleDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitModuleDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) ModuleDeclaration() (localctx IModuleDeclarationContext) {
	localctx = NewModuleDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, NebularisParserRULE_moduleDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(86)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserSquareOpen {
		{
			p.SetState(83)

			var _x = p.Attribute()

			localctx.(*ModuleDeclarationContext)._attribute = _x
		}
		localctx.(*ModuleDeclarationContext).attributes = append(localctx.(*ModuleDeclarationContext).attributes, localctx.(*ModuleDeclarationContext)._attribute)

		p.SetState(88)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(89)
		p.Match(NebularisParserModule)
	}
	{
		p.SetState(90)

		var _x = p.QualifiedName()

		localctx.(*ModuleDeclarationContext).name = _x
	}
	{
		p.SetState(91)
		p.Match(NebularisParserSemiColon)
	}

	return localctx
}

// IImportStatementContext is an interface to support dynamic dispatch.
type IImportStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetAlias returns the alias token.
	GetAlias() antlr.Token

	// SetAlias sets the alias token.
	SetAlias(antlr.Token)

	// Get_attribute returns the _attribute rule contexts.
	Get_attribute() IAttributeContext

	// GetModule returns the module rule contexts.
	GetModule() IModuleReferenceContext

	// Set_attribute sets the _attribute rule contexts.
	Set_attribute(IAttributeContext)

	// SetModule sets the module rule contexts.
	SetModule(IModuleReferenceContext)

	// GetAttributes returns the attributes rule context list.
	GetAttributes() []IAttributeContext

	// SetAttributes sets the attributes rule context list.
	SetAttributes([]IAttributeContext)

	// IsImportStatementContext differentiates from other interfaces.
	IsImportStatementContext()
}

type ImportStatementContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_attribute IAttributeContext
	attributes []IAttributeContext
	module     IModuleReferenceContext
	alias      antlr.Token
}

func NewEmptyImportStatementContext() *ImportStatementContext {
	var p = new(ImportStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_importStatement
	return p
}

func (*ImportStatementContext) IsImportStatementContext() {}

func NewImportStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImportStatementContext {
	var p = new(ImportStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_importStatement

	return p
}

func (s *ImportStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *ImportStatementContext) GetAlias() antlr.Token { return s.alias }

func (s *ImportStatementContext) SetAlias(v antlr.Token) { s.alias = v }

func (s *ImportStatementContext) Get_attribute() IAttributeContext { return s._attribute }

func (s *ImportStatementContext) GetModule() IModuleReferenceContext { return s.module }

func (s *ImportStatementContext) Set_attribute(v IAttributeContext) { s._attribute = v }

func (s *ImportStatementContext) SetModule(v IModuleReferenceContext) { s.module = v }

func (s *ImportStatementContext) GetAttributes() []IAttributeContext { return s.attributes }

func (s *ImportStatementContext) SetAttributes(v []IAttributeContext) { s.attributes = v }

func (s *ImportStatementContext) Import() antlr.TerminalNode {
	return s.GetToken(NebularisParserImport, 0)
}

func (s *ImportStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *ImportStatementContext) ModuleReference() IModuleReferenceContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IModuleReferenceContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IModuleReferenceContext)
}

func (s *ImportStatementContext) As() antlr.TerminalNode {
	return s.GetToken(NebularisParserAs, 0)
}

func (s *ImportStatementContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *ImportStatementContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *ImportStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *ImportStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImportStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImportStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterImportStatement(s)
	}
}

func (s *ImportStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitImportStatement(s)
	}
}

func (s *ImportStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitImportStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) ImportStatement() (localctx IImportStatementContext) {
	localctx = NewImportStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, NebularisParserRULE_importStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(96)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserSquareOpen {
		{
			p.SetState(93)

			var _x = p.Attribute()

			localctx.(*ImportStatementContext)._attribute = _x
		}
		localctx.(*ImportStatementContext).attributes = append(localctx.(*ImportStatementContext).attributes, localctx.(*ImportStatementContext)._attribute)

		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(99)
		p.Match(NebularisParserImport)
	}
	{
		p.SetState(100)

		var _x = p.ModuleReference()

		localctx.(*ImportStatementContext).module = _x
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserAs {
		{
			p.SetState(101)
			p.Match(NebularisParserAs)
		}
		{
			p.SetState(102)

			var _m = p.Match(NebularisParserIdentifier)

			localctx.(*ImportStatementContext).alias = _m
		}

	}
	{
		p.SetState(105)
		p.Match(NebularisParserSemiColon)
	}

	return localctx
}

// IModuleReferenceContext is an interface to support dynamic dispatch.
type IModuleReferenceContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IQualifiedNameContext

	// SetName sets the name rule contexts.
	SetName(IQualifiedNameContext)

	// IsModuleReferenceContext differentiates from other interfaces.
	IsModuleReferenceContext()
}

type ModuleReferenceContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   IQualifiedNameContext
}

func NewEmptyModuleReferenceContext() *ModuleReferenceContext {
	var p = new(ModuleReferenceContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_moduleReference
	return p
}

func (*ModuleReferenceContext) IsModuleReferenceContext() {}

func NewModuleReferenceContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ModuleReferenceContext {
	var p = new(ModuleReferenceContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_moduleReference

	return p
}

func (s *ModuleReferenceContext) GetParser() antlr.Parser { return s.parser }

func (s *ModuleReferenceContext) GetName() IQualifiedNameContext { return s.name }

func (s *ModuleReferenceContext) SetName(v IQualifiedNameContext) { s.name = v }

func (s *ModuleReferenceContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *ModuleReferenceContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ModuleReferenceContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ModuleReferenceContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterModuleReference(s)
	}
}

func (s *ModuleReferenceContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitModuleReference(s)
	}
}

func (s *ModuleReferenceContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitModuleReference(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) ModuleReference() (localctx IModuleReferenceContext) {
	localctx = NewModuleReferenceContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, NebularisParserRULE_moduleReference)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(107)

		var _x = p.QualifiedName()

		localctx.(*ModuleReferenceContext).name = _x
	}

	return localctx
}

// IDeclarationContext is an interface to support dynamic dispatch.
type IDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeDecl returns the typeDecl rule contexts.
	GetTypeDecl() ITypeDeclarationContext

	// GetFnDecl returns the fnDecl rule contexts.
	GetFnDecl() IFunctionDeclarationContext

	// SetTypeDecl sets the typeDecl rule contexts.
	SetTypeDecl(ITypeDeclarationContext)

	// SetFnDecl sets the fnDecl rule contexts.
	SetFnDecl(IFunctionDeclarationContext)

	// IsDeclarationContext differentiates from other interfaces.
	IsDeclarationContext()
}

type DeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser   antlr.Parser
	typeDecl ITypeDeclarationContext
	fnDecl   IFunctionDeclarationContext
}

func NewEmptyDeclarationContext() *DeclarationContext {
	var p = new(DeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_declaration
	return p
}

func (*DeclarationContext) IsDeclarationContext() {}

func NewDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclarationContext {
	var p = new(DeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_declaration

	return p
}

func (s *DeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclarationContext) GetTypeDecl() ITypeDeclarationContext { return s.typeDecl }

func (s *DeclarationContext) GetFnDecl() IFunctionDeclarationContext { return s.fnDecl }

func (s *DeclarationContext) SetTypeDecl(v ITypeDeclarationContext) { s.typeDecl = v }

func (s *DeclarationContext) SetFnDecl(v IFunctionDeclarationContext) { s.fnDecl = v }

func (s *DeclarationContext) TypeDeclaration() ITypeDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeDeclarationContext)
}

func (s *DeclarationContext) FunctionDeclaration() IFunctionDeclarationContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionDeclarationContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionDeclarationContext)
}

func (s *DeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterDeclaration(s)
	}
}

func (s *DeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitDeclaration(s)
	}
}

func (s *DeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Declaration() (localctx IDeclarationContext) {
	localctx = NewDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, NebularisParserRULE_declaration)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(111)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)

			var _x = p.TypeDeclaration()

			localctx.(*DeclarationContext).typeDecl = _x
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)

			var _x = p.FunctionDeclaration()

			localctx.(*DeclarationContext).fnDecl = _x
		}

	}

	return localctx
}

// IFunctionDeclarationContext is an interface to support dynamic dispatch.
type IFunctionDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExtern returns the extern token.
	GetExtern() antlr.Token

	// GetPub returns the pub token.
	GetPub() antlr.Token

	// GetPure returns the pure token.
	GetPure() antlr.Token

	// GetName returns the name token.
	GetName() antlr.Token

	// SetExtern sets the extern token.
	SetExtern(antlr.Token)

	// SetPub sets the pub token.
	SetPub(antlr.Token)

	// SetPure sets the pure token.
	SetPure(antlr.Token)

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_attribute returns the _attribute rule contexts.
	Get_attribute() IAttributeContext

	// GetTypeParams returns the typeParams rule contexts.
	GetTypeParams() ITypeParametersContext

	// GetTarget returns the target rule contexts.
	GetTarget() IReferencedTypeContext

	// GetParams returns the params rule contexts.
	GetParams() IFunctionParametersContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeSpecContext

	// Get_whereClause returns the _whereClause rule contexts.
	Get_whereClause() IWhereClauseContext

	// GetBody returns the body rule contexts.
	GetBody() ICodeBlockContext

	// Set_attribute sets the _attribute rule contexts.
	Set_attribute(IAttributeContext)

	// SetTypeParams sets the typeParams rule contexts.
	SetTypeParams(ITypeParametersContext)

	// SetTarget sets the target rule contexts.
	SetTarget(IReferencedTypeContext)

	// SetParams sets the params rule contexts.
	SetParams(IFunctionParametersContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeSpecContext)

	// Set_whereClause sets the _whereClause rule contexts.
	Set_whereClause(IWhereClauseContext)

	// SetBody sets the body rule contexts.
	SetBody(ICodeBlockContext)

	// GetAttributes returns the attributes rule context list.
	GetAttributes() []IAttributeContext

	// GetConstraints returns the constraints rule context list.
	GetConstraints() []IWhereClauseContext

	// SetAttributes sets the attributes rule context list.
	SetAttributes([]IAttributeContext)

	// SetConstraints sets the constraints rule context list.
	SetConstraints([]IWhereClauseContext)

	// IsFunctionDeclarationContext differentiates from other interfaces.
	IsFunctionDeclarationContext()
}

type FunctionDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_attribute   IAttributeContext
	attributes   []IAttributeContext
	extern       antlr.Token
	pub          antlr.Token
	pure         antlr.Token
	typeParams   ITypeParametersContext
	target       IReferencedTypeContext
	name         antlr.Token
	params       IFunctionParametersContext
	returnType   ITypeSpecContext
	_whereClause IWhereClauseContext
	constraints  []IWhereClauseContext
	body         ICodeBlockContext
}

func NewEmptyFunctionDeclarationContext() *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_functionDeclaration
	return p
}

func (*FunctionDeclarationContext) IsFunctionDeclarationContext() {}

func NewFunctionDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionDeclarationContext {
	var p = new(FunctionDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_functionDeclaration

	return p
}

func (s *FunctionDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionDeclarationContext) GetExtern() antlr.Token { return s.extern }

func (s *FunctionDeclarationContext) GetPub() antlr.Token { return s.pub }

func (s *FunctionDeclarationContext) GetPure() antlr.Token { return s.pure }

func (s *FunctionDeclarationContext) GetName() antlr.Token { return s.name }

func (s *FunctionDeclarationContext) SetExtern(v antlr.Token) { s.extern = v }

func (s *FunctionDeclarationContext) SetPub(v antlr.Token) { s.pub = v }

func (s *FunctionDeclarationContext) SetPure(v antlr.Token) { s.pure = v }

func (s *FunctionDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionDeclarationContext) Get_attribute() IAttributeContext { return s._attribute }

func (s *FunctionDeclarationContext) GetTypeParams() ITypeParametersContext { return s.typeParams }

func (s *FunctionDeclarationContext) GetTarget() IReferencedTypeContext { return s.target }

func (s *FunctionDeclarationContext) GetParams() IFunctionParametersContext { return s.params }

func (s *FunctionDeclarationContext) GetReturnType() ITypeSpecContext { return s.returnType }

func (s *FunctionDeclarationContext) Get_whereClause() IWhereClauseContext { return s._whereClause }

func (s *FunctionDeclarationContext) GetBody() ICodeBlockContext { return s.body }

func (s *FunctionDeclarationContext) Set_attribute(v IAttributeContext) { s._attribute = v }

func (s *FunctionDeclarationContext) SetTypeParams(v ITypeParametersContext) { s.typeParams = v }

func (s *FunctionDeclarationContext) SetTarget(v IReferencedTypeContext) { s.target = v }

func (s *FunctionDeclarationContext) SetParams(v IFunctionParametersContext) { s.params = v }

func (s *FunctionDeclarationContext) SetReturnType(v ITypeSpecContext) { s.returnType = v }

func (s *FunctionDeclarationContext) Set_whereClause(v IWhereClauseContext) { s._whereClause = v }

func (s *FunctionDeclarationContext) SetBody(v ICodeBlockContext) { s.body = v }

func (s *FunctionDeclarationContext) GetAttributes() []IAttributeContext { return s.attributes }

func (s *FunctionDeclarationContext) GetConstraints() []IWhereClauseContext { return s.constraints }

func (s *FunctionDeclarationContext) SetAttributes(v []IAttributeContext) { s.attributes = v }

func (s *FunctionDeclarationContext) SetConstraints(v []IWhereClauseContext) { s.constraints = v }

func (s *FunctionDeclarationContext) Fn() antlr.TerminalNode {
	return s.GetToken(NebularisParserFn, 0)
}

func (s *FunctionDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *FunctionDeclarationContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionDeclarationContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *FunctionDeclarationContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenOpen, 0)
}

func (s *FunctionDeclarationContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenClose, 0)
}

func (s *FunctionDeclarationContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *FunctionDeclarationContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *FunctionDeclarationContext) Extern() antlr.TerminalNode {
	return s.GetToken(NebularisParserExtern, 0)
}

func (s *FunctionDeclarationContext) Public() antlr.TerminalNode {
	return s.GetToken(NebularisParserPublic, 0)
}

func (s *FunctionDeclarationContext) Pure() antlr.TerminalNode {
	return s.GetToken(NebularisParserPure, 0)
}

func (s *FunctionDeclarationContext) TypeParameters() ITypeParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *FunctionDeclarationContext) ReferencedType() IReferencedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferencedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferencedTypeContext)
}

func (s *FunctionDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FunctionDeclarationContext) AllWhereClause() []IWhereClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem())
	var tst = make([]IWhereClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereClauseContext)
		}
	}

	return tst
}

func (s *FunctionDeclarationContext) WhereClause(i int) IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *FunctionDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitFunctionDeclaration(s)
	}
}

func (s *FunctionDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitFunctionDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) FunctionDeclaration() (localctx IFunctionDeclarationContext) {
	localctx = NewFunctionDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, NebularisParserRULE_functionDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(116)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserSquareOpen {
		{
			p.SetState(113)

			var _x = p.Attribute()

			localctx.(*FunctionDeclarationContext)._attribute = _x
		}
		localctx.(*FunctionDeclarationContext).attributes = append(localctx.(*FunctionDeclarationContext).attributes, localctx.(*FunctionDeclarationContext)._attribute)

		p.SetState(118)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(120)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserExtern {
		{
			p.SetState(119)

			var _m = p.Match(NebularisParserExtern)

			localctx.(*FunctionDeclarationContext).extern = _m
		}

	}
	p.SetState(123)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserPublic {
		{
			p.SetState(122)

			var _m = p.Match(NebularisParserPublic)

			localctx.(*FunctionDeclarationContext).pub = _m
		}

	}
	p.SetState(126)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserPure {
		{
			p.SetState(125)

			var _m = p.Match(NebularisParserPure)

			localctx.(*FunctionDeclarationContext).pure = _m
		}

	}
	{
		p.SetState(128)
		p.Match(NebularisParserFn)
	}
	p.SetState(130)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserLessThan {
		{
			p.SetState(129)

			var _x = p.TypeParameters()

			localctx.(*FunctionDeclarationContext).typeParams = _x
		}

	}
	p.SetState(136)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserParenOpen {
		{
			p.SetState(132)
			p.Match(NebularisParserParenOpen)
		}
		{
			p.SetState(133)

			var _x = p.ReferencedType()

			localctx.(*FunctionDeclarationContext).target = _x
		}
		{
			p.SetState(134)
			p.Match(NebularisParserParenClose)
		}

	}
	{
		p.SetState(138)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*FunctionDeclarationContext).name = _m
	}
	{
		p.SetState(139)

		var _x = p.FunctionParameters()

		localctx.(*FunctionDeclarationContext).params = _x
	}
	p.SetState(141)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NebularisParserBool)|(1<<NebularisParserByte)|(1<<NebularisParserChar)|(1<<NebularisParserFn)|(1<<NebularisParserInterface)|(1<<NebularisParserInt32)|(1<<NebularisParserStr)|(1<<NebularisParserStruct))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(NebularisParserQuestion-59))|(1<<(NebularisParserSquareOpen-59))|(1<<(NebularisParserIdentifier-59)))) != 0) {
		{
			p.SetState(140)

			var _x = p.TypeSpec()

			localctx.(*FunctionDeclarationContext).returnType = _x
		}

	}
	p.SetState(146)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserWhere {
		{
			p.SetState(143)

			var _x = p.WhereClause()

			localctx.(*FunctionDeclarationContext)._whereClause = _x
		}
		localctx.(*FunctionDeclarationContext).constraints = append(localctx.(*FunctionDeclarationContext).constraints, localctx.(*FunctionDeclarationContext)._whereClause)

		p.SetState(148)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(149)

		var _x = p.CodeBlock()

		localctx.(*FunctionDeclarationContext).body = _x
	}

	return localctx
}

// IFunctionParametersContext is an interface to support dynamic dispatch.
type IFunctionParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_functionParameter returns the _functionParameter rule contexts.
	Get_functionParameter() IFunctionParameterContext

	// Set_functionParameter sets the _functionParameter rule contexts.
	Set_functionParameter(IFunctionParameterContext)

	// GetParams returns the params rule context list.
	GetParams() []IFunctionParameterContext

	// SetParams sets the params rule context list.
	SetParams([]IFunctionParameterContext)

	// IsFunctionParametersContext differentiates from other interfaces.
	IsFunctionParametersContext()
}

type FunctionParametersContext struct {
	*antlr.BaseParserRuleContext
	parser             antlr.Parser
	_functionParameter IFunctionParameterContext
	params             []IFunctionParameterContext
}

func NewEmptyFunctionParametersContext() *FunctionParametersContext {
	var p = new(FunctionParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_functionParameters
	return p
}

func (*FunctionParametersContext) IsFunctionParametersContext() {}

func NewFunctionParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParametersContext {
	var p = new(FunctionParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_functionParameters

	return p
}

func (s *FunctionParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParametersContext) Get_functionParameter() IFunctionParameterContext {
	return s._functionParameter
}

func (s *FunctionParametersContext) Set_functionParameter(v IFunctionParameterContext) {
	s._functionParameter = v
}

func (s *FunctionParametersContext) GetParams() []IFunctionParameterContext { return s.params }

func (s *FunctionParametersContext) SetParams(v []IFunctionParameterContext) { s.params = v }

func (s *FunctionParametersContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenOpen, 0)
}

func (s *FunctionParametersContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenClose, 0)
}

func (s *FunctionParametersContext) AllFunctionParameter() []IFunctionParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem())
	var tst = make([]IFunctionParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFunctionParameterContext)
		}
	}

	return tst
}

func (s *FunctionParametersContext) FunctionParameter(i int) IFunctionParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFunctionParameterContext)
}

func (s *FunctionParametersContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *FunctionParametersContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *FunctionParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitFunctionParameters(s)
	}
}

func (s *FunctionParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitFunctionParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) FunctionParameters() (localctx IFunctionParametersContext) {
	localctx = NewFunctionParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, NebularisParserRULE_functionParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(151)
		p.Match(NebularisParserParenOpen)
	}
	p.SetState(160)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NebularisParserBool)|(1<<NebularisParserByte)|(1<<NebularisParserChar)|(1<<NebularisParserFn)|(1<<NebularisParserInterface)|(1<<NebularisParserInt32)|(1<<NebularisParserStr)|(1<<NebularisParserStruct))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(NebularisParserQuestion-59))|(1<<(NebularisParserSquareOpen-59))|(1<<(NebularisParserIdentifier-59)))) != 0) {
		{
			p.SetState(152)

			var _x = p.FunctionParameter()

			localctx.(*FunctionParametersContext)._functionParameter = _x
		}
		localctx.(*FunctionParametersContext).params = append(localctx.(*FunctionParametersContext).params, localctx.(*FunctionParametersContext)._functionParameter)
		p.SetState(157)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NebularisParserComma {
			{
				p.SetState(153)
				p.Match(NebularisParserComma)
			}
			{
				p.SetState(154)

				var _x = p.FunctionParameter()

				localctx.(*FunctionParametersContext)._functionParameter = _x
			}
			localctx.(*FunctionParametersContext).params = append(localctx.(*FunctionParametersContext).params, localctx.(*FunctionParametersContext)._functionParameter)

			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}
	{
		p.SetState(162)
		p.Match(NebularisParserParenClose)
	}

	return localctx
}

// IFunctionParameterContext is an interface to support dynamic dispatch.
type IFunctionParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetTypeRef returns the typeRef rule contexts.
	GetTypeRef() ITypeSpecContext

	// SetTypeRef sets the typeRef rule contexts.
	SetTypeRef(ITypeSpecContext)

	// IsFunctionParameterContext differentiates from other interfaces.
	IsFunctionParameterContext()
}

type FunctionParameterContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	name    antlr.Token
	typeRef ITypeSpecContext
}

func NewEmptyFunctionParameterContext() *FunctionParameterContext {
	var p = new(FunctionParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_functionParameter
	return p
}

func (*FunctionParameterContext) IsFunctionParameterContext() {}

func NewFunctionParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionParameterContext {
	var p = new(FunctionParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_functionParameter

	return p
}

func (s *FunctionParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionParameterContext) GetName() antlr.Token { return s.name }

func (s *FunctionParameterContext) SetName(v antlr.Token) { s.name = v }

func (s *FunctionParameterContext) GetTypeRef() ITypeSpecContext { return s.typeRef }

func (s *FunctionParameterContext) SetTypeRef(v ITypeSpecContext) { s.typeRef = v }

func (s *FunctionParameterContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FunctionParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *FunctionParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitFunctionParameter(s)
	}
}

func (s *FunctionParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitFunctionParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) FunctionParameter() (localctx IFunctionParameterContext) {
	localctx = NewFunctionParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, NebularisParserRULE_functionParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(165)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(164)

			var _m = p.Match(NebularisParserIdentifier)

			localctx.(*FunctionParameterContext).name = _m
		}

	}
	{
		p.SetState(167)

		var _x = p.TypeSpec()

		localctx.(*FunctionParameterContext).typeRef = _x
	}

	return localctx
}

// ICodeBlockContext is an interface to support dynamic dispatch.
type ICodeBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// IsCodeBlockContext differentiates from other interfaces.
	IsCodeBlockContext()
}

type CodeBlockContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_statement IStatementContext
	statements []IStatementContext
}

func NewEmptyCodeBlockContext() *CodeBlockContext {
	var p = new(CodeBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_codeBlock
	return p
}

func (*CodeBlockContext) IsCodeBlockContext() {}

func NewCodeBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CodeBlockContext {
	var p = new(CodeBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_codeBlock

	return p
}

func (s *CodeBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *CodeBlockContext) Get_statement() IStatementContext { return s._statement }

func (s *CodeBlockContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *CodeBlockContext) GetStatements() []IStatementContext { return s.statements }

func (s *CodeBlockContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *CodeBlockContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyOpen, 0)
}

func (s *CodeBlockContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyClose, 0)
}

func (s *CodeBlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CodeBlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CodeBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CodeBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CodeBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterCodeBlock(s)
	}
}

func (s *CodeBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitCodeBlock(s)
	}
}

func (s *CodeBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitCodeBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) CodeBlock() (localctx ICodeBlockContext) {
	localctx = NewCodeBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, NebularisParserRULE_codeBlock)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(169)
		p.Match(NebularisParserCurlyOpen)
	}
	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(NebularisParserIntegerLiteral-1))|(1<<(NebularisParserBoolLiteral-1))|(1<<(NebularisParserStringLiteral-1))|(1<<(NebularisParserBreak-1))|(1<<(NebularisParserContinue-1))|(1<<(NebularisParserFn-1))|(1<<(NebularisParserIf-1))|(1<<(NebularisParserReturn-1))|(1<<(NebularisParserSwitch-1))|(1<<(NebularisParserValue-1))|(1<<(NebularisParserVar-1))|(1<<(NebularisParserWhile-1)))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(NebularisParserEllipsis-42))|(1<<(NebularisParserNot-42))|(1<<(NebularisParserParenOpen-42))|(1<<(NebularisParserTilda-42))|(1<<(NebularisParserIdentifier-42)))) != 0) {
		{
			p.SetState(170)

			var _x = p.Statement()

			localctx.(*CodeBlockContext)._statement = _x
		}
		localctx.(*CodeBlockContext).statements = append(localctx.(*CodeBlockContext).statements, localctx.(*CodeBlockContext)._statement)

		p.SetState(175)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(176)
		p.Match(NebularisParserCurlyClose)
	}

	return localctx
}

// IStatementContext is an interface to support dynamic dispatch.
type IStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStatementContext differentiates from other interfaces.
	IsStatementContext()
}

type StatementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatementContext() *StatementContext {
	var p = new(StatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) CopyFrom(ctx *StatementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type BreakStatementContext struct {
	*StatementContext
	label antlr.Token
}

func NewBreakStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BreakStatementContext {
	var p = new(BreakStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *BreakStatementContext) GetLabel() antlr.Token { return s.label }

func (s *BreakStatementContext) SetLabel(v antlr.Token) { s.label = v }

func (s *BreakStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreakStatementContext) Break() antlr.TerminalNode {
	return s.GetToken(NebularisParserBreak, 0)
}

func (s *BreakStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *BreakStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *BreakStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterBreakStatement(s)
	}
}

func (s *BreakStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitBreakStatement(s)
	}
}

func (s *BreakStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitBreakStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type SwitchStatementContext struct {
	*StatementContext
	expr              IExpressionContext
	_caseStatement    ICaseStatementContext
	cases             []ICaseStatementContext
	_defaultStatement IDefaultStatementContext
	defauls           []IDefaultStatementContext
}

func NewSwitchStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SwitchStatementContext {
	var p = new(SwitchStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *SwitchStatementContext) GetExpr() IExpressionContext { return s.expr }

func (s *SwitchStatementContext) Get_caseStatement() ICaseStatementContext { return s._caseStatement }

func (s *SwitchStatementContext) Get_defaultStatement() IDefaultStatementContext {
	return s._defaultStatement
}

func (s *SwitchStatementContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *SwitchStatementContext) Set_caseStatement(v ICaseStatementContext) { s._caseStatement = v }

func (s *SwitchStatementContext) Set_defaultStatement(v IDefaultStatementContext) {
	s._defaultStatement = v
}

func (s *SwitchStatementContext) GetCases() []ICaseStatementContext { return s.cases }

func (s *SwitchStatementContext) GetDefauls() []IDefaultStatementContext { return s.defauls }

func (s *SwitchStatementContext) SetCases(v []ICaseStatementContext) { s.cases = v }

func (s *SwitchStatementContext) SetDefauls(v []IDefaultStatementContext) { s.defauls = v }

func (s *SwitchStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchStatementContext) Switch() antlr.TerminalNode {
	return s.GetToken(NebularisParserSwitch, 0)
}

func (s *SwitchStatementContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyOpen, 0)
}

func (s *SwitchStatementContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyClose, 0)
}

func (s *SwitchStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *SwitchStatementContext) AllCaseStatement() []ICaseStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem())
	var tst = make([]ICaseStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICaseStatementContext)
		}
	}

	return tst
}

func (s *SwitchStatementContext) CaseStatement(i int) ICaseStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICaseStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICaseStatementContext)
}

func (s *SwitchStatementContext) AllDefaultStatement() []IDefaultStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDefaultStatementContext)(nil)).Elem())
	var tst = make([]IDefaultStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDefaultStatementContext)
		}
	}

	return tst
}

func (s *SwitchStatementContext) DefaultStatement(i int) IDefaultStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDefaultStatementContext)
}

func (s *SwitchStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitSwitchStatement(s)
	}
}

func (s *SwitchStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitSwitchStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type VariableStatementContext struct {
	*StatementContext
	name        antlr.Token
	typeRef     ITypeSpecContext
	initializer IExpressionContext
}

func NewVariableStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *VariableStatementContext {
	var p = new(VariableStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *VariableStatementContext) GetName() antlr.Token { return s.name }

func (s *VariableStatementContext) SetName(v antlr.Token) { s.name = v }

func (s *VariableStatementContext) GetTypeRef() ITypeSpecContext { return s.typeRef }

func (s *VariableStatementContext) GetInitializer() IExpressionContext { return s.initializer }

func (s *VariableStatementContext) SetTypeRef(v ITypeSpecContext) { s.typeRef = v }

func (s *VariableStatementContext) SetInitializer(v IExpressionContext) { s.initializer = v }

func (s *VariableStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VariableStatementContext) Var() antlr.TerminalNode {
	return s.GetToken(NebularisParserVar, 0)
}

func (s *VariableStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *VariableStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *VariableStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *VariableStatementContext) Equals() antlr.TerminalNode {
	return s.GetToken(NebularisParserEquals, 0)
}

func (s *VariableStatementContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *VariableStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *VariableStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterVariableStatement(s)
	}
}

func (s *VariableStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitVariableStatement(s)
	}
}

func (s *VariableStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitVariableStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type LabelStatementContext struct {
	*StatementContext
	label antlr.Token
}

func NewLabelStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LabelStatementContext {
	var p = new(LabelStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *LabelStatementContext) GetLabel() antlr.Token { return s.label }

func (s *LabelStatementContext) SetLabel(v antlr.Token) { s.label = v }

func (s *LabelStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LabelStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *LabelStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *LabelStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterLabelStatement(s)
	}
}

func (s *LabelStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitLabelStatement(s)
	}
}

func (s *LabelStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitLabelStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type AssignmentStatementContext struct {
	*StatementContext
	left  IExpressionContext
	right IExpressionContext
}

func NewAssignmentStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AssignmentStatementContext {
	var p = new(AssignmentStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *AssignmentStatementContext) GetLeft() IExpressionContext { return s.left }

func (s *AssignmentStatementContext) GetRight() IExpressionContext { return s.right }

func (s *AssignmentStatementContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *AssignmentStatementContext) SetRight(v IExpressionContext) { s.right = v }

func (s *AssignmentStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AssignmentStatementContext) Equals() antlr.TerminalNode {
	return s.GetToken(NebularisParserEquals, 0)
}

func (s *AssignmentStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *AssignmentStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *AssignmentStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *AssignmentStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitAssignmentStatement(s)
	}
}

func (s *AssignmentStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitAssignmentStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotYetImplementedStatementContext struct {
	*StatementContext
}

func NewNotYetImplementedStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotYetImplementedStatementContext {
	var p = new(NotYetImplementedStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *NotYetImplementedStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotYetImplementedStatementContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(NebularisParserEllipsis, 0)
}

func (s *NotYetImplementedStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterNotYetImplementedStatement(s)
	}
}

func (s *NotYetImplementedStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitNotYetImplementedStatement(s)
	}
}

func (s *NotYetImplementedStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitNotYetImplementedStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConditionalStatementContext struct {
	*StatementContext
	cond    IExpressionContext
	onTrue  ICodeBlockContext
	onFalse ICodeBlockContext
}

func NewConditionalStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConditionalStatementContext {
	var p = new(ConditionalStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ConditionalStatementContext) GetCond() IExpressionContext { return s.cond }

func (s *ConditionalStatementContext) GetOnTrue() ICodeBlockContext { return s.onTrue }

func (s *ConditionalStatementContext) GetOnFalse() ICodeBlockContext { return s.onFalse }

func (s *ConditionalStatementContext) SetCond(v IExpressionContext) { s.cond = v }

func (s *ConditionalStatementContext) SetOnTrue(v ICodeBlockContext) { s.onTrue = v }

func (s *ConditionalStatementContext) SetOnFalse(v ICodeBlockContext) { s.onFalse = v }

func (s *ConditionalStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalStatementContext) If() antlr.TerminalNode {
	return s.GetToken(NebularisParserIf, 0)
}

func (s *ConditionalStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionalStatementContext) AllCodeBlock() []ICodeBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem())
	var tst = make([]ICodeBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ICodeBlockContext)
		}
	}

	return tst
}

func (s *ConditionalStatementContext) CodeBlock(i int) ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *ConditionalStatementContext) Else() antlr.TerminalNode {
	return s.GetToken(NebularisParserElse, 0)
}

func (s *ConditionalStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterConditionalStatement(s)
	}
}

func (s *ConditionalStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitConditionalStatement(s)
	}
}

func (s *ConditionalStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitConditionalStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryStatementContext struct {
	*StatementContext
	op IPostfixOpContext
}

func NewUnaryStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryStatementContext {
	var p = new(UnaryStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *UnaryStatementContext) GetOp() IPostfixOpContext { return s.op }

func (s *UnaryStatementContext) SetOp(v IPostfixOpContext) { s.op = v }

func (s *UnaryStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *UnaryStatementContext) PostfixOp() IPostfixOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPostfixOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPostfixOpContext)
}

func (s *UnaryStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterUnaryStatement(s)
	}
}

func (s *UnaryStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitUnaryStatement(s)
	}
}

func (s *UnaryStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitUnaryStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExpressionStatementContext struct {
	*StatementContext
}

func NewExpressionStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionStatementContext {
	var p = new(ExpressionStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ExpressionStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *ExpressionStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitExpressionStatement(s)
	}
}

func (s *ExpressionStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitExpressionStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReturnStatementContext struct {
	*StatementContext
	value IExpressionContext
}

func NewReturnStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReturnStatementContext {
	var p = new(ReturnStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ReturnStatementContext) GetValue() IExpressionContext { return s.value }

func (s *ReturnStatementContext) SetValue(v IExpressionContext) { s.value = v }

func (s *ReturnStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReturnStatementContext) Return() antlr.TerminalNode {
	return s.GetToken(NebularisParserReturn, 0)
}

func (s *ReturnStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *ReturnStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ReturnStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterReturnStatement(s)
	}
}

func (s *ReturnStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitReturnStatement(s)
	}
}

func (s *ReturnStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitReturnStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type WhileStatementContext struct {
	*StatementContext
}

func NewWhileStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *WhileStatementContext {
	var p = new(WhileStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *WhileStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhileStatementContext) While() antlr.TerminalNode {
	return s.GetToken(NebularisParserWhile, 0)
}

func (s *WhileStatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhileStatementContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *WhileStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterWhileStatement(s)
	}
}

func (s *WhileStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitWhileStatement(s)
	}
}

func (s *WhileStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitWhileStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

type ContinueStatementContext struct {
	*StatementContext
	label antlr.Token
}

func NewContinueStatementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ContinueStatementContext {
	var p = new(ContinueStatementContext)

	p.StatementContext = NewEmptyStatementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StatementContext))

	return p
}

func (s *ContinueStatementContext) GetLabel() antlr.Token { return s.label }

func (s *ContinueStatementContext) SetLabel(v antlr.Token) { s.label = v }

func (s *ContinueStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinueStatementContext) Continue() antlr.TerminalNode {
	return s.GetToken(NebularisParserContinue, 0)
}

func (s *ContinueStatementContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *ContinueStatementContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *ContinueStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterContinueStatement(s)
	}
}

func (s *ContinueStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitContinueStatement(s)
	}
}

func (s *ContinueStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitContinueStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, NebularisParserRULE_statement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(242)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		localctx = NewVariableStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(178)
			p.Match(NebularisParserVar)
		}
		{
			p.SetState(179)

			var _m = p.Match(NebularisParserIdentifier)

			localctx.(*VariableStatementContext).name = _m
		}
		p.SetState(182)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NebularisParserColon {
			{
				p.SetState(180)
				p.Match(NebularisParserColon)
			}
			{
				p.SetState(181)

				var _x = p.TypeSpec()

				localctx.(*VariableStatementContext).typeRef = _x
			}

		}
		p.SetState(186)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NebularisParserEquals {
			{
				p.SetState(184)
				p.Match(NebularisParserEquals)
			}
			{
				p.SetState(185)

				var _x = p.expression(0)

				localctx.(*VariableStatementContext).initializer = _x
			}

		}
		{
			p.SetState(188)
			p.Match(NebularisParserSemiColon)
		}

	case 2:
		localctx = NewReturnStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(189)
			p.Match(NebularisParserReturn)
		}
		p.SetState(191)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NebularisParserIntegerLiteral)|(1<<NebularisParserBoolLiteral)|(1<<NebularisParserStringLiteral)|(1<<NebularisParserFn)|(1<<NebularisParserValue))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(NebularisParserEllipsis-42))|(1<<(NebularisParserNot-42))|(1<<(NebularisParserParenOpen-42))|(1<<(NebularisParserTilda-42))|(1<<(NebularisParserIdentifier-42)))) != 0) {
			{
				p.SetState(190)

				var _x = p.expression(0)

				localctx.(*ReturnStatementContext).value = _x
			}

		}
		{
			p.SetState(193)
			p.Match(NebularisParserSemiColon)
		}

	case 3:
		localctx = NewConditionalStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(194)
			p.Match(NebularisParserIf)
		}
		{
			p.SetState(195)

			var _x = p.expression(0)

			localctx.(*ConditionalStatementContext).cond = _x
		}
		{
			p.SetState(196)

			var _x = p.CodeBlock()

			localctx.(*ConditionalStatementContext).onTrue = _x
		}
		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NebularisParserElse {
			{
				p.SetState(197)
				p.Match(NebularisParserElse)
			}
			{
				p.SetState(198)

				var _x = p.CodeBlock()

				localctx.(*ConditionalStatementContext).onFalse = _x
			}

		}

	case 4:
		localctx = NewWhileStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(201)
			p.Match(NebularisParserWhile)
		}
		{
			p.SetState(202)
			p.expression(0)
		}
		{
			p.SetState(203)
			p.CodeBlock()
		}

	case 5:
		localctx = NewLabelStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(205)

			var _m = p.Match(NebularisParserIdentifier)

			localctx.(*LabelStatementContext).label = _m
		}
		{
			p.SetState(206)
			p.Match(NebularisParserColon)
		}

	case 6:
		localctx = NewContinueStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(207)
			p.Match(NebularisParserContinue)
		}
		p.SetState(209)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NebularisParserIdentifier {
			{
				p.SetState(208)

				var _m = p.Match(NebularisParserIdentifier)

				localctx.(*ContinueStatementContext).label = _m
			}

		}
		{
			p.SetState(211)
			p.Match(NebularisParserSemiColon)
		}

	case 7:
		localctx = NewBreakStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(212)
			p.Match(NebularisParserBreak)
		}
		p.SetState(214)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NebularisParserIdentifier {
			{
				p.SetState(213)

				var _m = p.Match(NebularisParserIdentifier)

				localctx.(*BreakStatementContext).label = _m
			}

		}
		{
			p.SetState(216)
			p.Match(NebularisParserSemiColon)
		}

	case 8:
		localctx = NewSwitchStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(217)
			p.Match(NebularisParserSwitch)
		}
		{
			p.SetState(218)

			var _x = p.expression(0)

			localctx.(*SwitchStatementContext).expr = _x
		}
		{
			p.SetState(219)
			p.Match(NebularisParserCurlyOpen)
		}
		p.SetState(224)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NebularisParserCase || _la == NebularisParserDefault {
			p.SetState(222)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case NebularisParserCase:
				{
					p.SetState(220)

					var _x = p.CaseStatement()

					localctx.(*SwitchStatementContext)._caseStatement = _x
				}
				localctx.(*SwitchStatementContext).cases = append(localctx.(*SwitchStatementContext).cases, localctx.(*SwitchStatementContext)._caseStatement)

			case NebularisParserDefault:
				{
					p.SetState(221)

					var _x = p.DefaultStatement()

					localctx.(*SwitchStatementContext)._defaultStatement = _x
				}
				localctx.(*SwitchStatementContext).defauls = append(localctx.(*SwitchStatementContext).defauls, localctx.(*SwitchStatementContext)._defaultStatement)

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(226)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(227)
			p.Match(NebularisParserCurlyClose)
		}

	case 9:
		localctx = NewUnaryStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(229)
			p.expression(0)
		}
		{
			p.SetState(230)

			var _x = p.PostfixOp()

			localctx.(*UnaryStatementContext).op = _x
		}
		{
			p.SetState(231)
			p.Match(NebularisParserSemiColon)
		}

	case 10:
		localctx = NewAssignmentStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(233)

			var _x = p.expression(0)

			localctx.(*AssignmentStatementContext).left = _x
		}
		{
			p.SetState(234)
			p.Match(NebularisParserEquals)
		}
		{
			p.SetState(235)

			var _x = p.expression(0)

			localctx.(*AssignmentStatementContext).right = _x
		}
		{
			p.SetState(236)
			p.Match(NebularisParserSemiColon)
		}

	case 11:
		localctx = NewNotYetImplementedStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(238)
			p.Match(NebularisParserEllipsis)
		}

	case 12:
		localctx = NewExpressionStatementContext(p, localctx)
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(239)
			p.expression(0)
		}
		{
			p.SetState(240)
			p.Match(NebularisParserSemiColon)
		}

	}

	return localctx
}

// ICaseStatementContext is an interface to support dynamic dispatch.
type ICaseStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// GetBlock returns the block rule contexts.
	GetBlock() ICodeBlockContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// SetBlock sets the block rule contexts.
	SetBlock(ICodeBlockContext)

	// GetExpr returns the expr rule context list.
	GetExpr() []IExpressionContext

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetExpr sets the expr rule context list.
	SetExpr([]IExpressionContext)

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// IsCaseStatementContext differentiates from other interfaces.
	IsCaseStatementContext()
}

type CaseStatementContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_expression IExpressionContext
	expr        []IExpressionContext
	_statement  IStatementContext
	statements  []IStatementContext
	block       ICodeBlockContext
}

func NewEmptyCaseStatementContext() *CaseStatementContext {
	var p = new(CaseStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_caseStatement
	return p
}

func (*CaseStatementContext) IsCaseStatementContext() {}

func NewCaseStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CaseStatementContext {
	var p = new(CaseStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_caseStatement

	return p
}

func (s *CaseStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *CaseStatementContext) Get_expression() IExpressionContext { return s._expression }

func (s *CaseStatementContext) Get_statement() IStatementContext { return s._statement }

func (s *CaseStatementContext) GetBlock() ICodeBlockContext { return s.block }

func (s *CaseStatementContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *CaseStatementContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *CaseStatementContext) SetBlock(v ICodeBlockContext) { s.block = v }

func (s *CaseStatementContext) GetExpr() []IExpressionContext { return s.expr }

func (s *CaseStatementContext) GetStatements() []IStatementContext { return s.statements }

func (s *CaseStatementContext) SetExpr(v []IExpressionContext) { s.expr = v }

func (s *CaseStatementContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *CaseStatementContext) Case() antlr.TerminalNode {
	return s.GetToken(NebularisParserCase, 0)
}

func (s *CaseStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *CaseStatementContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CaseStatementContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CaseStatementContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *CaseStatementContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *CaseStatementContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *CaseStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *CaseStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *CaseStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CaseStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CaseStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterCaseStatement(s)
	}
}

func (s *CaseStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitCaseStatement(s)
	}
}

func (s *CaseStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitCaseStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) CaseStatement() (localctx ICaseStatementContext) {
	localctx = NewCaseStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, NebularisParserRULE_caseStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(244)
		p.Match(NebularisParserCase)
	}
	{
		p.SetState(245)

		var _x = p.expression(0)

		localctx.(*CaseStatementContext)._expression = _x
	}
	localctx.(*CaseStatementContext).expr = append(localctx.(*CaseStatementContext).expr, localctx.(*CaseStatementContext)._expression)
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserComma {
		{
			p.SetState(246)
			p.Match(NebularisParserComma)
		}
		{
			p.SetState(247)

			var _x = p.expression(0)

			localctx.(*CaseStatementContext)._expression = _x
		}
		localctx.(*CaseStatementContext).expr = append(localctx.(*CaseStatementContext).expr, localctx.(*CaseStatementContext)._expression)

		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(253)
		p.Match(NebularisParserColon)
	}
	p.SetState(260)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NebularisParserIntegerLiteral, NebularisParserBoolLiteral, NebularisParserStringLiteral, NebularisParserBreak, NebularisParserContinue, NebularisParserFn, NebularisParserIf, NebularisParserReturn, NebularisParserSwitch, NebularisParserValue, NebularisParserVar, NebularisParserWhile, NebularisParserEllipsis, NebularisParserNot, NebularisParserParenOpen, NebularisParserTilda, NebularisParserIdentifier:
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(NebularisParserIntegerLiteral-1))|(1<<(NebularisParserBoolLiteral-1))|(1<<(NebularisParserStringLiteral-1))|(1<<(NebularisParserBreak-1))|(1<<(NebularisParserContinue-1))|(1<<(NebularisParserFn-1))|(1<<(NebularisParserIf-1))|(1<<(NebularisParserReturn-1))|(1<<(NebularisParserSwitch-1))|(1<<(NebularisParserValue-1))|(1<<(NebularisParserVar-1))|(1<<(NebularisParserWhile-1)))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(NebularisParserEllipsis-42))|(1<<(NebularisParserNot-42))|(1<<(NebularisParserParenOpen-42))|(1<<(NebularisParserTilda-42))|(1<<(NebularisParserIdentifier-42)))) != 0) {
			{
				p.SetState(254)

				var _x = p.Statement()

				localctx.(*CaseStatementContext)._statement = _x
			}
			localctx.(*CaseStatementContext).statements = append(localctx.(*CaseStatementContext).statements, localctx.(*CaseStatementContext)._statement)

			p.SetState(257)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case NebularisParserCurlyOpen:
		{
			p.SetState(259)

			var _x = p.CodeBlock()

			localctx.(*CaseStatementContext).block = _x
		}

	case NebularisParserCase, NebularisParserDefault, NebularisParserCurlyClose:

	default:
	}

	return localctx
}

// IDefaultStatementContext is an interface to support dynamic dispatch.
type IDefaultStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_statement returns the _statement rule contexts.
	Get_statement() IStatementContext

	// GetBlock returns the block rule contexts.
	GetBlock() ICodeBlockContext

	// Set_statement sets the _statement rule contexts.
	Set_statement(IStatementContext)

	// SetBlock sets the block rule contexts.
	SetBlock(ICodeBlockContext)

	// GetStatements returns the statements rule context list.
	GetStatements() []IStatementContext

	// SetStatements sets the statements rule context list.
	SetStatements([]IStatementContext)

	// IsDefaultStatementContext differentiates from other interfaces.
	IsDefaultStatementContext()
}

type DefaultStatementContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	_statement IStatementContext
	statements []IStatementContext
	block      ICodeBlockContext
}

func NewEmptyDefaultStatementContext() *DefaultStatementContext {
	var p = new(DefaultStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_defaultStatement
	return p
}

func (*DefaultStatementContext) IsDefaultStatementContext() {}

func NewDefaultStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultStatementContext {
	var p = new(DefaultStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_defaultStatement

	return p
}

func (s *DefaultStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultStatementContext) Get_statement() IStatementContext { return s._statement }

func (s *DefaultStatementContext) GetBlock() ICodeBlockContext { return s.block }

func (s *DefaultStatementContext) Set_statement(v IStatementContext) { s._statement = v }

func (s *DefaultStatementContext) SetBlock(v ICodeBlockContext) { s.block = v }

func (s *DefaultStatementContext) GetStatements() []IStatementContext { return s.statements }

func (s *DefaultStatementContext) SetStatements(v []IStatementContext) { s.statements = v }

func (s *DefaultStatementContext) Default() antlr.TerminalNode {
	return s.GetToken(NebularisParserDefault, 0)
}

func (s *DefaultStatementContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *DefaultStatementContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *DefaultStatementContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *DefaultStatementContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *DefaultStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterDefaultStatement(s)
	}
}

func (s *DefaultStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitDefaultStatement(s)
	}
}

func (s *DefaultStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitDefaultStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) DefaultStatement() (localctx IDefaultStatementContext) {
	localctx = NewDefaultStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, NebularisParserRULE_defaultStatement)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(262)
		p.Match(NebularisParserDefault)
	}
	{
		p.SetState(263)
		p.Match(NebularisParserColon)
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NebularisParserIntegerLiteral, NebularisParserBoolLiteral, NebularisParserStringLiteral, NebularisParserBreak, NebularisParserContinue, NebularisParserFn, NebularisParserIf, NebularisParserReturn, NebularisParserSwitch, NebularisParserValue, NebularisParserVar, NebularisParserWhile, NebularisParserEllipsis, NebularisParserNot, NebularisParserParenOpen, NebularisParserTilda, NebularisParserIdentifier:
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = (((_la-1)&-(0x1f+1)) == 0 && ((1<<uint((_la-1)))&((1<<(NebularisParserIntegerLiteral-1))|(1<<(NebularisParserBoolLiteral-1))|(1<<(NebularisParserStringLiteral-1))|(1<<(NebularisParserBreak-1))|(1<<(NebularisParserContinue-1))|(1<<(NebularisParserFn-1))|(1<<(NebularisParserIf-1))|(1<<(NebularisParserReturn-1))|(1<<(NebularisParserSwitch-1))|(1<<(NebularisParserValue-1))|(1<<(NebularisParserVar-1))|(1<<(NebularisParserWhile-1)))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(NebularisParserEllipsis-42))|(1<<(NebularisParserNot-42))|(1<<(NebularisParserParenOpen-42))|(1<<(NebularisParserTilda-42))|(1<<(NebularisParserIdentifier-42)))) != 0) {
			{
				p.SetState(264)

				var _x = p.Statement()

				localctx.(*DefaultStatementContext)._statement = _x
			}
			localctx.(*DefaultStatementContext).statements = append(localctx.(*DefaultStatementContext).statements, localctx.(*DefaultStatementContext)._statement)

			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case NebularisParserCurlyOpen:
		{
			p.SetState(269)

			var _x = p.CodeBlock()

			localctx.(*DefaultStatementContext).block = _x
		}

	case NebularisParserCase, NebularisParserDefault, NebularisParserCurlyClose:

	default:
	}

	return localctx
}

// ITypeDeclarationContext is an interface to support dynamic dispatch.
type ITypeDeclarationContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_attribute returns the _attribute rule contexts.
	Get_attribute() IAttributeContext

	// GetTypeParams returns the typeParams rule contexts.
	GetTypeParams() ITypeParametersContext

	// GetSpec returns the spec rule contexts.
	GetSpec() ITypeSpecContext

	// Get_whereClause returns the _whereClause rule contexts.
	Get_whereClause() IWhereClauseContext

	// Set_attribute sets the _attribute rule contexts.
	Set_attribute(IAttributeContext)

	// SetTypeParams sets the typeParams rule contexts.
	SetTypeParams(ITypeParametersContext)

	// SetSpec sets the spec rule contexts.
	SetSpec(ITypeSpecContext)

	// Set_whereClause sets the _whereClause rule contexts.
	Set_whereClause(IWhereClauseContext)

	// GetAttributes returns the attributes rule context list.
	GetAttributes() []IAttributeContext

	// GetConstraints returns the constraints rule context list.
	GetConstraints() []IWhereClauseContext

	// SetAttributes sets the attributes rule context list.
	SetAttributes([]IAttributeContext)

	// SetConstraints sets the constraints rule context list.
	SetConstraints([]IWhereClauseContext)

	// IsTypeDeclarationContext differentiates from other interfaces.
	IsTypeDeclarationContext()
}

type TypeDeclarationContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_attribute   IAttributeContext
	attributes   []IAttributeContext
	name         antlr.Token
	typeParams   ITypeParametersContext
	spec         ITypeSpecContext
	_whereClause IWhereClauseContext
	constraints  []IWhereClauseContext
}

func NewEmptyTypeDeclarationContext() *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_typeDeclaration
	return p
}

func (*TypeDeclarationContext) IsTypeDeclarationContext() {}

func NewTypeDeclarationContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeDeclarationContext {
	var p = new(TypeDeclarationContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_typeDeclaration

	return p
}

func (s *TypeDeclarationContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeDeclarationContext) GetName() antlr.Token { return s.name }

func (s *TypeDeclarationContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeDeclarationContext) Get_attribute() IAttributeContext { return s._attribute }

func (s *TypeDeclarationContext) GetTypeParams() ITypeParametersContext { return s.typeParams }

func (s *TypeDeclarationContext) GetSpec() ITypeSpecContext { return s.spec }

func (s *TypeDeclarationContext) Get_whereClause() IWhereClauseContext { return s._whereClause }

func (s *TypeDeclarationContext) Set_attribute(v IAttributeContext) { s._attribute = v }

func (s *TypeDeclarationContext) SetTypeParams(v ITypeParametersContext) { s.typeParams = v }

func (s *TypeDeclarationContext) SetSpec(v ITypeSpecContext) { s.spec = v }

func (s *TypeDeclarationContext) Set_whereClause(v IWhereClauseContext) { s._whereClause = v }

func (s *TypeDeclarationContext) GetAttributes() []IAttributeContext { return s.attributes }

func (s *TypeDeclarationContext) GetConstraints() []IWhereClauseContext { return s.constraints }

func (s *TypeDeclarationContext) SetAttributes(v []IAttributeContext) { s.attributes = v }

func (s *TypeDeclarationContext) SetConstraints(v []IWhereClauseContext) { s.constraints = v }

func (s *TypeDeclarationContext) Type() antlr.TerminalNode {
	return s.GetToken(NebularisParserType, 0)
}

func (s *TypeDeclarationContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *TypeDeclarationContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *TypeDeclarationContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *TypeDeclarationContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *TypeDeclarationContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *TypeDeclarationContext) TypeParameters() ITypeParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *TypeDeclarationContext) AllWhereClause() []IWhereClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem())
	var tst = make([]IWhereClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereClauseContext)
		}
	}

	return tst
}

func (s *TypeDeclarationContext) WhereClause(i int) IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *TypeDeclarationContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeDeclarationContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeDeclarationContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitTypeDeclaration(s)
	}
}

func (s *TypeDeclarationContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitTypeDeclaration(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) TypeDeclaration() (localctx ITypeDeclarationContext) {
	localctx = NewTypeDeclarationContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, NebularisParserRULE_typeDeclaration)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserSquareOpen {
		{
			p.SetState(272)

			var _x = p.Attribute()

			localctx.(*TypeDeclarationContext)._attribute = _x
		}
		localctx.(*TypeDeclarationContext).attributes = append(localctx.(*TypeDeclarationContext).attributes, localctx.(*TypeDeclarationContext)._attribute)

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(278)
		p.Match(NebularisParserType)
	}
	{
		p.SetState(279)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*TypeDeclarationContext).name = _m
	}
	p.SetState(281)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserLessThan {
		{
			p.SetState(280)

			var _x = p.TypeParameters()

			localctx.(*TypeDeclarationContext).typeParams = _x
		}

	}
	{
		p.SetState(283)

		var _x = p.TypeSpec()

		localctx.(*TypeDeclarationContext).spec = _x
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserWhere {
		{
			p.SetState(284)

			var _x = p.WhereClause()

			localctx.(*TypeDeclarationContext)._whereClause = _x
		}
		localctx.(*TypeDeclarationContext).constraints = append(localctx.(*TypeDeclarationContext).constraints, localctx.(*TypeDeclarationContext)._whereClause)

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(291)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserSemiColon {
		{
			p.SetState(290)
			p.Match(NebularisParserSemiColon)
		}

	}

	return localctx
}

// ITypeSpecContext is an interface to support dynamic dispatch.
type ITypeSpecContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTypeSpecContext differentiates from other interfaces.
	IsTypeSpecContext()
}

type TypeSpecContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeSpecContext() *TypeSpecContext {
	var p = new(TypeSpecContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_typeSpec
	return p
}

func (*TypeSpecContext) IsTypeSpecContext() {}

func NewTypeSpecContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeSpecContext {
	var p = new(TypeSpecContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_typeSpec

	return p
}

func (s *TypeSpecContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeSpecContext) CopyFrom(ctx *TypeSpecContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *TypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeSpecContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StructTypeSpecContext struct {
	*TypeSpecContext
	_field IFieldContext
	fields []IFieldContext
}

func NewStructTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructTypeSpecContext {
	var p = new(StructTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *StructTypeSpecContext) Get_field() IFieldContext { return s._field }

func (s *StructTypeSpecContext) Set_field(v IFieldContext) { s._field = v }

func (s *StructTypeSpecContext) GetFields() []IFieldContext { return s.fields }

func (s *StructTypeSpecContext) SetFields(v []IFieldContext) { s.fields = v }

func (s *StructTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructTypeSpecContext) Struct() antlr.TerminalNode {
	return s.GetToken(NebularisParserStruct, 0)
}

func (s *StructTypeSpecContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyOpen, 0)
}

func (s *StructTypeSpecContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyClose, 0)
}

func (s *StructTypeSpecContext) AllField() []IFieldContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldContext)(nil)).Elem())
	var tst = make([]IFieldContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldContext)
		}
	}

	return tst
}

func (s *StructTypeSpecContext) Field(i int) IFieldContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldContext)
}

func (s *StructTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitStructTypeSpec(s)
	}
}

func (s *StructTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitStructTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

type SpanTypeSpecContext struct {
	*TypeSpecContext
	element ITypeSpecContext
}

func NewSpanTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SpanTypeSpecContext {
	var p = new(SpanTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *SpanTypeSpecContext) GetElement() ITypeSpecContext { return s.element }

func (s *SpanTypeSpecContext) SetElement(v ITypeSpecContext) { s.element = v }

func (s *SpanTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpanTypeSpecContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserSquareOpen, 0)
}

func (s *SpanTypeSpecContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserSquareClose, 0)
}

func (s *SpanTypeSpecContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *SpanTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterSpanTypeSpec(s)
	}
}

func (s *SpanTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitSpanTypeSpec(s)
	}
}

func (s *SpanTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitSpanTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

type NullableTypeSpecContext struct {
	*TypeSpecContext
	element ITypeSpecContext
}

func NewNullableTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NullableTypeSpecContext {
	var p = new(NullableTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *NullableTypeSpecContext) GetElement() ITypeSpecContext { return s.element }

func (s *NullableTypeSpecContext) SetElement(v ITypeSpecContext) { s.element = v }

func (s *NullableTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NullableTypeSpecContext) Question() antlr.TerminalNode {
	return s.GetToken(NebularisParserQuestion, 0)
}

func (s *NullableTypeSpecContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *NullableTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterNullableTypeSpec(s)
	}
}

func (s *NullableTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitNullableTypeSpec(s)
	}
}

func (s *NullableTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitNullableTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionTypeSpecContext struct {
	*TypeSpecContext
	params     IFunctionParametersContext
	returnType ITypeSpecContext
}

func NewFunctionTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionTypeSpecContext {
	var p = new(FunctionTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *FunctionTypeSpecContext) GetParams() IFunctionParametersContext { return s.params }

func (s *FunctionTypeSpecContext) GetReturnType() ITypeSpecContext { return s.returnType }

func (s *FunctionTypeSpecContext) SetParams(v IFunctionParametersContext) { s.params = v }

func (s *FunctionTypeSpecContext) SetReturnType(v ITypeSpecContext) { s.returnType = v }

func (s *FunctionTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionTypeSpecContext) Fn() antlr.TerminalNode {
	return s.GetToken(NebularisParserFn, 0)
}

func (s *FunctionTypeSpecContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *FunctionTypeSpecContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FunctionTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterFunctionTypeSpec(s)
	}
}

func (s *FunctionTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitFunctionTypeSpec(s)
	}
}

func (s *FunctionTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitFunctionTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

type ReferencedTypeSpecContext struct {
	*TypeSpecContext
	refType IReferencedTypeContext
}

func NewReferencedTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ReferencedTypeSpecContext {
	var p = new(ReferencedTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *ReferencedTypeSpecContext) GetRefType() IReferencedTypeContext { return s.refType }

func (s *ReferencedTypeSpecContext) SetRefType(v IReferencedTypeContext) { s.refType = v }

func (s *ReferencedTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferencedTypeSpecContext) ReferencedType() IReferencedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferencedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferencedTypeContext)
}

func (s *ReferencedTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterReferencedTypeSpec(s)
	}
}

func (s *ReferencedTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitReferencedTypeSpec(s)
	}
}

func (s *ReferencedTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitReferencedTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrimitiveTypeSpecContext struct {
	*TypeSpecContext
	primitive IPrimitiveTypeContext
}

func NewPrimitiveTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrimitiveTypeSpecContext {
	var p = new(PrimitiveTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *PrimitiveTypeSpecContext) GetPrimitive() IPrimitiveTypeContext { return s.primitive }

func (s *PrimitiveTypeSpecContext) SetPrimitive(v IPrimitiveTypeContext) { s.primitive = v }

func (s *PrimitiveTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeSpecContext) PrimitiveType() IPrimitiveTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitiveTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitiveTypeContext)
}

func (s *PrimitiveTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterPrimitiveTypeSpec(s)
	}
}

func (s *PrimitiveTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitPrimitiveTypeSpec(s)
	}
}

func (s *PrimitiveTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitPrimitiveTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

type InterfaceTypeSpecContext struct {
	*TypeSpecContext
	extends IExtendsClauseContext
	_method IMethodContext
	methods []IMethodContext
}

func NewInterfaceTypeSpecContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InterfaceTypeSpecContext {
	var p = new(InterfaceTypeSpecContext)

	p.TypeSpecContext = NewEmptyTypeSpecContext()
	p.parser = parser
	p.CopyFrom(ctx.(*TypeSpecContext))

	return p
}

func (s *InterfaceTypeSpecContext) GetExtends() IExtendsClauseContext { return s.extends }

func (s *InterfaceTypeSpecContext) Get_method() IMethodContext { return s._method }

func (s *InterfaceTypeSpecContext) SetExtends(v IExtendsClauseContext) { s.extends = v }

func (s *InterfaceTypeSpecContext) Set_method(v IMethodContext) { s._method = v }

func (s *InterfaceTypeSpecContext) GetMethods() []IMethodContext { return s.methods }

func (s *InterfaceTypeSpecContext) SetMethods(v []IMethodContext) { s.methods = v }

func (s *InterfaceTypeSpecContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InterfaceTypeSpecContext) Interface() antlr.TerminalNode {
	return s.GetToken(NebularisParserInterface, 0)
}

func (s *InterfaceTypeSpecContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyOpen, 0)
}

func (s *InterfaceTypeSpecContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyClose, 0)
}

func (s *InterfaceTypeSpecContext) ExtendsClause() IExtendsClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExtendsClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExtendsClauseContext)
}

func (s *InterfaceTypeSpecContext) AllMethod() []IMethodContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IMethodContext)(nil)).Elem())
	var tst = make([]IMethodContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IMethodContext)
		}
	}

	return tst
}

func (s *InterfaceTypeSpecContext) Method(i int) IMethodContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMethodContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IMethodContext)
}

func (s *InterfaceTypeSpecContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterInterfaceTypeSpec(s)
	}
}

func (s *InterfaceTypeSpecContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitInterfaceTypeSpec(s)
	}
}

func (s *InterfaceTypeSpecContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitInterfaceTypeSpec(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) TypeSpec() (localctx ITypeSpecContext) {
	localctx = NewTypeSpecContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, NebularisParserRULE_typeSpec)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(326)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NebularisParserQuestion:
		localctx = NewNullableTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(293)
			p.Match(NebularisParserQuestion)
		}
		{
			p.SetState(294)

			var _x = p.TypeSpec()

			localctx.(*NullableTypeSpecContext).element = _x
		}

	case NebularisParserSquareOpen:
		localctx = NewSpanTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(295)
			p.Match(NebularisParserSquareOpen)
		}
		{
			p.SetState(296)
			p.Match(NebularisParserSquareClose)
		}
		{
			p.SetState(297)

			var _x = p.TypeSpec()

			localctx.(*SpanTypeSpecContext).element = _x
		}

	case NebularisParserIdentifier:
		localctx = NewReferencedTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(298)

			var _x = p.ReferencedType()

			localctx.(*ReferencedTypeSpecContext).refType = _x
		}

	case NebularisParserBool, NebularisParserByte, NebularisParserChar, NebularisParserInt32, NebularisParserStr:
		localctx = NewPrimitiveTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(299)

			var _x = p.PrimitiveType()

			localctx.(*PrimitiveTypeSpecContext).primitive = _x
		}

	case NebularisParserFn:
		localctx = NewFunctionTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(300)
			p.Match(NebularisParserFn)
		}
		{
			p.SetState(301)

			var _x = p.FunctionParameters()

			localctx.(*FunctionTypeSpecContext).params = _x
		}
		p.SetState(303)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(302)

				var _x = p.TypeSpec()

				localctx.(*FunctionTypeSpecContext).returnType = _x
			}

		}

	case NebularisParserStruct:
		localctx = NewStructTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(305)
			p.Match(NebularisParserStruct)
		}
		{
			p.SetState(306)
			p.Match(NebularisParserCurlyOpen)
		}
		p.SetState(310)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NebularisParserSquareOpen || _la == NebularisParserIdentifier {
			{
				p.SetState(307)

				var _x = p.Field()

				localctx.(*StructTypeSpecContext)._field = _x
			}
			localctx.(*StructTypeSpecContext).fields = append(localctx.(*StructTypeSpecContext).fields, localctx.(*StructTypeSpecContext)._field)

			p.SetState(312)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(313)
			p.Match(NebularisParserCurlyClose)
		}

	case NebularisParserInterface:
		localctx = NewInterfaceTypeSpecContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(314)
			p.Match(NebularisParserInterface)
		}
		p.SetState(316)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == NebularisParserColon {
			{
				p.SetState(315)

				var _x = p.ExtendsClause()

				localctx.(*InterfaceTypeSpecContext).extends = _x
			}

		}
		{
			p.SetState(318)
			p.Match(NebularisParserCurlyOpen)
		}
		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NebularisParserSquareOpen || _la == NebularisParserIdentifier {
			{
				p.SetState(319)

				var _x = p.Method()

				localctx.(*InterfaceTypeSpecContext)._method = _x
			}
			localctx.(*InterfaceTypeSpecContext).methods = append(localctx.(*InterfaceTypeSpecContext).methods, localctx.(*InterfaceTypeSpecContext)._method)

			p.SetState(324)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(325)
			p.Match(NebularisParserCurlyClose)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IReferencedTypeContext is an interface to support dynamic dispatch.
type IReferencedTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name rule contexts.
	GetName() IQualifiedNameContext

	// Get_typeSpec returns the _typeSpec rule contexts.
	Get_typeSpec() ITypeSpecContext

	// SetName sets the name rule contexts.
	SetName(IQualifiedNameContext)

	// Set_typeSpec sets the _typeSpec rule contexts.
	Set_typeSpec(ITypeSpecContext)

	// GetTypeArgs returns the typeArgs rule context list.
	GetTypeArgs() []ITypeSpecContext

	// SetTypeArgs sets the typeArgs rule context list.
	SetTypeArgs([]ITypeSpecContext)

	// IsReferencedTypeContext differentiates from other interfaces.
	IsReferencedTypeContext()
}

type ReferencedTypeContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	name      IQualifiedNameContext
	_typeSpec ITypeSpecContext
	typeArgs  []ITypeSpecContext
}

func NewEmptyReferencedTypeContext() *ReferencedTypeContext {
	var p = new(ReferencedTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_referencedType
	return p
}

func (*ReferencedTypeContext) IsReferencedTypeContext() {}

func NewReferencedTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ReferencedTypeContext {
	var p = new(ReferencedTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_referencedType

	return p
}

func (s *ReferencedTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *ReferencedTypeContext) GetName() IQualifiedNameContext { return s.name }

func (s *ReferencedTypeContext) Get_typeSpec() ITypeSpecContext { return s._typeSpec }

func (s *ReferencedTypeContext) SetName(v IQualifiedNameContext) { s.name = v }

func (s *ReferencedTypeContext) Set_typeSpec(v ITypeSpecContext) { s._typeSpec = v }

func (s *ReferencedTypeContext) GetTypeArgs() []ITypeSpecContext { return s.typeArgs }

func (s *ReferencedTypeContext) SetTypeArgs(v []ITypeSpecContext) { s.typeArgs = v }

func (s *ReferencedTypeContext) QualifiedName() IQualifiedNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IQualifiedNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IQualifiedNameContext)
}

func (s *ReferencedTypeContext) LessThan() antlr.TerminalNode {
	return s.GetToken(NebularisParserLessThan, 0)
}

func (s *ReferencedTypeContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(NebularisParserGreaterThan, 0)
}

func (s *ReferencedTypeContext) AllTypeSpec() []ITypeSpecContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem())
	var tst = make([]ITypeSpecContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeSpecContext)
		}
	}

	return tst
}

func (s *ReferencedTypeContext) TypeSpec(i int) ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *ReferencedTypeContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *ReferencedTypeContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *ReferencedTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ReferencedTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ReferencedTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterReferencedType(s)
	}
}

func (s *ReferencedTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitReferencedType(s)
	}
}

func (s *ReferencedTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitReferencedType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) ReferencedType() (localctx IReferencedTypeContext) {
	localctx = NewReferencedTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, NebularisParserRULE_referencedType)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(328)

		var _x = p.QualifiedName()

		localctx.(*ReferencedTypeContext).name = _x
	}
	p.SetState(340)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserLessThan {
		{
			p.SetState(329)
			p.Match(NebularisParserLessThan)
		}
		{
			p.SetState(330)

			var _x = p.TypeSpec()

			localctx.(*ReferencedTypeContext)._typeSpec = _x
		}
		localctx.(*ReferencedTypeContext).typeArgs = append(localctx.(*ReferencedTypeContext).typeArgs, localctx.(*ReferencedTypeContext)._typeSpec)
		p.SetState(335)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NebularisParserComma {
			{
				p.SetState(331)
				p.Match(NebularisParserComma)
			}
			{
				p.SetState(332)

				var _x = p.TypeSpec()

				localctx.(*ReferencedTypeContext)._typeSpec = _x
			}
			localctx.(*ReferencedTypeContext).typeArgs = append(localctx.(*ReferencedTypeContext).typeArgs, localctx.(*ReferencedTypeContext)._typeSpec)

			p.SetState(337)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(338)
			p.Match(NebularisParserGreaterThan)
		}

	}

	return localctx
}

// IExtendsClauseContext is an interface to support dynamic dispatch.
type IExtendsClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_referencedType returns the _referencedType rule contexts.
	Get_referencedType() IReferencedTypeContext

	// Set_referencedType sets the _referencedType rule contexts.
	Set_referencedType(IReferencedTypeContext)

	// GetRefTypes returns the refTypes rule context list.
	GetRefTypes() []IReferencedTypeContext

	// SetRefTypes sets the refTypes rule context list.
	SetRefTypes([]IReferencedTypeContext)

	// IsExtendsClauseContext differentiates from other interfaces.
	IsExtendsClauseContext()
}

type ExtendsClauseContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	_referencedType IReferencedTypeContext
	refTypes        []IReferencedTypeContext
}

func NewEmptyExtendsClauseContext() *ExtendsClauseContext {
	var p = new(ExtendsClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_extendsClause
	return p
}

func (*ExtendsClauseContext) IsExtendsClauseContext() {}

func NewExtendsClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExtendsClauseContext {
	var p = new(ExtendsClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_extendsClause

	return p
}

func (s *ExtendsClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *ExtendsClauseContext) Get_referencedType() IReferencedTypeContext { return s._referencedType }

func (s *ExtendsClauseContext) Set_referencedType(v IReferencedTypeContext) { s._referencedType = v }

func (s *ExtendsClauseContext) GetRefTypes() []IReferencedTypeContext { return s.refTypes }

func (s *ExtendsClauseContext) SetRefTypes(v []IReferencedTypeContext) { s.refTypes = v }

func (s *ExtendsClauseContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *ExtendsClauseContext) AllReferencedType() []IReferencedTypeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IReferencedTypeContext)(nil)).Elem())
	var tst = make([]IReferencedTypeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IReferencedTypeContext)
		}
	}

	return tst
}

func (s *ExtendsClauseContext) ReferencedType(i int) IReferencedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferencedTypeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IReferencedTypeContext)
}

func (s *ExtendsClauseContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *ExtendsClauseContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *ExtendsClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExtendsClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExtendsClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterExtendsClause(s)
	}
}

func (s *ExtendsClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitExtendsClause(s)
	}
}

func (s *ExtendsClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitExtendsClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) ExtendsClause() (localctx IExtendsClauseContext) {
	localctx = NewExtendsClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, NebularisParserRULE_extendsClause)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(342)
		p.Match(NebularisParserColon)
	}
	p.SetState(351)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserIdentifier {
		{
			p.SetState(343)

			var _x = p.ReferencedType()

			localctx.(*ExtendsClauseContext)._referencedType = _x
		}
		localctx.(*ExtendsClauseContext).refTypes = append(localctx.(*ExtendsClauseContext).refTypes, localctx.(*ExtendsClauseContext)._referencedType)
		p.SetState(348)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == NebularisParserComma {
			{
				p.SetState(344)
				p.Match(NebularisParserComma)
			}
			{
				p.SetState(345)

				var _x = p.ReferencedType()

				localctx.(*ExtendsClauseContext)._referencedType = _x
			}
			localctx.(*ExtendsClauseContext).refTypes = append(localctx.(*ExtendsClauseContext).refTypes, localctx.(*ExtendsClauseContext)._referencedType)

			p.SetState(350)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	}

	return localctx
}

// IMethodContext is an interface to support dynamic dispatch.
type IMethodContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_attribute returns the _attribute rule contexts.
	Get_attribute() IAttributeContext

	// GetParams returns the params rule contexts.
	GetParams() IFunctionParametersContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeSpecContext

	// Get_whereClause returns the _whereClause rule contexts.
	Get_whereClause() IWhereClauseContext

	// Set_attribute sets the _attribute rule contexts.
	Set_attribute(IAttributeContext)

	// SetParams sets the params rule contexts.
	SetParams(IFunctionParametersContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeSpecContext)

	// Set_whereClause sets the _whereClause rule contexts.
	Set_whereClause(IWhereClauseContext)

	// GetAttributes returns the attributes rule context list.
	GetAttributes() []IAttributeContext

	// GetConstraints returns the constraints rule context list.
	GetConstraints() []IWhereClauseContext

	// SetAttributes sets the attributes rule context list.
	SetAttributes([]IAttributeContext)

	// SetConstraints sets the constraints rule context list.
	SetConstraints([]IWhereClauseContext)

	// IsMethodContext differentiates from other interfaces.
	IsMethodContext()
}

type MethodContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_attribute   IAttributeContext
	attributes   []IAttributeContext
	name         antlr.Token
	params       IFunctionParametersContext
	returnType   ITypeSpecContext
	_whereClause IWhereClauseContext
	constraints  []IWhereClauseContext
}

func NewEmptyMethodContext() *MethodContext {
	var p = new(MethodContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_method
	return p
}

func (*MethodContext) IsMethodContext() {}

func NewMethodContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MethodContext {
	var p = new(MethodContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_method

	return p
}

func (s *MethodContext) GetParser() antlr.Parser { return s.parser }

func (s *MethodContext) GetName() antlr.Token { return s.name }

func (s *MethodContext) SetName(v antlr.Token) { s.name = v }

func (s *MethodContext) Get_attribute() IAttributeContext { return s._attribute }

func (s *MethodContext) GetParams() IFunctionParametersContext { return s.params }

func (s *MethodContext) GetReturnType() ITypeSpecContext { return s.returnType }

func (s *MethodContext) Get_whereClause() IWhereClauseContext { return s._whereClause }

func (s *MethodContext) Set_attribute(v IAttributeContext) { s._attribute = v }

func (s *MethodContext) SetParams(v IFunctionParametersContext) { s.params = v }

func (s *MethodContext) SetReturnType(v ITypeSpecContext) { s.returnType = v }

func (s *MethodContext) Set_whereClause(v IWhereClauseContext) { s._whereClause = v }

func (s *MethodContext) GetAttributes() []IAttributeContext { return s.attributes }

func (s *MethodContext) GetConstraints() []IWhereClauseContext { return s.constraints }

func (s *MethodContext) SetAttributes(v []IAttributeContext) { s.attributes = v }

func (s *MethodContext) SetConstraints(v []IWhereClauseContext) { s.constraints = v }

func (s *MethodContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *MethodContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *MethodContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *MethodContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *MethodContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *MethodContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *MethodContext) AllWhereClause() []IWhereClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem())
	var tst = make([]IWhereClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereClauseContext)
		}
	}

	return tst
}

func (s *MethodContext) WhereClause(i int) IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *MethodContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MethodContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MethodContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterMethod(s)
	}
}

func (s *MethodContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitMethod(s)
	}
}

func (s *MethodContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitMethod(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Method() (localctx IMethodContext) {
	localctx = NewMethodContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, NebularisParserRULE_method)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(356)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserSquareOpen {
		{
			p.SetState(353)

			var _x = p.Attribute()

			localctx.(*MethodContext)._attribute = _x
		}
		localctx.(*MethodContext).attributes = append(localctx.(*MethodContext).attributes, localctx.(*MethodContext)._attribute)

		p.SetState(358)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(359)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*MethodContext).name = _m
	}
	{
		p.SetState(360)

		var _x = p.FunctionParameters()

		localctx.(*MethodContext).params = _x
	}
	p.SetState(362)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(361)

			var _x = p.TypeSpec()

			localctx.(*MethodContext).returnType = _x
		}

	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserWhere {
		{
			p.SetState(364)

			var _x = p.WhereClause()

			localctx.(*MethodContext)._whereClause = _x
		}
		localctx.(*MethodContext).constraints = append(localctx.(*MethodContext).constraints, localctx.(*MethodContext)._whereClause)

		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserSemiColon {
		{
			p.SetState(370)
			p.Match(NebularisParserSemiColon)
		}

	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExpr returns the expr rule contexts.
	GetExpr() IExpressionContext

	// SetExpr sets the expr rule contexts.
	SetExpr(IExpressionContext)

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	expr   IExpressionContext
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) GetExpr() IExpressionContext { return s.expr }

func (s *WhereClauseContext) SetExpr(v IExpressionContext) { s.expr = v }

func (s *WhereClauseContext) Where() antlr.TerminalNode {
	return s.GetToken(NebularisParserWhere, 0)
}

func (s *WhereClauseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, NebularisParserRULE_whereClause)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(373)
		p.Match(NebularisParserWhere)
	}
	{
		p.SetState(374)

		var _x = p.expression(0)

		localctx.(*WhereClauseContext).expr = _x
	}

	return localctx
}

// ITypeParametersContext is an interface to support dynamic dispatch.
type ITypeParametersContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_typeParameter returns the _typeParameter rule contexts.
	Get_typeParameter() ITypeParameterContext

	// Set_typeParameter sets the _typeParameter rule contexts.
	Set_typeParameter(ITypeParameterContext)

	// GetParams returns the params rule context list.
	GetParams() []ITypeParameterContext

	// SetParams sets the params rule context list.
	SetParams([]ITypeParameterContext)

	// IsTypeParametersContext differentiates from other interfaces.
	IsTypeParametersContext()
}

type TypeParametersContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	_typeParameter ITypeParameterContext
	params         []ITypeParameterContext
}

func NewEmptyTypeParametersContext() *TypeParametersContext {
	var p = new(TypeParametersContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_typeParameters
	return p
}

func (*TypeParametersContext) IsTypeParametersContext() {}

func NewTypeParametersContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParametersContext {
	var p = new(TypeParametersContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_typeParameters

	return p
}

func (s *TypeParametersContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParametersContext) Get_typeParameter() ITypeParameterContext { return s._typeParameter }

func (s *TypeParametersContext) Set_typeParameter(v ITypeParameterContext) { s._typeParameter = v }

func (s *TypeParametersContext) GetParams() []ITypeParameterContext { return s.params }

func (s *TypeParametersContext) SetParams(v []ITypeParameterContext) { s.params = v }

func (s *TypeParametersContext) LessThan() antlr.TerminalNode {
	return s.GetToken(NebularisParserLessThan, 0)
}

func (s *TypeParametersContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(NebularisParserGreaterThan, 0)
}

func (s *TypeParametersContext) AllTypeParameter() []ITypeParameterContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ITypeParameterContext)(nil)).Elem())
	var tst = make([]ITypeParameterContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ITypeParameterContext)
		}
	}

	return tst
}

func (s *TypeParametersContext) TypeParameter(i int) ITypeParameterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParameterContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ITypeParameterContext)
}

func (s *TypeParametersContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *TypeParametersContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *TypeParametersContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParametersContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParametersContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterTypeParameters(s)
	}
}

func (s *TypeParametersContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitTypeParameters(s)
	}
}

func (s *TypeParametersContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitTypeParameters(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) TypeParameters() (localctx ITypeParametersContext) {
	localctx = NewTypeParametersContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, NebularisParserRULE_typeParameters)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(376)
		p.Match(NebularisParserLessThan)
	}
	{
		p.SetState(377)

		var _x = p.TypeParameter()

		localctx.(*TypeParametersContext)._typeParameter = _x
	}
	localctx.(*TypeParametersContext).params = append(localctx.(*TypeParametersContext).params, localctx.(*TypeParametersContext)._typeParameter)
	p.SetState(382)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserComma {
		{
			p.SetState(378)
			p.Match(NebularisParserComma)
		}
		{
			p.SetState(379)

			var _x = p.TypeParameter()

			localctx.(*TypeParametersContext)._typeParameter = _x
		}
		localctx.(*TypeParametersContext).params = append(localctx.(*TypeParametersContext).params, localctx.(*TypeParametersContext)._typeParameter)

		p.SetState(384)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(385)
		p.Match(NebularisParserGreaterThan)
	}

	return localctx
}

// ITypeParameterContext is an interface to support dynamic dispatch.
type ITypeParameterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// IsTypeParameterContext differentiates from other interfaces.
	IsTypeParameterContext()
}

type TypeParameterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
}

func NewEmptyTypeParameterContext() *TypeParameterContext {
	var p = new(TypeParameterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_typeParameter
	return p
}

func (*TypeParameterContext) IsTypeParameterContext() {}

func NewTypeParameterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeParameterContext {
	var p = new(TypeParameterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_typeParameter

	return p
}

func (s *TypeParameterContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeParameterContext) GetName() antlr.Token { return s.name }

func (s *TypeParameterContext) SetName(v antlr.Token) { s.name = v }

func (s *TypeParameterContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *TypeParameterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeParameterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypeParameterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterTypeParameter(s)
	}
}

func (s *TypeParameterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitTypeParameter(s)
	}
}

func (s *TypeParameterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitTypeParameter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) TypeParameter() (localctx ITypeParameterContext) {
	localctx = NewTypeParameterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, NebularisParserRULE_typeParameter)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(387)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*TypeParameterContext).name = _m
	}

	return localctx
}

// IPrimitiveTypeContext is an interface to support dynamic dispatch.
type IPrimitiveTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTInt32 returns the tInt32 token.
	GetTInt32() antlr.Token

	// GetTBool returns the tBool token.
	GetTBool() antlr.Token

	// GetTStr returns the tStr token.
	GetTStr() antlr.Token

	// GetTByte returns the tByte token.
	GetTByte() antlr.Token

	// GetTChar returns the tChar token.
	GetTChar() antlr.Token

	// SetTInt32 sets the tInt32 token.
	SetTInt32(antlr.Token)

	// SetTBool sets the tBool token.
	SetTBool(antlr.Token)

	// SetTStr sets the tStr token.
	SetTStr(antlr.Token)

	// SetTByte sets the tByte token.
	SetTByte(antlr.Token)

	// SetTChar sets the tChar token.
	SetTChar(antlr.Token)

	// IsPrimitiveTypeContext differentiates from other interfaces.
	IsPrimitiveTypeContext()
}

type PrimitiveTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tInt32 antlr.Token
	tBool  antlr.Token
	tStr   antlr.Token
	tByte  antlr.Token
	tChar  antlr.Token
}

func NewEmptyPrimitiveTypeContext() *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_primitiveType
	return p
}

func (*PrimitiveTypeContext) IsPrimitiveTypeContext() {}

func NewPrimitiveTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitiveTypeContext {
	var p = new(PrimitiveTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_primitiveType

	return p
}

func (s *PrimitiveTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitiveTypeContext) GetTInt32() antlr.Token { return s.tInt32 }

func (s *PrimitiveTypeContext) GetTBool() antlr.Token { return s.tBool }

func (s *PrimitiveTypeContext) GetTStr() antlr.Token { return s.tStr }

func (s *PrimitiveTypeContext) GetTByte() antlr.Token { return s.tByte }

func (s *PrimitiveTypeContext) GetTChar() antlr.Token { return s.tChar }

func (s *PrimitiveTypeContext) SetTInt32(v antlr.Token) { s.tInt32 = v }

func (s *PrimitiveTypeContext) SetTBool(v antlr.Token) { s.tBool = v }

func (s *PrimitiveTypeContext) SetTStr(v antlr.Token) { s.tStr = v }

func (s *PrimitiveTypeContext) SetTByte(v antlr.Token) { s.tByte = v }

func (s *PrimitiveTypeContext) SetTChar(v antlr.Token) { s.tChar = v }

func (s *PrimitiveTypeContext) Int32() antlr.TerminalNode {
	return s.GetToken(NebularisParserInt32, 0)
}

func (s *PrimitiveTypeContext) Bool() antlr.TerminalNode {
	return s.GetToken(NebularisParserBool, 0)
}

func (s *PrimitiveTypeContext) Str() antlr.TerminalNode {
	return s.GetToken(NebularisParserStr, 0)
}

func (s *PrimitiveTypeContext) Byte() antlr.TerminalNode {
	return s.GetToken(NebularisParserByte, 0)
}

func (s *PrimitiveTypeContext) Char() antlr.TerminalNode {
	return s.GetToken(NebularisParserChar, 0)
}

func (s *PrimitiveTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitiveTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitiveTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitPrimitiveType(s)
	}
}

func (s *PrimitiveTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitPrimitiveType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) PrimitiveType() (localctx IPrimitiveTypeContext) {
	localctx = NewPrimitiveTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, NebularisParserRULE_primitiveType)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(394)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NebularisParserInt32:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(389)

			var _m = p.Match(NebularisParserInt32)

			localctx.(*PrimitiveTypeContext).tInt32 = _m
		}

	case NebularisParserBool:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(390)

			var _m = p.Match(NebularisParserBool)

			localctx.(*PrimitiveTypeContext).tBool = _m
		}

	case NebularisParserStr:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(391)

			var _m = p.Match(NebularisParserStr)

			localctx.(*PrimitiveTypeContext).tStr = _m
		}

	case NebularisParserByte:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(392)

			var _m = p.Match(NebularisParserByte)

			localctx.(*PrimitiveTypeContext).tByte = _m
		}

	case NebularisParserChar:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(393)

			var _m = p.Match(NebularisParserChar)

			localctx.(*PrimitiveTypeContext).tChar = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFieldContext is an interface to support dynamic dispatch.
type IFieldContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// Get_attribute returns the _attribute rule contexts.
	Get_attribute() IAttributeContext

	// GetTypeRef returns the typeRef rule contexts.
	GetTypeRef() ITypeSpecContext

	// Get_whereClause returns the _whereClause rule contexts.
	Get_whereClause() IWhereClauseContext

	// Set_attribute sets the _attribute rule contexts.
	Set_attribute(IAttributeContext)

	// SetTypeRef sets the typeRef rule contexts.
	SetTypeRef(ITypeSpecContext)

	// Set_whereClause sets the _whereClause rule contexts.
	Set_whereClause(IWhereClauseContext)

	// GetAttributes returns the attributes rule context list.
	GetAttributes() []IAttributeContext

	// GetConstraints returns the constraints rule context list.
	GetConstraints() []IWhereClauseContext

	// SetAttributes sets the attributes rule context list.
	SetAttributes([]IAttributeContext)

	// SetConstraints sets the constraints rule context list.
	SetConstraints([]IWhereClauseContext)

	// IsFieldContext differentiates from other interfaces.
	IsFieldContext()
}

type FieldContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	_attribute   IAttributeContext
	attributes   []IAttributeContext
	name         antlr.Token
	typeRef      ITypeSpecContext
	_whereClause IWhereClauseContext
	constraints  []IWhereClauseContext
}

func NewEmptyFieldContext() *FieldContext {
	var p = new(FieldContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_field
	return p
}

func (*FieldContext) IsFieldContext() {}

func NewFieldContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldContext {
	var p = new(FieldContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_field

	return p
}

func (s *FieldContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldContext) GetName() antlr.Token { return s.name }

func (s *FieldContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldContext) Get_attribute() IAttributeContext { return s._attribute }

func (s *FieldContext) GetTypeRef() ITypeSpecContext { return s.typeRef }

func (s *FieldContext) Get_whereClause() IWhereClauseContext { return s._whereClause }

func (s *FieldContext) Set_attribute(v IAttributeContext) { s._attribute = v }

func (s *FieldContext) SetTypeRef(v ITypeSpecContext) { s.typeRef = v }

func (s *FieldContext) Set_whereClause(v IWhereClauseContext) { s._whereClause = v }

func (s *FieldContext) GetAttributes() []IAttributeContext { return s.attributes }

func (s *FieldContext) GetConstraints() []IWhereClauseContext { return s.constraints }

func (s *FieldContext) SetAttributes(v []IAttributeContext) { s.attributes = v }

func (s *FieldContext) SetConstraints(v []IWhereClauseContext) { s.constraints = v }

func (s *FieldContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *FieldContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *FieldContext) SemiColon() antlr.TerminalNode {
	return s.GetToken(NebularisParserSemiColon, 0)
}

func (s *FieldContext) AllAttribute() []IAttributeContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IAttributeContext)(nil)).Elem())
	var tst = make([]IAttributeContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IAttributeContext)
		}
	}

	return tst
}

func (s *FieldContext) Attribute(i int) IAttributeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAttributeContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IAttributeContext)
}

func (s *FieldContext) AllWhereClause() []IWhereClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem())
	var tst = make([]IWhereClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereClauseContext)
		}
	}

	return tst
}

func (s *FieldContext) WhereClause(i int) IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *FieldContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterField(s)
	}
}

func (s *FieldContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitField(s)
	}
}

func (s *FieldContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitField(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Field() (localctx IFieldContext) {
	localctx = NewFieldContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, NebularisParserRULE_field)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(399)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserSquareOpen {
		{
			p.SetState(396)

			var _x = p.Attribute()

			localctx.(*FieldContext)._attribute = _x
		}
		localctx.(*FieldContext).attributes = append(localctx.(*FieldContext).attributes, localctx.(*FieldContext)._attribute)

		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(402)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*FieldContext).name = _m
	}
	{
		p.SetState(403)

		var _x = p.TypeSpec()

		localctx.(*FieldContext).typeRef = _x
	}
	p.SetState(407)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserWhere {
		{
			p.SetState(404)

			var _x = p.WhereClause()

			localctx.(*FieldContext)._whereClause = _x
		}
		localctx.(*FieldContext).constraints = append(localctx.(*FieldContext).constraints, localctx.(*FieldContext)._whereClause)

		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(411)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserSemiColon {
		{
			p.SetState(410)
			p.Match(NebularisParserSemiColon)
		}

	}

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) CopyFrom(ctx *ExpressionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type TernaryExpressionContext struct {
	*ExpressionContext
	predicate IExpressionContext
	onTrue    IExpressionContext
	onFalse   IExpressionContext
}

func NewTernaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *TernaryExpressionContext {
	var p = new(TernaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *TernaryExpressionContext) GetPredicate() IExpressionContext { return s.predicate }

func (s *TernaryExpressionContext) GetOnTrue() IExpressionContext { return s.onTrue }

func (s *TernaryExpressionContext) GetOnFalse() IExpressionContext { return s.onFalse }

func (s *TernaryExpressionContext) SetPredicate(v IExpressionContext) { s.predicate = v }

func (s *TernaryExpressionContext) SetOnTrue(v IExpressionContext) { s.onTrue = v }

func (s *TernaryExpressionContext) SetOnFalse(v IExpressionContext) { s.onFalse = v }

func (s *TernaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TernaryExpressionContext) Question() antlr.TerminalNode {
	return s.GetToken(NebularisParserQuestion, 0)
}

func (s *TernaryExpressionContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *TernaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *TernaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *TernaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitTernaryExpression(s)
	}
}

func (s *TernaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitTernaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type MemberAccessExpressionContext struct {
	*ExpressionContext
	target IExpressionContext
	member antlr.Token
}

func NewMemberAccessExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MemberAccessExpressionContext {
	var p = new(MemberAccessExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *MemberAccessExpressionContext) GetMember() antlr.Token { return s.member }

func (s *MemberAccessExpressionContext) SetMember(v antlr.Token) { s.member = v }

func (s *MemberAccessExpressionContext) GetTarget() IExpressionContext { return s.target }

func (s *MemberAccessExpressionContext) SetTarget(v IExpressionContext) { s.target = v }

func (s *MemberAccessExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MemberAccessExpressionContext) Dot() antlr.TerminalNode {
	return s.GetToken(NebularisParserDot, 0)
}

func (s *MemberAccessExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MemberAccessExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *MemberAccessExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterMemberAccessExpression(s)
	}
}

func (s *MemberAccessExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitMemberAccessExpression(s)
	}
}

func (s *MemberAccessExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitMemberAccessExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenthesisExpressionContext struct {
	*ExpressionContext
}

func NewParenthesisExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenthesisExpressionContext {
	var p = new(ParenthesisExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ParenthesisExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenthesisExpressionContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenOpen, 0)
}

func (s *ParenthesisExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ParenthesisExpressionContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenClose, 0)
}

func (s *ParenthesisExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitParenthesisExpression(s)
	}
}

func (s *ParenthesisExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitParenthesisExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LiteralExpressionContext struct {
	*ExpressionContext
	value ILiteralContext
}

func NewLiteralExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LiteralExpressionContext {
	var p = new(LiteralExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LiteralExpressionContext) GetValue() ILiteralContext { return s.value }

func (s *LiteralExpressionContext) SetValue(v ILiteralContext) { s.value = v }

func (s *LiteralExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralExpressionContext) Literal() ILiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILiteralContext)
}

func (s *LiteralExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitLiteralExpression(s)
	}
}

func (s *LiteralExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitLiteralExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IndexExpressionContext struct {
	*ExpressionContext
	target IExpressionContext
	index  IExpressionContext
}

func NewIndexExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IndexExpressionContext {
	var p = new(IndexExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IndexExpressionContext) GetTarget() IExpressionContext { return s.target }

func (s *IndexExpressionContext) GetIndex() IExpressionContext { return s.index }

func (s *IndexExpressionContext) SetTarget(v IExpressionContext) { s.target = v }

func (s *IndexExpressionContext) SetIndex(v IExpressionContext) { s.index = v }

func (s *IndexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IndexExpressionContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserSquareOpen, 0)
}

func (s *IndexExpressionContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserSquareClose, 0)
}

func (s *IndexExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *IndexExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IndexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterIndexExpression(s)
	}
}

func (s *IndexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitIndexExpression(s)
	}
}

func (s *IndexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitIndexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type UnaryExpressionContext struct {
	*ExpressionContext
	op IPrefixOpContext
}

func NewUnaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *UnaryExpressionContext) GetOp() IPrefixOpContext { return s.op }

func (s *UnaryExpressionContext) SetOp(v IPrefixOpContext) { s.op = v }

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *UnaryExpressionContext) PrefixOp() IPrefixOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrefixOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrefixOpContext)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type InvocationExpressionContext struct {
	*ExpressionContext
	target      IExpressionContext
	_expression IExpressionContext
	args        []IExpressionContext
}

func NewInvocationExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InvocationExpressionContext {
	var p = new(InvocationExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *InvocationExpressionContext) GetTarget() IExpressionContext { return s.target }

func (s *InvocationExpressionContext) Get_expression() IExpressionContext { return s._expression }

func (s *InvocationExpressionContext) SetTarget(v IExpressionContext) { s.target = v }

func (s *InvocationExpressionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InvocationExpressionContext) GetArgs() []IExpressionContext { return s.args }

func (s *InvocationExpressionContext) SetArgs(v []IExpressionContext) { s.args = v }

func (s *InvocationExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InvocationExpressionContext) ParenOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenOpen, 0)
}

func (s *InvocationExpressionContext) ParenClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserParenClose, 0)
}

func (s *InvocationExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *InvocationExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InvocationExpressionContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *InvocationExpressionContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *InvocationExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterInvocationExpression(s)
	}
}

func (s *InvocationExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitInvocationExpression(s)
	}
}

func (s *InvocationExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitInvocationExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdentifierExpressionContext struct {
	*ExpressionContext
	ident antlr.Token
}

func NewIdentifierExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdentifierExpressionContext {
	var p = new(IdentifierExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *IdentifierExpressionContext) GetIdent() antlr.Token { return s.ident }

func (s *IdentifierExpressionContext) SetIdent(v antlr.Token) { s.ident = v }

func (s *IdentifierExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdentifierExpressionContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *IdentifierExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitIdentifierExpression(s)
	}
}

func (s *IdentifierExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitIdentifierExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type StructExpressionContext struct {
	*ExpressionContext
}

func NewStructExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StructExpressionContext {
	var p = new(StructExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *StructExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExpressionContext) StructExpr() IStructExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructExprContext)
}

func (s *StructExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterStructExpression(s)
	}
}

func (s *StructExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitStructExpression(s)
	}
}

func (s *StructExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitStructExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	op    IBinaryOpContext
	right IExpressionContext
}

func NewBinaryExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryExpressionContext {
	var p = new(BinaryExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *BinaryExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *BinaryExpressionContext) GetOp() IBinaryOpContext { return s.op }

func (s *BinaryExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *BinaryExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *BinaryExpressionContext) SetOp(v IBinaryOpContext) { s.op = v }

func (s *BinaryExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *BinaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *BinaryExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BinaryExpressionContext) BinaryOp() IBinaryOpContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBinaryOpContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBinaryOpContext)
}

func (s *BinaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitBinaryExpression(s)
	}
}

func (s *BinaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitBinaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type ValueExpressionContext struct {
	*ExpressionContext
}

func NewValueExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ValueExpressionContext {
	var p = new(ValueExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *ValueExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueExpressionContext) Value() antlr.TerminalNode {
	return s.GetToken(NebularisParserValue, 0)
}

func (s *ValueExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterValueExpression(s)
	}
}

func (s *ValueExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitValueExpression(s)
	}
}

func (s *ValueExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitValueExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type NotYetImplementedExpressionContext struct {
	*ExpressionContext
}

func NewNotYetImplementedExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NotYetImplementedExpressionContext {
	var p = new(NotYetImplementedExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *NotYetImplementedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NotYetImplementedExpressionContext) Ellipsis() antlr.TerminalNode {
	return s.GetToken(NebularisParserEllipsis, 0)
}

func (s *NotYetImplementedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterNotYetImplementedExpression(s)
	}
}

func (s *NotYetImplementedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitNotYetImplementedExpression(s)
	}
}

func (s *NotYetImplementedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitNotYetImplementedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type LambdaExpressionContext struct {
	*ExpressionContext
}

func NewLambdaExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LambdaExpressionContext {
	var p = new(LambdaExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LambdaExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExpressionContext) LambdaExpr() ILambdaExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILambdaExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILambdaExprContext)
}

func (s *LambdaExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterLambdaExpression(s)
	}
}

func (s *LambdaExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitLambdaExpression(s)
	}
}

func (s *LambdaExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitLambdaExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *NebularisParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, NebularisParserRULE_expression, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 55, p.GetParserRuleContext()) {
	case 1:
		localctx = NewNotYetImplementedExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(414)
			p.Match(NebularisParserEllipsis)
		}

	case 2:
		localctx = NewIdentifierExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(415)

			var _m = p.Match(NebularisParserIdentifier)

			localctx.(*IdentifierExpressionContext).ident = _m
		}

	case 3:
		localctx = NewLiteralExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(416)

			var _x = p.Literal()

			localctx.(*LiteralExpressionContext).value = _x
		}

	case 4:
		localctx = NewParenthesisExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(417)
			p.Match(NebularisParserParenOpen)
		}
		{
			p.SetState(418)
			p.expression(0)
		}
		{
			p.SetState(419)
			p.Match(NebularisParserParenClose)
		}

	case 5:
		localctx = NewValueExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(421)
			p.Match(NebularisParserValue)
		}

	case 6:
		localctx = NewStructExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(422)
			p.StructExpr()
		}

	case 7:
		localctx = NewLambdaExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(423)
			p.LambdaExpr()
		}

	case 8:
		localctx = NewUnaryExpressionContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(424)

			var _x = p.PrefixOp()

			localctx.(*UnaryExpressionContext).op = _x
		}
		{
			p.SetState(425)
			p.expression(2)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(462)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(460)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 58, p.GetParserRuleContext()) {
			case 1:
				localctx = NewTernaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*TernaryExpressionContext).predicate = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NebularisParserRULE_expression)
				p.SetState(429)

				if !(p.Precpred(p.GetParserRuleContext(), 8)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 8)", ""))
				}
				{
					p.SetState(430)
					p.Match(NebularisParserQuestion)
				}
				{
					p.SetState(431)

					var _x = p.expression(0)

					localctx.(*TernaryExpressionContext).onTrue = _x
				}
				{
					p.SetState(432)
					p.Match(NebularisParserColon)
				}
				{
					p.SetState(433)

					var _x = p.expression(9)

					localctx.(*TernaryExpressionContext).onFalse = _x
				}

			case 2:
				localctx = NewBinaryExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*BinaryExpressionContext).left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NebularisParserRULE_expression)
				p.SetState(435)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(436)

					var _x = p.BinaryOp()

					localctx.(*BinaryExpressionContext).op = _x
				}
				{
					p.SetState(437)

					var _x = p.expression(2)

					localctx.(*BinaryExpressionContext).right = _x
				}

			case 3:
				localctx = NewMemberAccessExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*MemberAccessExpressionContext).target = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NebularisParserRULE_expression)
				p.SetState(439)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(440)
					p.Match(NebularisParserDot)
				}
				{
					p.SetState(441)

					var _m = p.Match(NebularisParserIdentifier)

					localctx.(*MemberAccessExpressionContext).member = _m
				}

			case 4:
				localctx = NewInvocationExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*InvocationExpressionContext).target = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NebularisParserRULE_expression)
				p.SetState(442)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(443)
					p.Match(NebularisParserParenOpen)
				}
				p.SetState(452)
				p.GetErrorHandler().Sync(p)
				_la = p.GetTokenStream().LA(1)

				if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NebularisParserIntegerLiteral)|(1<<NebularisParserBoolLiteral)|(1<<NebularisParserStringLiteral)|(1<<NebularisParserFn)|(1<<NebularisParserValue))) != 0) || (((_la-42)&-(0x1f+1)) == 0 && ((1<<uint((_la-42)))&((1<<(NebularisParserEllipsis-42))|(1<<(NebularisParserNot-42))|(1<<(NebularisParserParenOpen-42))|(1<<(NebularisParserTilda-42))|(1<<(NebularisParserIdentifier-42)))) != 0) {
					{
						p.SetState(444)

						var _x = p.expression(0)

						localctx.(*InvocationExpressionContext)._expression = _x
					}
					localctx.(*InvocationExpressionContext).args = append(localctx.(*InvocationExpressionContext).args, localctx.(*InvocationExpressionContext)._expression)
					p.SetState(449)
					p.GetErrorHandler().Sync(p)
					_la = p.GetTokenStream().LA(1)

					for _la == NebularisParserComma {
						{
							p.SetState(445)
							p.Match(NebularisParserComma)
						}
						{
							p.SetState(446)

							var _x = p.expression(0)

							localctx.(*InvocationExpressionContext)._expression = _x
						}
						localctx.(*InvocationExpressionContext).args = append(localctx.(*InvocationExpressionContext).args, localctx.(*InvocationExpressionContext)._expression)

						p.SetState(451)
						p.GetErrorHandler().Sync(p)
						_la = p.GetTokenStream().LA(1)
					}

				}
				{
					p.SetState(454)
					p.Match(NebularisParserParenClose)
				}

			case 5:
				localctx = NewIndexExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
				localctx.(*IndexExpressionContext).target = _prevctx

				p.PushNewRecursionContext(localctx, _startState, NebularisParserRULE_expression)
				p.SetState(455)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(456)
					p.Match(NebularisParserSquareOpen)
				}
				{
					p.SetState(457)

					var _x = p.expression(0)

					localctx.(*IndexExpressionContext).index = _x
				}
				{
					p.SetState(458)
					p.Match(NebularisParserSquareClose)
				}

			}

		}
		p.SetState(464)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 59, p.GetParserRuleContext())
	}

	return localctx
}

// ILambdaExprContext is an interface to support dynamic dispatch.
type ILambdaExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeParams returns the typeParams rule contexts.
	GetTypeParams() ITypeParametersContext

	// GetParams returns the params rule contexts.
	GetParams() IFunctionParametersContext

	// GetReturnType returns the returnType rule contexts.
	GetReturnType() ITypeSpecContext

	// Get_whereClause returns the _whereClause rule contexts.
	Get_whereClause() IWhereClauseContext

	// GetBody returns the body rule contexts.
	GetBody() ICodeBlockContext

	// SetTypeParams sets the typeParams rule contexts.
	SetTypeParams(ITypeParametersContext)

	// SetParams sets the params rule contexts.
	SetParams(IFunctionParametersContext)

	// SetReturnType sets the returnType rule contexts.
	SetReturnType(ITypeSpecContext)

	// Set_whereClause sets the _whereClause rule contexts.
	Set_whereClause(IWhereClauseContext)

	// SetBody sets the body rule contexts.
	SetBody(ICodeBlockContext)

	// GetConstraints returns the constraints rule context list.
	GetConstraints() []IWhereClauseContext

	// SetConstraints sets the constraints rule context list.
	SetConstraints([]IWhereClauseContext)

	// IsLambdaExprContext differentiates from other interfaces.
	IsLambdaExprContext()
}

type LambdaExprContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	typeParams   ITypeParametersContext
	params       IFunctionParametersContext
	returnType   ITypeSpecContext
	_whereClause IWhereClauseContext
	constraints  []IWhereClauseContext
	body         ICodeBlockContext
}

func NewEmptyLambdaExprContext() *LambdaExprContext {
	var p = new(LambdaExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_lambdaExpr
	return p
}

func (*LambdaExprContext) IsLambdaExprContext() {}

func NewLambdaExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LambdaExprContext {
	var p = new(LambdaExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_lambdaExpr

	return p
}

func (s *LambdaExprContext) GetParser() antlr.Parser { return s.parser }

func (s *LambdaExprContext) GetTypeParams() ITypeParametersContext { return s.typeParams }

func (s *LambdaExprContext) GetParams() IFunctionParametersContext { return s.params }

func (s *LambdaExprContext) GetReturnType() ITypeSpecContext { return s.returnType }

func (s *LambdaExprContext) Get_whereClause() IWhereClauseContext { return s._whereClause }

func (s *LambdaExprContext) GetBody() ICodeBlockContext { return s.body }

func (s *LambdaExprContext) SetTypeParams(v ITypeParametersContext) { s.typeParams = v }

func (s *LambdaExprContext) SetParams(v IFunctionParametersContext) { s.params = v }

func (s *LambdaExprContext) SetReturnType(v ITypeSpecContext) { s.returnType = v }

func (s *LambdaExprContext) Set_whereClause(v IWhereClauseContext) { s._whereClause = v }

func (s *LambdaExprContext) SetBody(v ICodeBlockContext) { s.body = v }

func (s *LambdaExprContext) GetConstraints() []IWhereClauseContext { return s.constraints }

func (s *LambdaExprContext) SetConstraints(v []IWhereClauseContext) { s.constraints = v }

func (s *LambdaExprContext) Fn() antlr.TerminalNode {
	return s.GetToken(NebularisParserFn, 0)
}

func (s *LambdaExprContext) FunctionParameters() IFunctionParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionParametersContext)
}

func (s *LambdaExprContext) CodeBlock() ICodeBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICodeBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICodeBlockContext)
}

func (s *LambdaExprContext) TypeParameters() ITypeParametersContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeParametersContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeParametersContext)
}

func (s *LambdaExprContext) TypeSpec() ITypeSpecContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypeSpecContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypeSpecContext)
}

func (s *LambdaExprContext) AllWhereClause() []IWhereClauseContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem())
	var tst = make([]IWhereClauseContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IWhereClauseContext)
		}
	}

	return tst
}

func (s *LambdaExprContext) WhereClause(i int) IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *LambdaExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LambdaExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LambdaExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterLambdaExpr(s)
	}
}

func (s *LambdaExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitLambdaExpr(s)
	}
}

func (s *LambdaExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitLambdaExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) LambdaExpr() (localctx ILambdaExprContext) {
	localctx = NewLambdaExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, NebularisParserRULE_lambdaExpr)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(465)
		p.Match(NebularisParserFn)
	}
	p.SetState(467)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserLessThan {
		{
			p.SetState(466)

			var _x = p.TypeParameters()

			localctx.(*LambdaExprContext).typeParams = _x
		}

	}
	{
		p.SetState(469)

		var _x = p.FunctionParameters()

		localctx.(*LambdaExprContext).params = _x
	}
	p.SetState(471)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<NebularisParserBool)|(1<<NebularisParserByte)|(1<<NebularisParserChar)|(1<<NebularisParserFn)|(1<<NebularisParserInterface)|(1<<NebularisParserInt32)|(1<<NebularisParserStr)|(1<<NebularisParserStruct))) != 0) || (((_la-59)&-(0x1f+1)) == 0 && ((1<<uint((_la-59)))&((1<<(NebularisParserQuestion-59))|(1<<(NebularisParserSquareOpen-59))|(1<<(NebularisParserIdentifier-59)))) != 0) {
		{
			p.SetState(470)

			var _x = p.TypeSpec()

			localctx.(*LambdaExprContext).returnType = _x
		}

	}
	p.SetState(476)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserWhere {
		{
			p.SetState(473)

			var _x = p.WhereClause()

			localctx.(*LambdaExprContext)._whereClause = _x
		}
		localctx.(*LambdaExprContext).constraints = append(localctx.(*LambdaExprContext).constraints, localctx.(*LambdaExprContext)._whereClause)

		p.SetState(478)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(479)

		var _x = p.CodeBlock()

		localctx.(*LambdaExprContext).body = _x
	}

	return localctx
}

// IStructExprContext is an interface to support dynamic dispatch.
type IStructExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTypeRef returns the typeRef rule contexts.
	GetTypeRef() IReferencedTypeContext

	// GetAssignments returns the assignments rule contexts.
	GetAssignments() IFieldAssignmentsContext

	// SetTypeRef sets the typeRef rule contexts.
	SetTypeRef(IReferencedTypeContext)

	// SetAssignments sets the assignments rule contexts.
	SetAssignments(IFieldAssignmentsContext)

	// IsStructExprContext differentiates from other interfaces.
	IsStructExprContext()
}

type StructExprContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	typeRef     IReferencedTypeContext
	assignments IFieldAssignmentsContext
}

func NewEmptyStructExprContext() *StructExprContext {
	var p = new(StructExprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_structExpr
	return p
}

func (*StructExprContext) IsStructExprContext() {}

func NewStructExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StructExprContext {
	var p = new(StructExprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_structExpr

	return p
}

func (s *StructExprContext) GetParser() antlr.Parser { return s.parser }

func (s *StructExprContext) GetTypeRef() IReferencedTypeContext { return s.typeRef }

func (s *StructExprContext) GetAssignments() IFieldAssignmentsContext { return s.assignments }

func (s *StructExprContext) SetTypeRef(v IReferencedTypeContext) { s.typeRef = v }

func (s *StructExprContext) SetAssignments(v IFieldAssignmentsContext) { s.assignments = v }

func (s *StructExprContext) CurlyOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyOpen, 0)
}

func (s *StructExprContext) CurlyClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserCurlyClose, 0)
}

func (s *StructExprContext) ReferencedType() IReferencedTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IReferencedTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IReferencedTypeContext)
}

func (s *StructExprContext) FieldAssignments() IFieldAssignmentsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldAssignmentsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFieldAssignmentsContext)
}

func (s *StructExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StructExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StructExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterStructExpr(s)
	}
}

func (s *StructExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitStructExpr(s)
	}
}

func (s *StructExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitStructExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) StructExpr() (localctx IStructExprContext) {
	localctx = NewStructExprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, NebularisParserRULE_structExpr)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(481)

		var _x = p.ReferencedType()

		localctx.(*StructExprContext).typeRef = _x
	}
	{
		p.SetState(482)
		p.Match(NebularisParserCurlyOpen)
	}
	{
		p.SetState(483)

		var _x = p.FieldAssignments()

		localctx.(*StructExprContext).assignments = _x
	}
	{
		p.SetState(484)
		p.Match(NebularisParserCurlyClose)
	}

	return localctx
}

// IFieldAssignmentsContext is an interface to support dynamic dispatch.
type IFieldAssignmentsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_fieldAssignment returns the _fieldAssignment rule contexts.
	Get_fieldAssignment() IFieldAssignmentContext

	// Set_fieldAssignment sets the _fieldAssignment rule contexts.
	Set_fieldAssignment(IFieldAssignmentContext)

	// GetFields returns the fields rule context list.
	GetFields() []IFieldAssignmentContext

	// SetFields sets the fields rule context list.
	SetFields([]IFieldAssignmentContext)

	// IsFieldAssignmentsContext differentiates from other interfaces.
	IsFieldAssignmentsContext()
}

type FieldAssignmentsContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	_fieldAssignment IFieldAssignmentContext
	fields           []IFieldAssignmentContext
}

func NewEmptyFieldAssignmentsContext() *FieldAssignmentsContext {
	var p = new(FieldAssignmentsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_fieldAssignments
	return p
}

func (*FieldAssignmentsContext) IsFieldAssignmentsContext() {}

func NewFieldAssignmentsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldAssignmentsContext {
	var p = new(FieldAssignmentsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_fieldAssignments

	return p
}

func (s *FieldAssignmentsContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldAssignmentsContext) Get_fieldAssignment() IFieldAssignmentContext {
	return s._fieldAssignment
}

func (s *FieldAssignmentsContext) Set_fieldAssignment(v IFieldAssignmentContext) {
	s._fieldAssignment = v
}

func (s *FieldAssignmentsContext) GetFields() []IFieldAssignmentContext { return s.fields }

func (s *FieldAssignmentsContext) SetFields(v []IFieldAssignmentContext) { s.fields = v }

func (s *FieldAssignmentsContext) AllFieldAssignment() []IFieldAssignmentContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IFieldAssignmentContext)(nil)).Elem())
	var tst = make([]IFieldAssignmentContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IFieldAssignmentContext)
		}
	}

	return tst
}

func (s *FieldAssignmentsContext) FieldAssignment(i int) IFieldAssignmentContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFieldAssignmentContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IFieldAssignmentContext)
}

func (s *FieldAssignmentsContext) AllComma() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserComma)
}

func (s *FieldAssignmentsContext) Comma(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserComma, i)
}

func (s *FieldAssignmentsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAssignmentsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldAssignmentsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterFieldAssignments(s)
	}
}

func (s *FieldAssignmentsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitFieldAssignments(s)
	}
}

func (s *FieldAssignmentsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitFieldAssignments(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) FieldAssignments() (localctx IFieldAssignmentsContext) {
	localctx = NewFieldAssignmentsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, NebularisParserRULE_fieldAssignments)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(495)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == NebularisParserIdentifier {
		p.SetState(491)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())

		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(486)

					var _x = p.FieldAssignment()

					localctx.(*FieldAssignmentsContext)._fieldAssignment = _x
				}
				localctx.(*FieldAssignmentsContext).fields = append(localctx.(*FieldAssignmentsContext).fields, localctx.(*FieldAssignmentsContext)._fieldAssignment)
				{
					p.SetState(487)
					p.Match(NebularisParserComma)
				}

			}
			p.SetState(493)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 63, p.GetParserRuleContext())
		}
		{
			p.SetState(494)

			var _x = p.FieldAssignment()

			localctx.(*FieldAssignmentsContext)._fieldAssignment = _x
		}
		localctx.(*FieldAssignmentsContext).fields = append(localctx.(*FieldAssignmentsContext).fields, localctx.(*FieldAssignmentsContext)._fieldAssignment)

	}

	return localctx
}

// IFieldAssignmentContext is an interface to support dynamic dispatch.
type IFieldAssignmentContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetName returns the name token.
	GetName() antlr.Token

	// SetName sets the name token.
	SetName(antlr.Token)

	// GetValue returns the value rule contexts.
	GetValue() IExpressionContext

	// SetValue sets the value rule contexts.
	SetValue(IExpressionContext)

	// IsFieldAssignmentContext differentiates from other interfaces.
	IsFieldAssignmentContext()
}

type FieldAssignmentContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	name   antlr.Token
	value  IExpressionContext
}

func NewEmptyFieldAssignmentContext() *FieldAssignmentContext {
	var p = new(FieldAssignmentContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_fieldAssignment
	return p
}

func (*FieldAssignmentContext) IsFieldAssignmentContext() {}

func NewFieldAssignmentContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FieldAssignmentContext {
	var p = new(FieldAssignmentContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_fieldAssignment

	return p
}

func (s *FieldAssignmentContext) GetParser() antlr.Parser { return s.parser }

func (s *FieldAssignmentContext) GetName() antlr.Token { return s.name }

func (s *FieldAssignmentContext) SetName(v antlr.Token) { s.name = v }

func (s *FieldAssignmentContext) GetValue() IExpressionContext { return s.value }

func (s *FieldAssignmentContext) SetValue(v IExpressionContext) { s.value = v }

func (s *FieldAssignmentContext) Colon() antlr.TerminalNode {
	return s.GetToken(NebularisParserColon, 0)
}

func (s *FieldAssignmentContext) Identifier() antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, 0)
}

func (s *FieldAssignmentContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FieldAssignmentContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FieldAssignmentContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FieldAssignmentContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterFieldAssignment(s)
	}
}

func (s *FieldAssignmentContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitFieldAssignment(s)
	}
}

func (s *FieldAssignmentContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitFieldAssignment(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) FieldAssignment() (localctx IFieldAssignmentContext) {
	localctx = NewFieldAssignmentContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, NebularisParserRULE_fieldAssignment)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(497)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*FieldAssignmentContext).name = _m
	}
	{
		p.SetState(498)
		p.Match(NebularisParserColon)
	}
	{
		p.SetState(499)

		var _x = p.expression(0)

		localctx.(*FieldAssignmentContext).value = _x
	}

	return localctx
}

// IPrefixOpContext is an interface to support dynamic dispatch.
type IPrefixOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPrefixOpContext differentiates from other interfaces.
	IsPrefixOpContext()
}

type PrefixOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixOpContext() *PrefixOpContext {
	var p = new(PrefixOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_prefixOp
	return p
}

func (*PrefixOpContext) IsPrefixOpContext() {}

func NewPrefixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixOpContext {
	var p = new(PrefixOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_prefixOp

	return p
}

func (s *PrefixOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixOpContext) Not() antlr.TerminalNode {
	return s.GetToken(NebularisParserNot, 0)
}

func (s *PrefixOpContext) Tilda() antlr.TerminalNode {
	return s.GetToken(NebularisParserTilda, 0)
}

func (s *PrefixOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterPrefixOp(s)
	}
}

func (s *PrefixOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitPrefixOp(s)
	}
}

func (s *PrefixOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitPrefixOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) PrefixOp() (localctx IPrefixOpContext) {
	localctx = NewPrefixOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, NebularisParserRULE_prefixOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(501)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NebularisParserNot || _la == NebularisParserTilda) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IPostfixOpContext is an interface to support dynamic dispatch.
type IPostfixOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPostfixOpContext differentiates from other interfaces.
	IsPostfixOpContext()
}

type PostfixOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPostfixOpContext() *PostfixOpContext {
	var p = new(PostfixOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_postfixOp
	return p
}

func (*PostfixOpContext) IsPostfixOpContext() {}

func NewPostfixOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PostfixOpContext {
	var p = new(PostfixOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_postfixOp

	return p
}

func (s *PostfixOpContext) GetParser() antlr.Parser { return s.parser }

func (s *PostfixOpContext) MinusMinus() antlr.TerminalNode {
	return s.GetToken(NebularisParserMinusMinus, 0)
}

func (s *PostfixOpContext) PlusPlus() antlr.TerminalNode {
	return s.GetToken(NebularisParserPlusPlus, 0)
}

func (s *PostfixOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PostfixOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PostfixOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterPostfixOp(s)
	}
}

func (s *PostfixOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitPostfixOp(s)
	}
}

func (s *PostfixOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitPostfixOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) PostfixOp() (localctx IPostfixOpContext) {
	localctx = NewPostfixOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, NebularisParserRULE_postfixOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(503)
		_la = p.GetTokenStream().LA(1)

		if !(_la == NebularisParserMinusMinus || _la == NebularisParserPlusPlus) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IBinaryOpContext is an interface to support dynamic dispatch.
type IBinaryOpContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBinaryOpContext differentiates from other interfaces.
	IsBinaryOpContext()
}

type BinaryOpContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBinaryOpContext() *BinaryOpContext {
	var p = new(BinaryOpContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_binaryOp
	return p
}

func (*BinaryOpContext) IsBinaryOpContext() {}

func NewBinaryOpContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BinaryOpContext {
	var p = new(BinaryOpContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_binaryOp

	return p
}

func (s *BinaryOpContext) GetParser() antlr.Parser { return s.parser }

func (s *BinaryOpContext) LessThan() antlr.TerminalNode {
	return s.GetToken(NebularisParserLessThan, 0)
}

func (s *BinaryOpContext) LessThanOrEqualTo() antlr.TerminalNode {
	return s.GetToken(NebularisParserLessThanOrEqualTo, 0)
}

func (s *BinaryOpContext) GreaterThan() antlr.TerminalNode {
	return s.GetToken(NebularisParserGreaterThan, 0)
}

func (s *BinaryOpContext) GreaterThanOrEqualTo() antlr.TerminalNode {
	return s.GetToken(NebularisParserGreaterThanOrEqualTo, 0)
}

func (s *BinaryOpContext) EqualsEquals() antlr.TerminalNode {
	return s.GetToken(NebularisParserEqualsEquals, 0)
}

func (s *BinaryOpContext) NotEquals() antlr.TerminalNode {
	return s.GetToken(NebularisParserNotEquals, 0)
}

func (s *BinaryOpContext) Plus() antlr.TerminalNode {
	return s.GetToken(NebularisParserPlus, 0)
}

func (s *BinaryOpContext) Minus() antlr.TerminalNode {
	return s.GetToken(NebularisParserMinus, 0)
}

func (s *BinaryOpContext) Asterisk() antlr.TerminalNode {
	return s.GetToken(NebularisParserAsterisk, 0)
}

func (s *BinaryOpContext) Slash() antlr.TerminalNode {
	return s.GetToken(NebularisParserSlash, 0)
}

func (s *BinaryOpContext) Modulus() antlr.TerminalNode {
	return s.GetToken(NebularisParserModulus, 0)
}

func (s *BinaryOpContext) Caret() antlr.TerminalNode {
	return s.GetToken(NebularisParserCaret, 0)
}

func (s *BinaryOpContext) OrOr() antlr.TerminalNode {
	return s.GetToken(NebularisParserOrOr, 0)
}

func (s *BinaryOpContext) AndAnd() antlr.TerminalNode {
	return s.GetToken(NebularisParserAndAnd, 0)
}

func (s *BinaryOpContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryOpContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BinaryOpContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterBinaryOp(s)
	}
}

func (s *BinaryOpContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitBinaryOp(s)
	}
}

func (s *BinaryOpContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitBinaryOp(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) BinaryOp() (localctx IBinaryOpContext) {
	localctx = NewBinaryOpContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, NebularisParserRULE_binaryOp)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(NebularisParserAndAnd-33))|(1<<(NebularisParserAsterisk-33))|(1<<(NebularisParserCaret-33))|(1<<(NebularisParserEqualsEquals-33))|(1<<(NebularisParserGreaterThan-33))|(1<<(NebularisParserGreaterThanOrEqualTo-33))|(1<<(NebularisParserLessThan-33))|(1<<(NebularisParserLessThanOrEqualTo-33))|(1<<(NebularisParserPlus-33))|(1<<(NebularisParserMinus-33))|(1<<(NebularisParserModulus-33))|(1<<(NebularisParserNotEquals-33))|(1<<(NebularisParserOrOr-33))|(1<<(NebularisParserSlash-33)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IAttributeContext is an interface to support dynamic dispatch.
type IAttributeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetValue returns the value rule contexts.
	GetValue() IStructExprContext

	// SetValue sets the value rule contexts.
	SetValue(IStructExprContext)

	// IsAttributeContext differentiates from other interfaces.
	IsAttributeContext()
}

type AttributeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	value  IStructExprContext
}

func NewEmptyAttributeContext() *AttributeContext {
	var p = new(AttributeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_attribute
	return p
}

func (*AttributeContext) IsAttributeContext() {}

func NewAttributeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AttributeContext {
	var p = new(AttributeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_attribute

	return p
}

func (s *AttributeContext) GetParser() antlr.Parser { return s.parser }

func (s *AttributeContext) GetValue() IStructExprContext { return s.value }

func (s *AttributeContext) SetValue(v IStructExprContext) { s.value = v }

func (s *AttributeContext) SquareOpen() antlr.TerminalNode {
	return s.GetToken(NebularisParserSquareOpen, 0)
}

func (s *AttributeContext) SquareClose() antlr.TerminalNode {
	return s.GetToken(NebularisParserSquareClose, 0)
}

func (s *AttributeContext) StructExpr() IStructExprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStructExprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStructExprContext)
}

func (s *AttributeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AttributeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AttributeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterAttribute(s)
	}
}

func (s *AttributeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitAttribute(s)
	}
}

func (s *AttributeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitAttribute(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Attribute() (localctx IAttributeContext) {
	localctx = NewAttributeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, NebularisParserRULE_attribute)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(507)
		p.Match(NebularisParserSquareOpen)
	}
	{
		p.SetState(508)

		var _x = p.StructExpr()

		localctx.(*AttributeContext).value = _x
	}
	{
		p.SetState(509)
		p.Match(NebularisParserSquareClose)
	}

	return localctx
}

// IQualifiedNameContext is an interface to support dynamic dispatch.
type IQualifiedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_Identifier returns the _Identifier token.
	Get_Identifier() antlr.Token

	// Set_Identifier sets the _Identifier token.
	Set_Identifier(antlr.Token)

	// GetParts returns the parts token list.
	GetParts() []antlr.Token

	// SetParts sets the parts token list.
	SetParts([]antlr.Token)

	// IsQualifiedNameContext differentiates from other interfaces.
	IsQualifiedNameContext()
}

type QualifiedNameContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	_Identifier antlr.Token
	parts       []antlr.Token
}

func NewEmptyQualifiedNameContext() *QualifiedNameContext {
	var p = new(QualifiedNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_qualifiedName
	return p
}

func (*QualifiedNameContext) IsQualifiedNameContext() {}

func NewQualifiedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QualifiedNameContext {
	var p = new(QualifiedNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_qualifiedName

	return p
}

func (s *QualifiedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *QualifiedNameContext) Get_Identifier() antlr.Token { return s._Identifier }

func (s *QualifiedNameContext) Set_Identifier(v antlr.Token) { s._Identifier = v }

func (s *QualifiedNameContext) GetParts() []antlr.Token { return s.parts }

func (s *QualifiedNameContext) SetParts(v []antlr.Token) { s.parts = v }

func (s *QualifiedNameContext) AllIdentifier() []antlr.TerminalNode {
	return s.GetTokens(NebularisParserIdentifier)
}

func (s *QualifiedNameContext) Identifier(i int) antlr.TerminalNode {
	return s.GetToken(NebularisParserIdentifier, i)
}

func (s *QualifiedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QualifiedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QualifiedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterQualifiedName(s)
	}
}

func (s *QualifiedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitQualifiedName(s)
	}
}

func (s *QualifiedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitQualifiedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) QualifiedName() (localctx IQualifiedNameContext) {
	localctx = NewQualifiedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, NebularisParserRULE_qualifiedName)
	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(511)

		var _m = p.Match(NebularisParserIdentifier)

		localctx.(*QualifiedNameContext)._Identifier = _m
	}
	localctx.(*QualifiedNameContext).parts = append(localctx.(*QualifiedNameContext).parts, localctx.(*QualifiedNameContext)._Identifier)
	p.SetState(516)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == NebularisParserDot {
		{
			p.SetState(512)
			p.Match(NebularisParserDot)
		}
		{
			p.SetState(513)

			var _m = p.Match(NebularisParserIdentifier)

			localctx.(*QualifiedNameContext)._Identifier = _m
		}
		localctx.(*QualifiedNameContext).parts = append(localctx.(*QualifiedNameContext).parts, localctx.(*QualifiedNameContext)._Identifier)

		p.SetState(518)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILiteralContext is an interface to support dynamic dispatch.
type ILiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetIntLit returns the intLit token.
	GetIntLit() antlr.Token

	// GetBoolLit returns the boolLit token.
	GetBoolLit() antlr.Token

	// GetStringLit returns the stringLit token.
	GetStringLit() antlr.Token

	// SetIntLit sets the intLit token.
	SetIntLit(antlr.Token)

	// SetBoolLit sets the boolLit token.
	SetBoolLit(antlr.Token)

	// SetStringLit sets the stringLit token.
	SetStringLit(antlr.Token)

	// IsLiteralContext differentiates from other interfaces.
	IsLiteralContext()
}

type LiteralContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	intLit    antlr.Token
	boolLit   antlr.Token
	stringLit antlr.Token
}

func NewEmptyLiteralContext() *LiteralContext {
	var p = new(LiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = NebularisParserRULE_literal
	return p
}

func (*LiteralContext) IsLiteralContext() {}

func NewLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LiteralContext {
	var p = new(LiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = NebularisParserRULE_literal

	return p
}

func (s *LiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *LiteralContext) GetIntLit() antlr.Token { return s.intLit }

func (s *LiteralContext) GetBoolLit() antlr.Token { return s.boolLit }

func (s *LiteralContext) GetStringLit() antlr.Token { return s.stringLit }

func (s *LiteralContext) SetIntLit(v antlr.Token) { s.intLit = v }

func (s *LiteralContext) SetBoolLit(v antlr.Token) { s.boolLit = v }

func (s *LiteralContext) SetStringLit(v antlr.Token) { s.stringLit = v }

func (s *LiteralContext) IntegerLiteral() antlr.TerminalNode {
	return s.GetToken(NebularisParserIntegerLiteral, 0)
}

func (s *LiteralContext) BoolLiteral() antlr.TerminalNode {
	return s.GetToken(NebularisParserBoolLiteral, 0)
}

func (s *LiteralContext) StringLiteral() antlr.TerminalNode {
	return s.GetToken(NebularisParserStringLiteral, 0)
}

func (s *LiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.EnterLiteral(s)
	}
}

func (s *LiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(NebularisListener); ok {
		listenerT.ExitLiteral(s)
	}
}

func (s *LiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case NebularisVisitor:
		return t.VisitLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *NebularisParser) Literal() (localctx ILiteralContext) {
	localctx = NewLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, NebularisParserRULE_literal)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(522)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case NebularisParserIntegerLiteral:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(519)

			var _m = p.Match(NebularisParserIntegerLiteral)

			localctx.(*LiteralContext).intLit = _m
		}

	case NebularisParserBoolLiteral:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(520)

			var _m = p.Match(NebularisParserBoolLiteral)

			localctx.(*LiteralContext).boolLit = _m
		}

	case NebularisParserStringLiteral:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(521)

			var _m = p.Match(NebularisParserStringLiteral)

			localctx.(*LiteralContext).stringLit = _m
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *NebularisParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 22:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *NebularisParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 8)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
