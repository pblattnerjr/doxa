// Code generated from LML.g4 by ANTLR 4.8. DO NOT EDIT.

package parser // LML

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 92, 523,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 4,
	50, 9, 50, 3, 2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 106, 10, 2, 13, 2, 14, 2,
	107, 3, 2, 3, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5, 3, 127, 10, 3, 3, 4, 7, 4, 130,
	10, 4, 12, 4, 14, 4, 133, 11, 4, 3, 4, 3, 4, 7, 4, 137, 10, 4, 12, 4, 14,
	4, 140, 11, 4, 3, 4, 5, 4, 143, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 6, 5, 153, 10, 5, 13, 5, 14, 5, 154, 5, 5, 157, 10, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 6, 5, 164, 10, 5, 13, 5, 14, 5, 165, 5, 5, 168,
	10, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 174, 10, 5, 3, 5, 3, 5, 5, 5, 178,
	10, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8,
	6, 8, 191, 10, 8, 13, 8, 14, 8, 192, 3, 8, 5, 8, 196, 10, 8, 3, 9, 3, 9,
	6, 9, 200, 10, 9, 13, 9, 14, 9, 201, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 6, 10, 211, 10, 10, 13, 10, 14, 10, 212, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13, 233, 10, 13, 3, 14, 3, 14, 6, 14,
	237, 10, 14, 13, 14, 14, 14, 238, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3,
	16, 5, 16, 247, 10, 16, 3, 17, 3, 17, 3, 17, 6, 17, 252, 10, 17, 13, 17,
	14, 17, 253, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 7, 18, 262, 10,
	18, 12, 18, 14, 18, 265, 11, 18, 3, 19, 3, 19, 5, 19, 269, 10, 19, 3, 20,
	3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 5, 22, 280, 10,
	22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25,
	3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3,
	28, 3, 28, 3, 28, 3, 28, 3, 29, 3, 29, 3, 29, 6, 29, 309, 10, 29, 13, 29,
	14, 29, 310, 3, 30, 3, 30, 3, 30, 6, 30, 316, 10, 30, 13, 30, 14, 30, 317,
	3, 31, 3, 31, 3, 31, 6, 31, 323, 10, 31, 13, 31, 14, 31, 324, 3, 32, 3,
	32, 3, 32, 6, 32, 330, 10, 32, 13, 32, 14, 32, 331, 3, 33, 3, 33, 3, 33,
	6, 33, 337, 10, 33, 13, 33, 14, 33, 338, 3, 34, 3, 34, 3, 34, 6, 34, 344,
	10, 34, 13, 34, 14, 34, 345, 3, 35, 3, 35, 3, 35, 3, 35, 3, 36, 3, 36,
	3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3,
	39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 7, 40, 370, 10, 40, 12, 40, 14,
	40, 373, 11, 40, 3, 40, 3, 40, 3, 41, 3, 41, 7, 41, 379, 10, 41, 12, 41,
	14, 41, 382, 11, 41, 3, 41, 7, 41, 385, 10, 41, 12, 41, 14, 41, 388, 11,
	41, 3, 41, 3, 41, 3, 42, 6, 42, 393, 10, 42, 13, 42, 14, 42, 394, 3, 42,
	7, 42, 398, 10, 42, 12, 42, 14, 42, 401, 11, 42, 3, 42, 6, 42, 404, 10,
	42, 13, 42, 14, 42, 405, 3, 42, 3, 42, 7, 42, 410, 10, 42, 12, 42, 14,
	42, 413, 11, 42, 3, 42, 3, 42, 5, 42, 417, 10, 42, 3, 43, 3, 43, 3, 43,
	3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 3,
	43, 3, 43, 5, 43, 434, 10, 43, 3, 44, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44,
	441, 10, 44, 3, 45, 3, 45, 3, 45, 7, 45, 446, 10, 45, 12, 45, 14, 45, 449,
	11, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46, 5, 46, 457, 10, 46, 3,
	47, 3, 47, 3, 47, 7, 47, 462, 10, 47, 12, 47, 14, 47, 465, 11, 47, 3, 48,
	3, 48, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 3, 50, 5, 50, 480, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 5,
	50, 498, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50,
	3, 50, 3, 50, 5, 50, 510, 10, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3, 50, 3,
	50, 7, 50, 518, 10, 50, 12, 50, 14, 50, 521, 11, 50, 3, 50, 2, 3, 98, 51,
	2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38,
	40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74,
	76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96, 98, 2, 6, 3, 2, 71, 77, 3,
	2, 78, 80, 4, 2, 51, 52, 67, 69, 3, 2, 31, 32, 2, 569, 2, 100, 3, 2, 2,
	2, 4, 126, 3, 2, 2, 2, 6, 142, 3, 2, 2, 2, 8, 177, 3, 2, 2, 2, 10, 179,
	3, 2, 2, 2, 12, 181, 3, 2, 2, 2, 14, 184, 3, 2, 2, 2, 16, 197, 3, 2, 2,
	2, 18, 205, 3, 2, 2, 2, 20, 214, 3, 2, 2, 2, 22, 217, 3, 2, 2, 2, 24, 232,
	3, 2, 2, 2, 26, 234, 3, 2, 2, 2, 28, 240, 3, 2, 2, 2, 30, 246, 3, 2, 2,
	2, 32, 248, 3, 2, 2, 2, 34, 258, 3, 2, 2, 2, 36, 268, 3, 2, 2, 2, 38, 270,
	3, 2, 2, 2, 40, 273, 3, 2, 2, 2, 42, 276, 3, 2, 2, 2, 44, 281, 3, 2, 2,
	2, 46, 285, 3, 2, 2, 2, 48, 289, 3, 2, 2, 2, 50, 293, 3, 2, 2, 2, 52, 297,
	3, 2, 2, 2, 54, 301, 3, 2, 2, 2, 56, 305, 3, 2, 2, 2, 58, 312, 3, 2, 2,
	2, 60, 319, 3, 2, 2, 2, 62, 326, 3, 2, 2, 2, 64, 333, 3, 2, 2, 2, 66, 340,
	3, 2, 2, 2, 68, 347, 3, 2, 2, 2, 70, 351, 3, 2, 2, 2, 72, 355, 3, 2, 2,
	2, 74, 359, 3, 2, 2, 2, 76, 363, 3, 2, 2, 2, 78, 367, 3, 2, 2, 2, 80, 376,
	3, 2, 2, 2, 82, 416, 3, 2, 2, 2, 84, 433, 3, 2, 2, 2, 86, 440, 3, 2, 2,
	2, 88, 442, 3, 2, 2, 2, 90, 456, 3, 2, 2, 2, 92, 458, 3, 2, 2, 2, 94, 466,
	3, 2, 2, 2, 96, 468, 3, 2, 2, 2, 98, 509, 3, 2, 2, 2, 100, 101, 5, 52,
	27, 2, 101, 102, 5, 74, 38, 2, 102, 103, 5, 70, 36, 2, 103, 105, 5, 6,
	4, 2, 104, 106, 5, 8, 5, 2, 105, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2,
	107, 105, 3, 2, 2, 2, 107, 108, 3, 2, 2, 2, 108, 109, 3, 2, 2, 2, 109,
	110, 7, 2, 2, 3, 110, 3, 3, 2, 2, 2, 111, 127, 5, 50, 26, 2, 112, 127,
	5, 54, 28, 2, 113, 127, 5, 76, 39, 2, 114, 127, 5, 46, 24, 2, 115, 127,
	5, 48, 25, 2, 116, 127, 5, 72, 37, 2, 117, 127, 5, 50, 26, 2, 118, 127,
	5, 44, 23, 2, 119, 127, 5, 68, 35, 2, 120, 127, 5, 56, 29, 2, 121, 127,
	5, 60, 31, 2, 122, 127, 5, 58, 30, 2, 123, 127, 5, 62, 32, 2, 124, 127,
	5, 64, 33, 2, 125, 127, 5, 66, 34, 2, 126, 111, 3, 2, 2, 2, 126, 112, 3,
	2, 2, 2, 126, 113, 3, 2, 2, 2, 126, 114, 3, 2, 2, 2, 126, 115, 3, 2, 2,
	2, 126, 116, 3, 2, 2, 2, 126, 117, 3, 2, 2, 2, 126, 118, 3, 2, 2, 2, 126,
	119, 3, 2, 2, 2, 126, 120, 3, 2, 2, 2, 126, 121, 3, 2, 2, 2, 126, 122,
	3, 2, 2, 2, 126, 123, 3, 2, 2, 2, 126, 124, 3, 2, 2, 2, 126, 125, 3, 2,
	2, 2, 127, 5, 3, 2, 2, 2, 128, 130, 5, 4, 3, 2, 129, 128, 3, 2, 2, 2, 130,
	133, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 131, 132, 3, 2, 2, 2, 132, 143,
	3, 2, 2, 2, 133, 131, 3, 2, 2, 2, 134, 138, 7, 34, 2, 2, 135, 137, 5, 4,
	3, 2, 136, 135, 3, 2, 2, 2, 137, 140, 3, 2, 2, 2, 138, 136, 3, 2, 2, 2,
	138, 139, 3, 2, 2, 2, 139, 141, 3, 2, 2, 2, 140, 138, 3, 2, 2, 2, 141,
	143, 7, 35, 2, 2, 142, 131, 3, 2, 2, 2, 142, 134, 3, 2, 2, 2, 143, 7, 3,
	2, 2, 2, 144, 178, 5, 20, 11, 2, 145, 178, 5, 14, 8, 2, 146, 178, 5, 12,
	7, 2, 147, 148, 7, 45, 2, 2, 148, 149, 5, 98, 50, 2, 149, 156, 5, 78, 40,
	2, 150, 152, 7, 46, 2, 2, 151, 153, 5, 78, 40, 2, 152, 151, 3, 2, 2, 2,
	153, 154, 3, 2, 2, 2, 154, 152, 3, 2, 2, 2, 154, 155, 3, 2, 2, 2, 155,
	157, 3, 2, 2, 2, 156, 150, 3, 2, 2, 2, 156, 157, 3, 2, 2, 2, 157, 178,
	3, 2, 2, 2, 158, 159, 7, 45, 2, 2, 159, 160, 5, 98, 50, 2, 160, 167, 5,
	78, 40, 2, 161, 163, 7, 3, 2, 2, 162, 164, 5, 78, 40, 2, 163, 162, 3, 2,
	2, 2, 164, 165, 3, 2, 2, 2, 165, 163, 3, 2, 2, 2, 165, 166, 3, 2, 2, 2,
	166, 168, 3, 2, 2, 2, 167, 161, 3, 2, 2, 2, 167, 168, 3, 2, 2, 2, 168,
	178, 3, 2, 2, 2, 169, 173, 7, 39, 2, 2, 170, 174, 5, 94, 48, 2, 171, 174,
	7, 65, 2, 2, 172, 174, 7, 66, 2, 2, 173, 170, 3, 2, 2, 2, 173, 171, 3,
	2, 2, 2, 173, 172, 3, 2, 2, 2, 174, 175, 3, 2, 2, 2, 175, 178, 5, 80, 41,
	2, 176, 178, 5, 78, 40, 2, 177, 144, 3, 2, 2, 2, 177, 145, 3, 2, 2, 2,
	177, 146, 3, 2, 2, 2, 177, 147, 3, 2, 2, 2, 177, 158, 3, 2, 2, 2, 177,
	169, 3, 2, 2, 2, 177, 176, 3, 2, 2, 2, 178, 9, 3, 2, 2, 2, 179, 180, 9,
	2, 2, 2, 180, 11, 3, 2, 2, 2, 181, 182, 7, 4, 2, 2, 182, 183, 7, 87, 2,
	2, 183, 13, 3, 2, 2, 2, 184, 190, 7, 81, 2, 2, 185, 191, 5, 22, 12, 2,
	186, 191, 5, 34, 18, 2, 187, 191, 5, 42, 22, 2, 188, 191, 5, 18, 10, 2,
	189, 191, 5, 16, 9, 2, 190, 185, 3, 2, 2, 2, 190, 186, 3, 2, 2, 2, 190,
	187, 3, 2, 2, 2, 190, 188, 3, 2, 2, 2, 190, 189, 3, 2, 2, 2, 191, 192,
	3, 2, 2, 2, 192, 190, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 195, 3, 2,
	2, 2, 194, 196, 7, 86, 2, 2, 195, 194, 3, 2, 2, 2, 195, 196, 3, 2, 2, 2,
	196, 15, 3, 2, 2, 2, 197, 199, 7, 36, 2, 2, 198, 200, 5, 18, 10, 2, 199,
	198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201, 199, 3, 2, 2, 2, 201, 202,
	3, 2, 2, 2, 202, 203, 3, 2, 2, 2, 203, 204, 7, 37, 2, 2, 204, 17, 3, 2,
	2, 2, 205, 210, 7, 82, 2, 2, 206, 211, 5, 22, 12, 2, 207, 211, 5, 34, 18,
	2, 208, 211, 5, 42, 22, 2, 209, 211, 5, 16, 9, 2, 210, 206, 3, 2, 2, 2,
	210, 207, 3, 2, 2, 2, 210, 208, 3, 2, 2, 2, 210, 209, 3, 2, 2, 2, 211,
	212, 3, 2, 2, 2, 212, 210, 3, 2, 2, 2, 212, 213, 3, 2, 2, 2, 213, 19, 3,
	2, 2, 2, 214, 215, 7, 5, 2, 2, 215, 216, 5, 42, 22, 2, 216, 21, 3, 2, 2,
	2, 217, 218, 7, 6, 2, 2, 218, 219, 7, 87, 2, 2, 219, 23, 3, 2, 2, 2, 220,
	233, 7, 53, 2, 2, 221, 233, 7, 54, 2, 2, 222, 233, 7, 55, 2, 2, 223, 224,
	7, 56, 2, 2, 224, 233, 7, 57, 2, 2, 225, 233, 7, 58, 2, 2, 226, 233, 7,
	59, 2, 2, 227, 228, 7, 60, 2, 2, 228, 233, 7, 61, 2, 2, 229, 233, 7, 62,
	2, 2, 230, 233, 7, 63, 2, 2, 231, 233, 7, 64, 2, 2, 232, 220, 3, 2, 2,
	2, 232, 221, 3, 2, 2, 2, 232, 222, 3, 2, 2, 2, 232, 223, 3, 2, 2, 2, 232,
	225, 3, 2, 2, 2, 232, 226, 3, 2, 2, 2, 232, 227, 3, 2, 2, 2, 232, 229,
	3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 232, 231, 3, 2, 2, 2, 233, 25, 3, 2,
	2, 2, 234, 236, 5, 28, 15, 2, 235, 237, 5, 30, 16, 2, 236, 235, 3, 2, 2,
	2, 237, 238, 3, 2, 2, 2, 238, 236, 3, 2, 2, 2, 238, 239, 3, 2, 2, 2, 239,
	27, 3, 2, 2, 2, 240, 241, 9, 3, 2, 2, 241, 29, 3, 2, 2, 2, 242, 247, 5,
	32, 17, 2, 243, 247, 7, 83, 2, 2, 244, 247, 7, 85, 2, 2, 245, 247, 5, 22,
	12, 2, 246, 242, 3, 2, 2, 2, 246, 243, 3, 2, 2, 2, 246, 244, 3, 2, 2, 2,
	246, 245, 3, 2, 2, 2, 247, 31, 3, 2, 2, 2, 248, 251, 7, 84, 2, 2, 249,
	252, 5, 42, 22, 2, 250, 252, 5, 34, 18, 2, 251, 249, 3, 2, 2, 2, 251, 250,
	3, 2, 2, 2, 252, 253, 3, 2, 2, 2, 253, 251, 3, 2, 2, 2, 253, 254, 3, 2,
	2, 2, 254, 255, 3, 2, 2, 2, 255, 256, 7, 7, 2, 2, 256, 257, 7, 70, 2, 2,
	257, 33, 3, 2, 2, 2, 258, 259, 7, 8, 2, 2, 259, 263, 7, 87, 2, 2, 260,
	262, 5, 36, 19, 2, 261, 260, 3, 2, 2, 2, 262, 265, 3, 2, 2, 2, 263, 261,
	3, 2, 2, 2, 263, 264, 3, 2, 2, 2, 264, 35, 3, 2, 2, 2, 265, 263, 3, 2,
	2, 2, 266, 269, 5, 38, 20, 2, 267, 269, 5, 40, 21, 2, 268, 266, 3, 2, 2,
	2, 268, 267, 3, 2, 2, 2, 269, 37, 3, 2, 2, 2, 270, 271, 7, 9, 2, 2, 271,
	272, 7, 70, 2, 2, 272, 39, 3, 2, 2, 2, 273, 274, 7, 10, 2, 2, 274, 275,
	7, 70, 2, 2, 275, 41, 3, 2, 2, 2, 276, 277, 7, 11, 2, 2, 277, 279, 7, 87,
	2, 2, 278, 280, 7, 86, 2, 2, 279, 278, 3, 2, 2, 2, 279, 280, 3, 2, 2, 2,
	280, 43, 3, 2, 2, 2, 281, 282, 7, 12, 2, 2, 282, 283, 7, 33, 2, 2, 283,
	284, 7, 87, 2, 2, 284, 45, 3, 2, 2, 2, 285, 286, 7, 13, 2, 2, 286, 287,
	7, 33, 2, 2, 287, 288, 7, 87, 2, 2, 288, 47, 3, 2, 2, 2, 289, 290, 7, 14,
	2, 2, 290, 291, 7, 33, 2, 2, 291, 292, 7, 87, 2, 2, 292, 49, 3, 2, 2, 2,
	293, 294, 7, 15, 2, 2, 294, 295, 7, 33, 2, 2, 295, 296, 7, 70, 2, 2, 296,
	51, 3, 2, 2, 2, 297, 298, 7, 16, 2, 2, 298, 299, 7, 33, 2, 2, 299, 300,
	7, 87, 2, 2, 300, 53, 3, 2, 2, 2, 301, 302, 7, 17, 2, 2, 302, 303, 7, 33,
	2, 2, 303, 304, 7, 70, 2, 2, 304, 55, 3, 2, 2, 2, 305, 306, 7, 18, 2, 2,
	306, 308, 7, 33, 2, 2, 307, 309, 5, 26, 14, 2, 308, 307, 3, 2, 2, 2, 309,
	310, 3, 2, 2, 2, 310, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 57, 3,
	2, 2, 2, 312, 313, 7, 19, 2, 2, 313, 315, 7, 33, 2, 2, 314, 316, 5, 26,
	14, 2, 315, 314, 3, 2, 2, 2, 316, 317, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2,
	317, 318, 3, 2, 2, 2, 318, 59, 3, 2, 2, 2, 319, 320, 7, 20, 2, 2, 320,
	322, 7, 33, 2, 2, 321, 323, 5, 26, 14, 2, 322, 321, 3, 2, 2, 2, 323, 324,
	3, 2, 2, 2, 324, 322, 3, 2, 2, 2, 324, 325, 3, 2, 2, 2, 325, 61, 3, 2,
	2, 2, 326, 327, 7, 21, 2, 2, 327, 329, 7, 33, 2, 2, 328, 330, 5, 26, 14,
	2, 329, 328, 3, 2, 2, 2, 330, 331, 3, 2, 2, 2, 331, 329, 3, 2, 2, 2, 331,
	332, 3, 2, 2, 2, 332, 63, 3, 2, 2, 2, 333, 334, 7, 22, 2, 2, 334, 336,
	7, 33, 2, 2, 335, 337, 5, 26, 14, 2, 336, 335, 3, 2, 2, 2, 337, 338, 3,
	2, 2, 2, 338, 336, 3, 2, 2, 2, 338, 339, 3, 2, 2, 2, 339, 65, 3, 2, 2,
	2, 340, 341, 7, 23, 2, 2, 341, 343, 7, 33, 2, 2, 342, 344, 5, 26, 14, 2,
	343, 342, 3, 2, 2, 2, 344, 345, 3, 2, 2, 2, 345, 343, 3, 2, 2, 2, 345,
	346, 3, 2, 2, 2, 346, 67, 3, 2, 2, 2, 347, 348, 7, 24, 2, 2, 348, 349,
	7, 33, 2, 2, 349, 350, 7, 70, 2, 2, 350, 69, 3, 2, 2, 2, 351, 352, 7, 25,
	2, 2, 352, 353, 7, 33, 2, 2, 353, 354, 7, 87, 2, 2, 354, 71, 3, 2, 2, 2,
	355, 356, 7, 26, 2, 2, 356, 357, 7, 33, 2, 2, 357, 358, 7, 87, 2, 2, 358,
	73, 3, 2, 2, 2, 359, 360, 7, 27, 2, 2, 360, 361, 7, 33, 2, 2, 361, 362,
	7, 87, 2, 2, 362, 75, 3, 2, 2, 2, 363, 364, 7, 28, 2, 2, 364, 365, 7, 33,
	2, 2, 365, 366, 7, 70, 2, 2, 366, 77, 3, 2, 2, 2, 367, 371, 7, 34, 2, 2,
	368, 370, 5, 8, 5, 2, 369, 368, 3, 2, 2, 2, 370, 373, 3, 2, 2, 2, 371,
	369, 3, 2, 2, 2, 371, 372, 3, 2, 2, 2, 372, 374, 3, 2, 2, 2, 373, 371,
	3, 2, 2, 2, 374, 375, 7, 35, 2, 2, 375, 79, 3, 2, 2, 2, 376, 380, 7, 34,
	2, 2, 377, 379, 5, 82, 42, 2, 378, 377, 3, 2, 2, 2, 379, 382, 3, 2, 2,
	2, 380, 378, 3, 2, 2, 2, 380, 381, 3, 2, 2, 2, 381, 386, 3, 2, 2, 2, 382,
	380, 3, 2, 2, 2, 383, 385, 5, 84, 43, 2, 384, 383, 3, 2, 2, 2, 385, 388,
	3, 2, 2, 2, 386, 384, 3, 2, 2, 2, 386, 387, 3, 2, 2, 2, 387, 389, 3, 2,
	2, 2, 388, 386, 3, 2, 2, 2, 389, 390, 7, 35, 2, 2, 390, 81, 3, 2, 2, 2,
	391, 393, 5, 84, 43, 2, 392, 391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394,
	392, 3, 2, 2, 2, 394, 395, 3, 2, 2, 2, 395, 399, 3, 2, 2, 2, 396, 398,
	5, 8, 5, 2, 397, 396, 3, 2, 2, 2, 398, 401, 3, 2, 2, 2, 399, 397, 3, 2,
	2, 2, 399, 400, 3, 2, 2, 2, 400, 417, 3, 2, 2, 2, 401, 399, 3, 2, 2, 2,
	402, 404, 5, 84, 43, 2, 403, 402, 3, 2, 2, 2, 404, 405, 3, 2, 2, 2, 405,
	403, 3, 2, 2, 2, 405, 406, 3, 2, 2, 2, 406, 407, 3, 2, 2, 2, 407, 411,
	7, 34, 2, 2, 408, 410, 5, 8, 5, 2, 409, 408, 3, 2, 2, 2, 410, 413, 3, 2,
	2, 2, 411, 409, 3, 2, 2, 2, 411, 412, 3, 2, 2, 2, 412, 414, 3, 2, 2, 2,
	413, 411, 3, 2, 2, 2, 414, 415, 7, 35, 2, 2, 415, 417, 3, 2, 2, 2, 416,
	392, 3, 2, 2, 2, 416, 403, 3, 2, 2, 2, 417, 83, 3, 2, 2, 2, 418, 419, 7,
	47, 2, 2, 419, 420, 5, 86, 44, 2, 420, 421, 7, 38, 2, 2, 421, 434, 3, 2,
	2, 2, 422, 423, 7, 47, 2, 2, 423, 424, 5, 90, 46, 2, 424, 425, 7, 38, 2,
	2, 425, 434, 3, 2, 2, 2, 426, 427, 7, 47, 2, 2, 427, 428, 5, 24, 13, 2,
	428, 429, 5, 86, 44, 2, 429, 430, 7, 38, 2, 2, 430, 434, 3, 2, 2, 2, 431,
	432, 7, 29, 2, 2, 432, 434, 7, 38, 2, 2, 433, 418, 3, 2, 2, 2, 433, 422,
	3, 2, 2, 2, 433, 426, 3, 2, 2, 2, 433, 431, 3, 2, 2, 2, 434, 85, 3, 2,
	2, 2, 435, 441, 7, 70, 2, 2, 436, 437, 7, 70, 2, 2, 437, 438, 7, 48, 2,
	2, 438, 441, 7, 70, 2, 2, 439, 441, 5, 88, 45, 2, 440, 435, 3, 2, 2, 2,
	440, 436, 3, 2, 2, 2, 440, 439, 3, 2, 2, 2, 441, 87, 3, 2, 2, 2, 442, 447,
	7, 70, 2, 2, 443, 444, 7, 30, 2, 2, 444, 446, 7, 70, 2, 2, 445, 443, 3,
	2, 2, 2, 446, 449, 3, 2, 2, 2, 447, 445, 3, 2, 2, 2, 447, 448, 3, 2, 2,
	2, 448, 89, 3, 2, 2, 2, 449, 447, 3, 2, 2, 2, 450, 457, 5, 10, 6, 2, 451,
	452, 5, 10, 6, 2, 452, 453, 7, 48, 2, 2, 453, 454, 5, 10, 6, 2, 454, 457,
	3, 2, 2, 2, 455, 457, 5, 92, 47, 2, 456, 450, 3, 2, 2, 2, 456, 451, 3,
	2, 2, 2, 456, 455, 3, 2, 2, 2, 457, 91, 3, 2, 2, 2, 458, 463, 5, 10, 6,
	2, 459, 460, 7, 30, 2, 2, 460, 462, 5, 10, 6, 2, 461, 459, 3, 2, 2, 2,
	462, 465, 3, 2, 2, 2, 463, 461, 3, 2, 2, 2, 463, 464, 3, 2, 2, 2, 464,
	93, 3, 2, 2, 2, 465, 463, 3, 2, 2, 2, 466, 467, 9, 4, 2, 2, 467, 95, 3,
	2, 2, 2, 468, 469, 5, 24, 13, 2, 469, 470, 7, 70, 2, 2, 470, 97, 3, 2,
	2, 2, 471, 472, 8, 50, 1, 2, 472, 479, 5, 94, 48, 2, 473, 474, 7, 43, 2,
	2, 474, 480, 7, 33, 2, 2, 475, 476, 7, 42, 2, 2, 476, 480, 7, 33, 2, 2,
	477, 480, 7, 42, 2, 2, 478, 480, 7, 43, 2, 2, 479, 473, 3, 2, 2, 2, 479,
	475, 3, 2, 2, 2, 479, 477, 3, 2, 2, 2, 479, 478, 3, 2, 2, 2, 480, 481,
	3, 2, 2, 2, 481, 482, 7, 70, 2, 2, 482, 510, 3, 2, 2, 2, 483, 484, 5, 94,
	48, 2, 484, 485, 9, 5, 2, 2, 485, 486, 7, 70, 2, 2, 486, 510, 3, 2, 2,
	2, 487, 488, 7, 65, 2, 2, 488, 489, 9, 5, 2, 2, 489, 510, 5, 96, 49, 2,
	490, 497, 7, 65, 2, 2, 491, 492, 7, 43, 2, 2, 492, 498, 7, 33, 2, 2, 493,
	494, 7, 42, 2, 2, 494, 498, 7, 33, 2, 2, 495, 498, 7, 42, 2, 2, 496, 498,
	7, 43, 2, 2, 497, 491, 3, 2, 2, 2, 497, 493, 3, 2, 2, 2, 497, 495, 3, 2,
	2, 2, 497, 496, 3, 2, 2, 2, 498, 499, 3, 2, 2, 2, 499, 510, 5, 96, 49,
	2, 500, 501, 7, 50, 2, 2, 501, 510, 5, 34, 18, 2, 502, 503, 7, 66, 2, 2,
	503, 504, 9, 5, 2, 2, 504, 510, 5, 10, 6, 2, 505, 506, 7, 36, 2, 2, 506,
	507, 5, 98, 50, 2, 507, 508, 7, 37, 2, 2, 508, 510, 3, 2, 2, 2, 509, 471,
	3, 2, 2, 2, 509, 483, 3, 2, 2, 2, 509, 487, 3, 2, 2, 2, 509, 490, 3, 2,
	2, 2, 509, 500, 3, 2, 2, 2, 509, 502, 3, 2, 2, 2, 509, 505, 3, 2, 2, 2,
	510, 519, 3, 2, 2, 2, 511, 512, 12, 5, 2, 2, 512, 513, 7, 40, 2, 2, 513,
	518, 5, 98, 50, 6, 514, 515, 12, 4, 2, 2, 515, 516, 7, 41, 2, 2, 516, 518,
	5, 98, 50, 5, 517, 511, 3, 2, 2, 2, 517, 514, 3, 2, 2, 2, 518, 521, 3,
	2, 2, 2, 519, 517, 3, 2, 2, 2, 519, 520, 3, 2, 2, 2, 520, 99, 3, 2, 2,
	2, 521, 519, 3, 2, 2, 2, 51, 107, 126, 131, 138, 142, 154, 156, 165, 167,
	173, 177, 190, 192, 195, 201, 210, 212, 232, 238, 246, 251, 253, 263, 268,
	279, 310, 317, 324, 331, 338, 345, 371, 380, 386, 394, 399, 405, 411, 416,
	433, 440, 447, 456, 463, 479, 497, 509, 517, 519,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'else if'", "'insert'", "'Media'", "'nid'", "'lang'", "'rid'", "'@Mode'",
	"'@Day'", "'sid'", "'Calendar'", "'HtmlCss'", "'PdfCss'", "'Day'", "'ID'",
	"'Month'", "'PageHeaderEven'", "'PageFooterEven'", "'PageHeaderOdd'", "'PageFooterOdd'",
	"'PageHeader'", "'PageFooter'", "'SetPageNumber'", "'Status'", "'Title'",
	"'Type'", "'Year'", "'default'", "','", "'=='", "'!='", "'='", "'{'", "'}'",
	"'('", "')'", "':'", "'switch'", "'&&'", "'||'", "'>'", "'<'", "'>='",
	"'if'", "'else'", "'case'", "'thru'", "'!'", "'Exists'", "'ModeOfWeek'",
	"'MovableCycleDay'", "'Jan'", "'Feb'", "'Mar'", "'Apr'", "'May'", "'Jun'",
	"'Jul'", "'Aug'", "'Sep'", "'Oct'", "'Nov'", "'Dec'", "'Date'", "'NameOfDay'",
	"'LukanCycleDay'", "'SundayAfterElevationOfCross'", "'SundaysBeforeTriodion'",
	"", "'Sun'", "'Mon'", "'Tue'", "'Wed'", "'Thu'", "'Fri'", "'Sat'", "'left'",
	"'center'", "'right'", "", "", "'@Date'", "'@Lookup'", "'@PageNbr'", "'@Ver'",
}
var symbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "", "", "EQUALITY", "INEQUALITY", "ASSIGNMENT",
	"LBRACE", "RBRACE", "LPAREN", "RPAREN", "COLON", "SWITCH", "AND", "OR",
	"GT", "LT", "GTOE", "IF", "ELSE", "CASE", "THRU", "NOT", "EXISTS", "MODE_OF_WEEK",
	"MOVABLE_CYCLE_DAY", "JAN", "FEB", "MAR", "APR", "MAY", "JUN", "JUL", "AUG",
	"SEP", "OCT", "NOV", "DEC", "DATE", "NAME_OF_DAY", "LUKAN_CYCLE_DAY", "SUNDAY_AFTER_ELEVATION_OF_CROSS",
	"SUNDAYS_BEFORE_TRIODION", "INTEGER", "SUN", "MON", "TUE", "WED", "THU",
	"FRI", "SAT", "LEFT", "CENTER", "RIGHT", "PARA_STYLE", "SPAN_STYLE", "INSERT_DATE",
	"INSERT_LOOKUP", "INSERT_PAGE_NUMBER", "INSERT_VER", "STRING", "COMMENT",
	"LINE_COMMENT", "EOL", "WS", "UNKNOWN_CHAR",
}

var ruleNames = []string{
	"template", "property", "propertyBlock", "statement", "dowName", "insert",
	"para", "pspan", "span", "media", "nid", "monthName", "position", "positionType",
	"directive", "lookup", "rid", "override", "overrideMode", "overrideDay",
	"sid", "tmplCalendar", "tmplHtmlCss", "tmplPdfCss", "tmplDay", "tmplID",
	"tmplMonth", "tmplPageHeaderEven", "tmplPageFooterEven", "tmplPageHeaderOdd",
	"tmplPageFooterOdd", "tmplPageHeader", "tmplPageFooter", "tmplPageNumber",
	"tmplStatus", "tmplTitle", "tmplType", "tmplYear", "block", "switchBlock",
	"switchBlockStatementGroup", "switchLabel", "integerExpression", "integerList",
	"dowExpression", "dowList", "ldpInt", "monthDay", "expression",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type LMLParser struct {
	*antlr.BaseParser
}

func NewLMLParser(input antlr.TokenStream) *LMLParser {
	this := new(LMLParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "LML.g4"

	return this
}

// LMLParser tokens.
const (
	LMLParserEOF                             = antlr.TokenEOF
	LMLParserT__0                            = 1
	LMLParserT__1                            = 2
	LMLParserT__2                            = 3
	LMLParserT__3                            = 4
	LMLParserT__4                            = 5
	LMLParserT__5                            = 6
	LMLParserT__6                            = 7
	LMLParserT__7                            = 8
	LMLParserT__8                            = 9
	LMLParserT__9                            = 10
	LMLParserT__10                           = 11
	LMLParserT__11                           = 12
	LMLParserT__12                           = 13
	LMLParserT__13                           = 14
	LMLParserT__14                           = 15
	LMLParserT__15                           = 16
	LMLParserT__16                           = 17
	LMLParserT__17                           = 18
	LMLParserT__18                           = 19
	LMLParserT__19                           = 20
	LMLParserT__20                           = 21
	LMLParserT__21                           = 22
	LMLParserT__22                           = 23
	LMLParserT__23                           = 24
	LMLParserT__24                           = 25
	LMLParserT__25                           = 26
	LMLParserT__26                           = 27
	LMLParserT__27                           = 28
	LMLParserEQUALITY                        = 29
	LMLParserINEQUALITY                      = 30
	LMLParserASSIGNMENT                      = 31
	LMLParserLBRACE                          = 32
	LMLParserRBRACE                          = 33
	LMLParserLPAREN                          = 34
	LMLParserRPAREN                          = 35
	LMLParserCOLON                           = 36
	LMLParserSWITCH                          = 37
	LMLParserAND                             = 38
	LMLParserOR                              = 39
	LMLParserGT                              = 40
	LMLParserLT                              = 41
	LMLParserGTOE                            = 42
	LMLParserIF                              = 43
	LMLParserELSE                            = 44
	LMLParserCASE                            = 45
	LMLParserTHRU                            = 46
	LMLParserNOT                             = 47
	LMLParserEXISTS                          = 48
	LMLParserMODE_OF_WEEK                    = 49
	LMLParserMOVABLE_CYCLE_DAY               = 50
	LMLParserJAN                             = 51
	LMLParserFEB                             = 52
	LMLParserMAR                             = 53
	LMLParserAPR                             = 54
	LMLParserMAY                             = 55
	LMLParserJUN                             = 56
	LMLParserJUL                             = 57
	LMLParserAUG                             = 58
	LMLParserSEP                             = 59
	LMLParserOCT                             = 60
	LMLParserNOV                             = 61
	LMLParserDEC                             = 62
	LMLParserDATE                            = 63
	LMLParserNAME_OF_DAY                     = 64
	LMLParserLUKAN_CYCLE_DAY                 = 65
	LMLParserSUNDAY_AFTER_ELEVATION_OF_CROSS = 66
	LMLParserSUNDAYS_BEFORE_TRIODION         = 67
	LMLParserINTEGER                         = 68
	LMLParserSUN                             = 69
	LMLParserMON                             = 70
	LMLParserTUE                             = 71
	LMLParserWED                             = 72
	LMLParserTHU                             = 73
	LMLParserFRI                             = 74
	LMLParserSAT                             = 75
	LMLParserLEFT                            = 76
	LMLParserCENTER                          = 77
	LMLParserRIGHT                           = 78
	LMLParserPARA_STYLE                      = 79
	LMLParserSPAN_STYLE                      = 80
	LMLParserINSERT_DATE                     = 81
	LMLParserINSERT_LOOKUP                   = 82
	LMLParserINSERT_PAGE_NUMBER              = 83
	LMLParserINSERT_VER                      = 84
	LMLParserSTRING                          = 85
	LMLParserCOMMENT                         = 86
	LMLParserLINE_COMMENT                    = 87
	LMLParserEOL                             = 88
	LMLParserWS                              = 89
	LMLParserUNKNOWN_CHAR                    = 90
)

// LMLParser rules.
const (
	LMLParserRULE_template                  = 0
	LMLParserRULE_property                  = 1
	LMLParserRULE_propertyBlock             = 2
	LMLParserRULE_statement                 = 3
	LMLParserRULE_dowName                   = 4
	LMLParserRULE_insert                    = 5
	LMLParserRULE_para                      = 6
	LMLParserRULE_pspan                     = 7
	LMLParserRULE_span                      = 8
	LMLParserRULE_media                     = 9
	LMLParserRULE_nid                       = 10
	LMLParserRULE_monthName                 = 11
	LMLParserRULE_position                  = 12
	LMLParserRULE_positionType              = 13
	LMLParserRULE_directive                 = 14
	LMLParserRULE_lookup                    = 15
	LMLParserRULE_rid                       = 16
	LMLParserRULE_override                  = 17
	LMLParserRULE_overrideMode              = 18
	LMLParserRULE_overrideDay               = 19
	LMLParserRULE_sid                       = 20
	LMLParserRULE_tmplCalendar              = 21
	LMLParserRULE_tmplHtmlCss               = 22
	LMLParserRULE_tmplPdfCss                = 23
	LMLParserRULE_tmplDay                   = 24
	LMLParserRULE_tmplID                    = 25
	LMLParserRULE_tmplMonth                 = 26
	LMLParserRULE_tmplPageHeaderEven        = 27
	LMLParserRULE_tmplPageFooterEven        = 28
	LMLParserRULE_tmplPageHeaderOdd         = 29
	LMLParserRULE_tmplPageFooterOdd         = 30
	LMLParserRULE_tmplPageHeader            = 31
	LMLParserRULE_tmplPageFooter            = 32
	LMLParserRULE_tmplPageNumber            = 33
	LMLParserRULE_tmplStatus                = 34
	LMLParserRULE_tmplTitle                 = 35
	LMLParserRULE_tmplType                  = 36
	LMLParserRULE_tmplYear                  = 37
	LMLParserRULE_block                     = 38
	LMLParserRULE_switchBlock               = 39
	LMLParserRULE_switchBlockStatementGroup = 40
	LMLParserRULE_switchLabel               = 41
	LMLParserRULE_integerExpression         = 42
	LMLParserRULE_integerList               = 43
	LMLParserRULE_dowExpression             = 44
	LMLParserRULE_dowList                   = 45
	LMLParserRULE_ldpInt                    = 46
	LMLParserRULE_monthDay                  = 47
	LMLParserRULE_expression                = 48
)

// ITemplateContext is an interface to support dynamic dispatch.
type ITemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTemplateContext differentiates from other interfaces.
	IsTemplateContext()
}

type TemplateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTemplateContext() *TemplateContext {
	var p = new(TemplateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_template
	return p
}

func (*TemplateContext) IsTemplateContext() {}

func NewTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TemplateContext {
	var p = new(TemplateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_template

	return p
}

func (s *TemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *TemplateContext) TmplID() ITmplIDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplIDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplIDContext)
}

func (s *TemplateContext) TmplType() ITmplTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplTypeContext)
}

func (s *TemplateContext) TmplStatus() ITmplStatusContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplStatusContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplStatusContext)
}

func (s *TemplateContext) PropertyBlock() IPropertyBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPropertyBlockContext)
}

func (s *TemplateContext) EOF() antlr.TerminalNode {
	return s.GetToken(LMLParserEOF, 0)
}

func (s *TemplateContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *TemplateContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *TemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTemplate(s)
	}
}

func (s *TemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTemplate(s)
	}
}

func (s *TemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Template() (localctx ITemplateContext) {
	localctx = NewTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, LMLParserRULE_template)
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
		p.SetState(98)
		p.TmplID()
	}
	{
		p.SetState(99)
		p.TmplType()
	}
	{
		p.SetState(100)
		p.TmplStatus()
	}
	{
		p.SetState(101)
		p.PropertyBlock()
	}
	p.SetState(103)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
		{
			p.SetState(102)
			p.Statement()
		}

		p.SetState(105)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(107)
		p.Match(LMLParserEOF)
	}

	return localctx
}

// IPropertyContext is an interface to support dynamic dispatch.
type IPropertyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyContext differentiates from other interfaces.
	IsPropertyContext()
}

type PropertyContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyContext() *PropertyContext {
	var p = new(PropertyContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_property
	return p
}

func (*PropertyContext) IsPropertyContext() {}

func NewPropertyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyContext {
	var p = new(PropertyContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_property

	return p
}

func (s *PropertyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyContext) TmplDay() ITmplDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplDayContext)
}

func (s *PropertyContext) TmplMonth() ITmplMonthContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplMonthContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplMonthContext)
}

func (s *PropertyContext) TmplYear() ITmplYearContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplYearContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplYearContext)
}

func (s *PropertyContext) TmplHtmlCss() ITmplHtmlCssContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplHtmlCssContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplHtmlCssContext)
}

func (s *PropertyContext) TmplPdfCss() ITmplPdfCssContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPdfCssContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPdfCssContext)
}

func (s *PropertyContext) TmplTitle() ITmplTitleContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplTitleContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplTitleContext)
}

func (s *PropertyContext) TmplCalendar() ITmplCalendarContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplCalendarContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplCalendarContext)
}

func (s *PropertyContext) TmplPageNumber() ITmplPageNumberContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageNumberContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageNumberContext)
}

func (s *PropertyContext) TmplPageHeaderEven() ITmplPageHeaderEvenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageHeaderEvenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageHeaderEvenContext)
}

func (s *PropertyContext) TmplPageHeaderOdd() ITmplPageHeaderOddContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageHeaderOddContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageHeaderOddContext)
}

func (s *PropertyContext) TmplPageFooterEven() ITmplPageFooterEvenContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageFooterEvenContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageFooterEvenContext)
}

func (s *PropertyContext) TmplPageFooterOdd() ITmplPageFooterOddContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageFooterOddContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageFooterOddContext)
}

func (s *PropertyContext) TmplPageHeader() ITmplPageHeaderContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageHeaderContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageHeaderContext)
}

func (s *PropertyContext) TmplPageFooter() ITmplPageFooterContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITmplPageFooterContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITmplPageFooterContext)
}

func (s *PropertyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterProperty(s)
	}
}

func (s *PropertyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitProperty(s)
	}
}

func (s *PropertyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitProperty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Property() (localctx IPropertyContext) {
	localctx = NewPropertyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, LMLParserRULE_property)

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

	p.SetState(124)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(109)
			p.TmplDay()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.TmplMonth()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(111)
			p.TmplYear()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(112)
			p.TmplHtmlCss()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(113)
			p.TmplPdfCss()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(114)
			p.TmplTitle()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(115)
			p.TmplDay()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(116)
			p.TmplCalendar()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(117)
			p.TmplPageNumber()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(118)
			p.TmplPageHeaderEven()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(119)
			p.TmplPageHeaderOdd()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(120)
			p.TmplPageFooterEven()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(121)
			p.TmplPageFooterOdd()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(122)
			p.TmplPageHeader()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(123)
			p.TmplPageFooter()
		}

	}

	return localctx
}

// IPropertyBlockContext is an interface to support dynamic dispatch.
type IPropertyBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPropertyBlockContext differentiates from other interfaces.
	IsPropertyBlockContext()
}

type PropertyBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyBlockContext() *PropertyBlockContext {
	var p = new(PropertyBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_propertyBlock
	return p
}

func (*PropertyBlockContext) IsPropertyBlockContext() {}

func NewPropertyBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyBlockContext {
	var p = new(PropertyBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_propertyBlock

	return p
}

func (s *PropertyBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyBlockContext) AllProperty() []IPropertyContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPropertyContext)(nil)).Elem())
	var tst = make([]IPropertyContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPropertyContext)
		}
	}

	return tst
}

func (s *PropertyBlockContext) Property(i int) IPropertyContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPropertyContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPropertyContext)
}

func (s *PropertyBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserLBRACE, 0)
}

func (s *PropertyBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserRBRACE, 0)
}

func (s *PropertyBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterPropertyBlock(s)
	}
}

func (s *PropertyBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitPropertyBlock(s)
	}
}

func (s *PropertyBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitPropertyBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) PropertyBlock() (localctx IPropertyBlockContext) {
	localctx = NewPropertyBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, LMLParserRULE_propertyBlock)
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

	p.SetState(140)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(129)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LMLParserT__9)|(1<<LMLParserT__10)|(1<<LMLParserT__11)|(1<<LMLParserT__12)|(1<<LMLParserT__14)|(1<<LMLParserT__15)|(1<<LMLParserT__16)|(1<<LMLParserT__17)|(1<<LMLParserT__18)|(1<<LMLParserT__19)|(1<<LMLParserT__20)|(1<<LMLParserT__21)|(1<<LMLParserT__23)|(1<<LMLParserT__25))) != 0 {
			{
				p.SetState(126)
				p.Property()
			}

			p.SetState(131)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(132)
			p.Match(LMLParserLBRACE)
		}
		p.SetState(136)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LMLParserT__9)|(1<<LMLParserT__10)|(1<<LMLParserT__11)|(1<<LMLParserT__12)|(1<<LMLParserT__14)|(1<<LMLParserT__15)|(1<<LMLParserT__16)|(1<<LMLParserT__17)|(1<<LMLParserT__18)|(1<<LMLParserT__19)|(1<<LMLParserT__20)|(1<<LMLParserT__21)|(1<<LMLParserT__23)|(1<<LMLParserT__25))) != 0 {
			{
				p.SetState(133)
				p.Property()
			}

			p.SetState(138)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(139)
			p.Match(LMLParserRBRACE)
		}

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
	p.RuleIndex = LMLParserRULE_statement
	return p
}

func (*StatementContext) IsStatementContext() {}

func NewStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatementContext {
	var p = new(StatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_statement

	return p
}

func (s *StatementContext) GetParser() antlr.Parser { return s.parser }

func (s *StatementContext) Media() IMediaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMediaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMediaContext)
}

func (s *StatementContext) Para() IParaContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IParaContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IParaContext)
}

func (s *StatementContext) Insert() IInsertContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInsertContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInsertContext)
}

func (s *StatementContext) IF() antlr.TerminalNode {
	return s.GetToken(LMLParserIF, 0)
}

func (s *StatementContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *StatementContext) AllBlock() []IBlockContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IBlockContext)(nil)).Elem())
	var tst = make([]IBlockContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IBlockContext)
		}
	}

	return tst
}

func (s *StatementContext) Block(i int) IBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBlockContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IBlockContext)
}

func (s *StatementContext) ELSE() antlr.TerminalNode {
	return s.GetToken(LMLParserELSE, 0)
}

func (s *StatementContext) SWITCH() antlr.TerminalNode {
	return s.GetToken(LMLParserSWITCH, 0)
}

func (s *StatementContext) SwitchBlock() ISwitchBlockContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBlockContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockContext)
}

func (s *StatementContext) LdpInt() ILdpIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILdpIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILdpIntContext)
}

func (s *StatementContext) DATE() antlr.TerminalNode {
	return s.GetToken(LMLParserDATE, 0)
}

func (s *StatementContext) NAME_OF_DAY() antlr.TerminalNode {
	return s.GetToken(LMLParserNAME_OF_DAY, 0)
}

func (s *StatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterStatement(s)
	}
}

func (s *StatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitStatement(s)
	}
}

func (s *StatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Statement() (localctx IStatementContext) {
	localctx = NewStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, LMLParserRULE_statement)
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

	p.SetState(175)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(142)
			p.Media()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(143)
			p.Para()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(144)
			p.Insert()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(145)
			p.Match(LMLParserIF)
		}
		{
			p.SetState(146)
			p.expression(0)
		}
		{
			p.SetState(147)
			p.Block()
		}
		p.SetState(154)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LMLParserELSE {
			{
				p.SetState(148)
				p.Match(LMLParserELSE)
			}
			p.SetState(150)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(149)
						p.Block()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(152)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(156)
			p.Match(LMLParserIF)
		}
		{
			p.SetState(157)
			p.expression(0)
		}
		{
			p.SetState(158)
			p.Block()
		}
		p.SetState(165)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LMLParserT__0 {
			{
				p.SetState(159)
				p.Match(LMLParserT__0)
			}
			p.SetState(161)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(160)
						p.Block()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(163)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(167)
			p.Match(LMLParserSWITCH)
		}
		p.SetState(171)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LMLParserMODE_OF_WEEK, LMLParserMOVABLE_CYCLE_DAY, LMLParserLUKAN_CYCLE_DAY, LMLParserSUNDAY_AFTER_ELEVATION_OF_CROSS, LMLParserSUNDAYS_BEFORE_TRIODION:
			{
				p.SetState(168)
				p.LdpInt()
			}

		case LMLParserDATE:
			{
				p.SetState(169)
				p.Match(LMLParserDATE)
			}

		case LMLParserNAME_OF_DAY:
			{
				p.SetState(170)
				p.Match(LMLParserNAME_OF_DAY)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(173)
			p.SwitchBlock()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(174)
			p.Block()
		}

	}

	return localctx
}

// IDowNameContext is an interface to support dynamic dispatch.
type IDowNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDowNameContext differentiates from other interfaces.
	IsDowNameContext()
}

type DowNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDowNameContext() *DowNameContext {
	var p = new(DowNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_dowName
	return p
}

func (*DowNameContext) IsDowNameContext() {}

func NewDowNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DowNameContext {
	var p = new(DowNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_dowName

	return p
}

func (s *DowNameContext) GetParser() antlr.Parser { return s.parser }

func (s *DowNameContext) SUN() antlr.TerminalNode {
	return s.GetToken(LMLParserSUN, 0)
}

func (s *DowNameContext) MON() antlr.TerminalNode {
	return s.GetToken(LMLParserMON, 0)
}

func (s *DowNameContext) TUE() antlr.TerminalNode {
	return s.GetToken(LMLParserTUE, 0)
}

func (s *DowNameContext) WED() antlr.TerminalNode {
	return s.GetToken(LMLParserWED, 0)
}

func (s *DowNameContext) THU() antlr.TerminalNode {
	return s.GetToken(LMLParserTHU, 0)
}

func (s *DowNameContext) FRI() antlr.TerminalNode {
	return s.GetToken(LMLParserFRI, 0)
}

func (s *DowNameContext) SAT() antlr.TerminalNode {
	return s.GetToken(LMLParserSAT, 0)
}

func (s *DowNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DowNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DowNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterDowName(s)
	}
}

func (s *DowNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitDowName(s)
	}
}

func (s *DowNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitDowName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) DowName() (localctx IDowNameContext) {
	localctx = NewDowNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, LMLParserRULE_dowName)
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
		p.SetState(177)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-69)&-(0x1f+1)) == 0 && ((1<<uint((_la-69)))&((1<<(LMLParserSUN-69))|(1<<(LMLParserMON-69))|(1<<(LMLParserTUE-69))|(1<<(LMLParserWED-69))|(1<<(LMLParserTHU-69))|(1<<(LMLParserFRI-69))|(1<<(LMLParserSAT-69)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IInsertContext is an interface to support dynamic dispatch.
type IInsertContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsInsertContext differentiates from other interfaces.
	IsInsertContext()
}

type InsertContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyInsertContext() *InsertContext {
	var p = new(InsertContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_insert
	return p
}

func (*InsertContext) IsInsertContext() {}

func NewInsertContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InsertContext {
	var p = new(InsertContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_insert

	return p
}

func (s *InsertContext) GetParser() antlr.Parser { return s.parser }

func (s *InsertContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *InsertContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InsertContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InsertContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterInsert(s)
	}
}

func (s *InsertContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitInsert(s)
	}
}

func (s *InsertContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitInsert(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Insert() (localctx IInsertContext) {
	localctx = NewInsertContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, LMLParserRULE_insert)

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
		p.SetState(179)
		p.Match(LMLParserT__1)
	}
	{
		p.SetState(180)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// IParaContext is an interface to support dynamic dispatch.
type IParaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsParaContext differentiates from other interfaces.
	IsParaContext()
}

type ParaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyParaContext() *ParaContext {
	var p = new(ParaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_para
	return p
}

func (*ParaContext) IsParaContext() {}

func NewParaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ParaContext {
	var p = new(ParaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_para

	return p
}

func (s *ParaContext) GetParser() antlr.Parser { return s.parser }

func (s *ParaContext) PARA_STYLE() antlr.TerminalNode {
	return s.GetToken(LMLParserPARA_STYLE, 0)
}

func (s *ParaContext) AllNid() []INidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INidContext)(nil)).Elem())
	var tst = make([]INidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INidContext)
		}
	}

	return tst
}

func (s *ParaContext) Nid(i int) INidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INidContext)
}

func (s *ParaContext) AllRid() []IRidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRidContext)(nil)).Elem())
	var tst = make([]IRidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRidContext)
		}
	}

	return tst
}

func (s *ParaContext) Rid(i int) IRidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRidContext)
}

func (s *ParaContext) AllSid() []ISidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISidContext)(nil)).Elem())
	var tst = make([]ISidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISidContext)
		}
	}

	return tst
}

func (s *ParaContext) Sid(i int) ISidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISidContext)
}

func (s *ParaContext) AllSpan() []ISpanContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpanContext)(nil)).Elem())
	var tst = make([]ISpanContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpanContext)
		}
	}

	return tst
}

func (s *ParaContext) Span(i int) ISpanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpanContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpanContext)
}

func (s *ParaContext) AllPspan() []IPspanContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPspanContext)(nil)).Elem())
	var tst = make([]IPspanContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPspanContext)
		}
	}

	return tst
}

func (s *ParaContext) Pspan(i int) IPspanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPspanContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPspanContext)
}

func (s *ParaContext) INSERT_VER() antlr.TerminalNode {
	return s.GetToken(LMLParserINSERT_VER, 0)
}

func (s *ParaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ParaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterPara(s)
	}
}

func (s *ParaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitPara(s)
	}
}

func (s *ParaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitPara(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Para() (localctx IParaContext) {
	localctx = NewParaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, LMLParserRULE_para)
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
		p.SetState(182)
		p.Match(LMLParserPARA_STYLE)
	}
	p.SetState(188)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-4)&-(0x1f+1)) == 0 && ((1<<uint((_la-4)))&((1<<(LMLParserT__3-4))|(1<<(LMLParserT__5-4))|(1<<(LMLParserT__8-4))|(1<<(LMLParserLPAREN-4)))) != 0) || _la == LMLParserSPAN_STYLE {
		p.SetState(188)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LMLParserT__3:
			{
				p.SetState(183)
				p.Nid()
			}

		case LMLParserT__5:
			{
				p.SetState(184)
				p.Rid()
			}

		case LMLParserT__8:
			{
				p.SetState(185)
				p.Sid()
			}

		case LMLParserSPAN_STYLE:
			{
				p.SetState(186)
				p.Span()
			}

		case LMLParserLPAREN:
			{
				p.SetState(187)
				p.Pspan()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(193)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LMLParserINSERT_VER {
		{
			p.SetState(192)
			p.Match(LMLParserINSERT_VER)
		}

	}

	return localctx
}

// IPspanContext is an interface to support dynamic dispatch.
type IPspanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPspanContext differentiates from other interfaces.
	IsPspanContext()
}

type PspanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPspanContext() *PspanContext {
	var p = new(PspanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_pspan
	return p
}

func (*PspanContext) IsPspanContext() {}

func NewPspanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PspanContext {
	var p = new(PspanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_pspan

	return p
}

func (s *PspanContext) GetParser() antlr.Parser { return s.parser }

func (s *PspanContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LMLParserLPAREN, 0)
}

func (s *PspanContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LMLParserRPAREN, 0)
}

func (s *PspanContext) AllSpan() []ISpanContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpanContext)(nil)).Elem())
	var tst = make([]ISpanContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpanContext)
		}
	}

	return tst
}

func (s *PspanContext) Span(i int) ISpanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpanContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpanContext)
}

func (s *PspanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PspanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PspanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterPspan(s)
	}
}

func (s *PspanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitPspan(s)
	}
}

func (s *PspanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitPspan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Pspan() (localctx IPspanContext) {
	localctx = NewPspanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, LMLParserRULE_pspan)
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
		p.SetState(195)
		p.Match(LMLParserLPAREN)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserSPAN_STYLE {
		{
			p.SetState(196)
			p.Span()
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(201)
		p.Match(LMLParserRPAREN)
	}

	return localctx
}

// ISpanContext is an interface to support dynamic dispatch.
type ISpanContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSpanContext differentiates from other interfaces.
	IsSpanContext()
}

type SpanContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySpanContext() *SpanContext {
	var p = new(SpanContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_span
	return p
}

func (*SpanContext) IsSpanContext() {}

func NewSpanContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SpanContext {
	var p = new(SpanContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_span

	return p
}

func (s *SpanContext) GetParser() antlr.Parser { return s.parser }

func (s *SpanContext) SPAN_STYLE() antlr.TerminalNode {
	return s.GetToken(LMLParserSPAN_STYLE, 0)
}

func (s *SpanContext) AllNid() []INidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*INidContext)(nil)).Elem())
	var tst = make([]INidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(INidContext)
		}
	}

	return tst
}

func (s *SpanContext) Nid(i int) INidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(INidContext)
}

func (s *SpanContext) AllRid() []IRidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRidContext)(nil)).Elem())
	var tst = make([]IRidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRidContext)
		}
	}

	return tst
}

func (s *SpanContext) Rid(i int) IRidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRidContext)
}

func (s *SpanContext) AllSid() []ISidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISidContext)(nil)).Elem())
	var tst = make([]ISidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISidContext)
		}
	}

	return tst
}

func (s *SpanContext) Sid(i int) ISidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISidContext)
}

func (s *SpanContext) AllPspan() []IPspanContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPspanContext)(nil)).Elem())
	var tst = make([]IPspanContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPspanContext)
		}
	}

	return tst
}

func (s *SpanContext) Pspan(i int) IPspanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPspanContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPspanContext)
}

func (s *SpanContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SpanContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SpanContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterSpan(s)
	}
}

func (s *SpanContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitSpan(s)
	}
}

func (s *SpanContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitSpan(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Span() (localctx ISpanContext) {
	localctx = NewSpanContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, LMLParserRULE_span)

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
	{
		p.SetState(203)
		p.Match(LMLParserSPAN_STYLE)
	}
	p.SetState(208)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(208)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case LMLParserT__3:
				{
					p.SetState(204)
					p.Nid()
				}

			case LMLParserT__5:
				{
					p.SetState(205)
					p.Rid()
				}

			case LMLParserT__8:
				{
					p.SetState(206)
					p.Sid()
				}

			case LMLParserLPAREN:
				{
					p.SetState(207)
					p.Pspan()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(210)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// IMediaContext is an interface to support dynamic dispatch.
type IMediaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMediaContext differentiates from other interfaces.
	IsMediaContext()
}

type MediaContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMediaContext() *MediaContext {
	var p = new(MediaContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_media
	return p
}

func (*MediaContext) IsMediaContext() {}

func NewMediaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MediaContext {
	var p = new(MediaContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_media

	return p
}

func (s *MediaContext) GetParser() antlr.Parser { return s.parser }

func (s *MediaContext) Sid() ISidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISidContext)
}

func (s *MediaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MediaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MediaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterMedia(s)
	}
}

func (s *MediaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitMedia(s)
	}
}

func (s *MediaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitMedia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Media() (localctx IMediaContext) {
	localctx = NewMediaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, LMLParserRULE_media)

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
		p.SetState(212)
		p.Match(LMLParserT__2)
	}
	{
		p.SetState(213)
		p.Sid()
	}

	return localctx
}

// INidContext is an interface to support dynamic dispatch.
type INidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsNidContext differentiates from other interfaces.
	IsNidContext()
}

type NidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNidContext() *NidContext {
	var p = new(NidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_nid
	return p
}

func (*NidContext) IsNidContext() {}

func NewNidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NidContext {
	var p = new(NidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_nid

	return p
}

func (s *NidContext) GetParser() antlr.Parser { return s.parser }

func (s *NidContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *NidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterNid(s)
	}
}

func (s *NidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitNid(s)
	}
}

func (s *NidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitNid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Nid() (localctx INidContext) {
	localctx = NewNidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, LMLParserRULE_nid)

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
		p.SetState(215)
		p.Match(LMLParserT__3)
	}
	{
		p.SetState(216)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// IMonthNameContext is an interface to support dynamic dispatch.
type IMonthNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthNameContext differentiates from other interfaces.
	IsMonthNameContext()
}

type MonthNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthNameContext() *MonthNameContext {
	var p = new(MonthNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_monthName
	return p
}

func (*MonthNameContext) IsMonthNameContext() {}

func NewMonthNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthNameContext {
	var p = new(MonthNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_monthName

	return p
}

func (s *MonthNameContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthNameContext) JAN() antlr.TerminalNode {
	return s.GetToken(LMLParserJAN, 0)
}

func (s *MonthNameContext) FEB() antlr.TerminalNode {
	return s.GetToken(LMLParserFEB, 0)
}

func (s *MonthNameContext) MAR() antlr.TerminalNode {
	return s.GetToken(LMLParserMAR, 0)
}

func (s *MonthNameContext) APR() antlr.TerminalNode {
	return s.GetToken(LMLParserAPR, 0)
}

func (s *MonthNameContext) MAY() antlr.TerminalNode {
	return s.GetToken(LMLParserMAY, 0)
}

func (s *MonthNameContext) JUN() antlr.TerminalNode {
	return s.GetToken(LMLParserJUN, 0)
}

func (s *MonthNameContext) JUL() antlr.TerminalNode {
	return s.GetToken(LMLParserJUL, 0)
}

func (s *MonthNameContext) AUG() antlr.TerminalNode {
	return s.GetToken(LMLParserAUG, 0)
}

func (s *MonthNameContext) SEP() antlr.TerminalNode {
	return s.GetToken(LMLParserSEP, 0)
}

func (s *MonthNameContext) OCT() antlr.TerminalNode {
	return s.GetToken(LMLParserOCT, 0)
}

func (s *MonthNameContext) NOV() antlr.TerminalNode {
	return s.GetToken(LMLParserNOV, 0)
}

func (s *MonthNameContext) DEC() antlr.TerminalNode {
	return s.GetToken(LMLParserDEC, 0)
}

func (s *MonthNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterMonthName(s)
	}
}

func (s *MonthNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitMonthName(s)
	}
}

func (s *MonthNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitMonthName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) MonthName() (localctx IMonthNameContext) {
	localctx = NewMonthNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, LMLParserRULE_monthName)

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

	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LMLParserJAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(218)
			p.Match(LMLParserJAN)
		}

	case LMLParserFEB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(219)
			p.Match(LMLParserFEB)
		}

	case LMLParserMAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(220)
			p.Match(LMLParserMAR)
		}

	case LMLParserAPR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(221)
			p.Match(LMLParserAPR)
		}
		{
			p.SetState(222)
			p.Match(LMLParserMAY)
		}

	case LMLParserJUN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(223)
			p.Match(LMLParserJUN)
		}

	case LMLParserJUL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(224)
			p.Match(LMLParserJUL)
		}

	case LMLParserAUG:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(225)
			p.Match(LMLParserAUG)
		}
		{
			p.SetState(226)
			p.Match(LMLParserSEP)
		}

	case LMLParserOCT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(227)
			p.Match(LMLParserOCT)
		}

	case LMLParserNOV:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(228)
			p.Match(LMLParserNOV)
		}

	case LMLParserDEC:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(229)
			p.Match(LMLParserDEC)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IPositionContext is an interface to support dynamic dispatch.
type IPositionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPositionContext differentiates from other interfaces.
	IsPositionContext()
}

type PositionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositionContext() *PositionContext {
	var p = new(PositionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_position
	return p
}

func (*PositionContext) IsPositionContext() {}

func NewPositionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PositionContext {
	var p = new(PositionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_position

	return p
}

func (s *PositionContext) GetParser() antlr.Parser { return s.parser }

func (s *PositionContext) PositionType() IPositionTypeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionTypeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPositionTypeContext)
}

func (s *PositionContext) AllDirective() []IDirectiveContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDirectiveContext)(nil)).Elem())
	var tst = make([]IDirectiveContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDirectiveContext)
		}
	}

	return tst
}

func (s *PositionContext) Directive(i int) IDirectiveContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDirectiveContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDirectiveContext)
}

func (s *PositionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PositionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterPosition(s)
	}
}

func (s *PositionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitPosition(s)
	}
}

func (s *PositionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitPosition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Position() (localctx IPositionContext) {
	localctx = NewPositionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, LMLParserRULE_position)
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
		p.SetState(232)
		p.PositionType()
	}
	p.SetState(234)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserT__3 || (((_la-81)&-(0x1f+1)) == 0 && ((1<<uint((_la-81)))&((1<<(LMLParserINSERT_DATE-81))|(1<<(LMLParserINSERT_LOOKUP-81))|(1<<(LMLParserINSERT_PAGE_NUMBER-81)))) != 0) {
		{
			p.SetState(233)
			p.Directive()
		}

		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IPositionTypeContext is an interface to support dynamic dispatch.
type IPositionTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPositionTypeContext differentiates from other interfaces.
	IsPositionTypeContext()
}

type PositionTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPositionTypeContext() *PositionTypeContext {
	var p = new(PositionTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_positionType
	return p
}

func (*PositionTypeContext) IsPositionTypeContext() {}

func NewPositionTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PositionTypeContext {
	var p = new(PositionTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_positionType

	return p
}

func (s *PositionTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *PositionTypeContext) LEFT() antlr.TerminalNode {
	return s.GetToken(LMLParserLEFT, 0)
}

func (s *PositionTypeContext) CENTER() antlr.TerminalNode {
	return s.GetToken(LMLParserCENTER, 0)
}

func (s *PositionTypeContext) RIGHT() antlr.TerminalNode {
	return s.GetToken(LMLParserRIGHT, 0)
}

func (s *PositionTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PositionTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PositionTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterPositionType(s)
	}
}

func (s *PositionTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitPositionType(s)
	}
}

func (s *PositionTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitPositionType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) PositionType() (localctx IPositionTypeContext) {
	localctx = NewPositionTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, LMLParserRULE_positionType)
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
		p.SetState(238)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IDirectiveContext is an interface to support dynamic dispatch.
type IDirectiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDirectiveContext differentiates from other interfaces.
	IsDirectiveContext()
}

type DirectiveContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDirectiveContext() *DirectiveContext {
	var p = new(DirectiveContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_directive
	return p
}

func (*DirectiveContext) IsDirectiveContext() {}

func NewDirectiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DirectiveContext {
	var p = new(DirectiveContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_directive

	return p
}

func (s *DirectiveContext) GetParser() antlr.Parser { return s.parser }

func (s *DirectiveContext) Lookup() ILookupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILookupContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILookupContext)
}

func (s *DirectiveContext) INSERT_DATE() antlr.TerminalNode {
	return s.GetToken(LMLParserINSERT_DATE, 0)
}

func (s *DirectiveContext) INSERT_PAGE_NUMBER() antlr.TerminalNode {
	return s.GetToken(LMLParserINSERT_PAGE_NUMBER, 0)
}

func (s *DirectiveContext) Nid() INidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*INidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(INidContext)
}

func (s *DirectiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DirectiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DirectiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterDirective(s)
	}
}

func (s *DirectiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitDirective(s)
	}
}

func (s *DirectiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitDirective(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Directive() (localctx IDirectiveContext) {
	localctx = NewDirectiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, LMLParserRULE_directive)

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

	p.SetState(244)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LMLParserINSERT_LOOKUP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(240)
			p.Lookup()
		}

	case LMLParserINSERT_DATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Match(LMLParserINSERT_DATE)
		}

	case LMLParserINSERT_PAGE_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(242)
			p.Match(LMLParserINSERT_PAGE_NUMBER)
		}

	case LMLParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(243)
			p.Nid()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILookupContext is an interface to support dynamic dispatch.
type ILookupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLookupContext differentiates from other interfaces.
	IsLookupContext()
}

type LookupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLookupContext() *LookupContext {
	var p = new(LookupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_lookup
	return p
}

func (*LookupContext) IsLookupContext() {}

func NewLookupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LookupContext {
	var p = new(LookupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_lookup

	return p
}

func (s *LookupContext) GetParser() antlr.Parser { return s.parser }

func (s *LookupContext) INSERT_LOOKUP() antlr.TerminalNode {
	return s.GetToken(LMLParserINSERT_LOOKUP, 0)
}

func (s *LookupContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *LookupContext) AllSid() []ISidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISidContext)(nil)).Elem())
	var tst = make([]ISidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISidContext)
		}
	}

	return tst
}

func (s *LookupContext) Sid(i int) ISidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISidContext)
}

func (s *LookupContext) AllRid() []IRidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IRidContext)(nil)).Elem())
	var tst = make([]IRidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IRidContext)
		}
	}

	return tst
}

func (s *LookupContext) Rid(i int) IRidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IRidContext)
}

func (s *LookupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LookupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LookupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterLookup(s)
	}
}

func (s *LookupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitLookup(s)
	}
}

func (s *LookupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitLookup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Lookup() (localctx ILookupContext) {
	localctx = NewLookupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, LMLParserRULE_lookup)
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
		p.SetState(246)
		p.Match(LMLParserINSERT_LOOKUP)
	}
	p.SetState(249)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserT__5 || _la == LMLParserT__8 {
		p.SetState(249)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LMLParserT__8:
			{
				p.SetState(247)
				p.Sid()
			}

		case LMLParserT__5:
			{
				p.SetState(248)
				p.Rid()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(251)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(253)
		p.Match(LMLParserT__4)
	}
	{
		p.SetState(254)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// IRidContext is an interface to support dynamic dispatch.
type IRidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsRidContext differentiates from other interfaces.
	IsRidContext()
}

type RidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRidContext() *RidContext {
	var p = new(RidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_rid
	return p
}

func (*RidContext) IsRidContext() {}

func NewRidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RidContext {
	var p = new(RidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_rid

	return p
}

func (s *RidContext) GetParser() antlr.Parser { return s.parser }

func (s *RidContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *RidContext) AllOverride() []IOverrideContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOverrideContext)(nil)).Elem())
	var tst = make([]IOverrideContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOverrideContext)
		}
	}

	return tst
}

func (s *RidContext) Override(i int) IOverrideContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOverrideContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOverrideContext)
}

func (s *RidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterRid(s)
	}
}

func (s *RidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitRid(s)
	}
}

func (s *RidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitRid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Rid() (localctx IRidContext) {
	localctx = NewRidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, LMLParserRULE_rid)

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
	{
		p.SetState(256)
		p.Match(LMLParserT__5)
	}
	{
		p.SetState(257)
		p.Match(LMLParserSTRING)
	}
	p.SetState(261)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(258)
				p.Override()
			}

		}
		p.SetState(263)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
	}

	return localctx
}

// IOverrideContext is an interface to support dynamic dispatch.
type IOverrideContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOverrideContext differentiates from other interfaces.
	IsOverrideContext()
}

type OverrideContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOverrideContext() *OverrideContext {
	var p = new(OverrideContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_override
	return p
}

func (*OverrideContext) IsOverrideContext() {}

func NewOverrideContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OverrideContext {
	var p = new(OverrideContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_override

	return p
}

func (s *OverrideContext) GetParser() antlr.Parser { return s.parser }

func (s *OverrideContext) OverrideMode() IOverrideModeContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOverrideModeContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOverrideModeContext)
}

func (s *OverrideContext) OverrideDay() IOverrideDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOverrideDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOverrideDayContext)
}

func (s *OverrideContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverrideContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OverrideContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterOverride(s)
	}
}

func (s *OverrideContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitOverride(s)
	}
}

func (s *OverrideContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitOverride(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Override() (localctx IOverrideContext) {
	localctx = NewOverrideContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, LMLParserRULE_override)

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

	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LMLParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(264)
			p.OverrideMode()
		}

	case LMLParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(265)
			p.OverrideDay()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IOverrideModeContext is an interface to support dynamic dispatch.
type IOverrideModeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOverrideModeContext differentiates from other interfaces.
	IsOverrideModeContext()
}

type OverrideModeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOverrideModeContext() *OverrideModeContext {
	var p = new(OverrideModeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_overrideMode
	return p
}

func (*OverrideModeContext) IsOverrideModeContext() {}

func NewOverrideModeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OverrideModeContext {
	var p = new(OverrideModeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_overrideMode

	return p
}

func (s *OverrideModeContext) GetParser() antlr.Parser { return s.parser }

func (s *OverrideModeContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *OverrideModeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverrideModeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OverrideModeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterOverrideMode(s)
	}
}

func (s *OverrideModeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitOverrideMode(s)
	}
}

func (s *OverrideModeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitOverrideMode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) OverrideMode() (localctx IOverrideModeContext) {
	localctx = NewOverrideModeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, LMLParserRULE_overrideMode)

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
		p.SetState(268)
		p.Match(LMLParserT__6)
	}
	{
		p.SetState(269)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// IOverrideDayContext is an interface to support dynamic dispatch.
type IOverrideDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOverrideDayContext differentiates from other interfaces.
	IsOverrideDayContext()
}

type OverrideDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOverrideDayContext() *OverrideDayContext {
	var p = new(OverrideDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_overrideDay
	return p
}

func (*OverrideDayContext) IsOverrideDayContext() {}

func NewOverrideDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OverrideDayContext {
	var p = new(OverrideDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_overrideDay

	return p
}

func (s *OverrideDayContext) GetParser() antlr.Parser { return s.parser }

func (s *OverrideDayContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *OverrideDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OverrideDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OverrideDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterOverrideDay(s)
	}
}

func (s *OverrideDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitOverrideDay(s)
	}
}

func (s *OverrideDayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitOverrideDay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) OverrideDay() (localctx IOverrideDayContext) {
	localctx = NewOverrideDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, LMLParserRULE_overrideDay)

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
		p.SetState(271)
		p.Match(LMLParserT__7)
	}
	{
		p.SetState(272)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// ISidContext is an interface to support dynamic dispatch.
type ISidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSidContext differentiates from other interfaces.
	IsSidContext()
}

type SidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySidContext() *SidContext {
	var p = new(SidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_sid
	return p
}

func (*SidContext) IsSidContext() {}

func NewSidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SidContext {
	var p = new(SidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_sid

	return p
}

func (s *SidContext) GetParser() antlr.Parser { return s.parser }

func (s *SidContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *SidContext) INSERT_VER() antlr.TerminalNode {
	return s.GetToken(LMLParserINSERT_VER, 0)
}

func (s *SidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterSid(s)
	}
}

func (s *SidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitSid(s)
	}
}

func (s *SidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitSid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Sid() (localctx ISidContext) {
	localctx = NewSidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, LMLParserRULE_sid)

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
		p.SetState(274)
		p.Match(LMLParserT__8)
	}
	{
		p.SetState(275)
		p.Match(LMLParserSTRING)
	}
	p.SetState(277)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(276)
			p.Match(LMLParserINSERT_VER)
		}

	}

	return localctx
}

// ITmplCalendarContext is an interface to support dynamic dispatch.
type ITmplCalendarContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplCalendarContext differentiates from other interfaces.
	IsTmplCalendarContext()
}

type TmplCalendarContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplCalendarContext() *TmplCalendarContext {
	var p = new(TmplCalendarContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplCalendar
	return p
}

func (*TmplCalendarContext) IsTmplCalendarContext() {}

func NewTmplCalendarContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplCalendarContext {
	var p = new(TmplCalendarContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplCalendar

	return p
}

func (s *TmplCalendarContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplCalendarContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplCalendarContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplCalendarContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplCalendarContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplCalendarContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplCalendar(s)
	}
}

func (s *TmplCalendarContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplCalendar(s)
	}
}

func (s *TmplCalendarContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplCalendar(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplCalendar() (localctx ITmplCalendarContext) {
	localctx = NewTmplCalendarContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, LMLParserRULE_tmplCalendar)

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
		p.SetState(279)
		p.Match(LMLParserT__9)
	}
	{
		p.SetState(280)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(281)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplHtmlCssContext is an interface to support dynamic dispatch.
type ITmplHtmlCssContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplHtmlCssContext differentiates from other interfaces.
	IsTmplHtmlCssContext()
}

type TmplHtmlCssContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplHtmlCssContext() *TmplHtmlCssContext {
	var p = new(TmplHtmlCssContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplHtmlCss
	return p
}

func (*TmplHtmlCssContext) IsTmplHtmlCssContext() {}

func NewTmplHtmlCssContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplHtmlCssContext {
	var p = new(TmplHtmlCssContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplHtmlCss

	return p
}

func (s *TmplHtmlCssContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplHtmlCssContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplHtmlCssContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplHtmlCssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplHtmlCssContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplHtmlCssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplHtmlCss(s)
	}
}

func (s *TmplHtmlCssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplHtmlCss(s)
	}
}

func (s *TmplHtmlCssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplHtmlCss(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplHtmlCss() (localctx ITmplHtmlCssContext) {
	localctx = NewTmplHtmlCssContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, LMLParserRULE_tmplHtmlCss)

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
		p.SetState(283)
		p.Match(LMLParserT__10)
	}
	{
		p.SetState(284)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(285)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplPdfCssContext is an interface to support dynamic dispatch.
type ITmplPdfCssContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPdfCssContext differentiates from other interfaces.
	IsTmplPdfCssContext()
}

type TmplPdfCssContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPdfCssContext() *TmplPdfCssContext {
	var p = new(TmplPdfCssContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPdfCss
	return p
}

func (*TmplPdfCssContext) IsTmplPdfCssContext() {}

func NewTmplPdfCssContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPdfCssContext {
	var p = new(TmplPdfCssContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPdfCss

	return p
}

func (s *TmplPdfCssContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPdfCssContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPdfCssContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplPdfCssContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPdfCssContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPdfCssContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPdfCss(s)
	}
}

func (s *TmplPdfCssContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPdfCss(s)
	}
}

func (s *TmplPdfCssContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPdfCss(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPdfCss() (localctx ITmplPdfCssContext) {
	localctx = NewTmplPdfCssContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, LMLParserRULE_tmplPdfCss)

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
		p.SetState(287)
		p.Match(LMLParserT__11)
	}
	{
		p.SetState(288)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(289)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplDayContext is an interface to support dynamic dispatch.
type ITmplDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplDayContext differentiates from other interfaces.
	IsTmplDayContext()
}

type TmplDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplDayContext() *TmplDayContext {
	var p = new(TmplDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplDay
	return p
}

func (*TmplDayContext) IsTmplDayContext() {}

func NewTmplDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplDayContext {
	var p = new(TmplDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplDay

	return p
}

func (s *TmplDayContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplDayContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplDayContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *TmplDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplDay(s)
	}
}

func (s *TmplDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplDay(s)
	}
}

func (s *TmplDayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplDay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplDay() (localctx ITmplDayContext) {
	localctx = NewTmplDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, LMLParserRULE_tmplDay)

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
		p.SetState(291)
		p.Match(LMLParserT__12)
	}
	{
		p.SetState(292)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(293)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// ITmplIDContext is an interface to support dynamic dispatch.
type ITmplIDContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplIDContext differentiates from other interfaces.
	IsTmplIDContext()
}

type TmplIDContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplIDContext() *TmplIDContext {
	var p = new(TmplIDContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplID
	return p
}

func (*TmplIDContext) IsTmplIDContext() {}

func NewTmplIDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplIDContext {
	var p = new(TmplIDContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplID

	return p
}

func (s *TmplIDContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplIDContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplIDContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplIDContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplIDContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplIDContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplID(s)
	}
}

func (s *TmplIDContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplID(s)
	}
}

func (s *TmplIDContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplID(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplID() (localctx ITmplIDContext) {
	localctx = NewTmplIDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, LMLParserRULE_tmplID)

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
		p.SetState(295)
		p.Match(LMLParserT__13)
	}
	{
		p.SetState(296)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(297)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplMonthContext is an interface to support dynamic dispatch.
type ITmplMonthContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplMonthContext differentiates from other interfaces.
	IsTmplMonthContext()
}

type TmplMonthContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplMonthContext() *TmplMonthContext {
	var p = new(TmplMonthContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplMonth
	return p
}

func (*TmplMonthContext) IsTmplMonthContext() {}

func NewTmplMonthContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplMonthContext {
	var p = new(TmplMonthContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplMonth

	return p
}

func (s *TmplMonthContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplMonthContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplMonthContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *TmplMonthContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplMonthContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplMonthContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplMonth(s)
	}
}

func (s *TmplMonthContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplMonth(s)
	}
}

func (s *TmplMonthContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplMonth(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplMonth() (localctx ITmplMonthContext) {
	localctx = NewTmplMonthContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, LMLParserRULE_tmplMonth)

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
		p.SetState(299)
		p.Match(LMLParserT__14)
	}
	{
		p.SetState(300)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(301)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// ITmplPageHeaderEvenContext is an interface to support dynamic dispatch.
type ITmplPageHeaderEvenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageHeaderEvenContext differentiates from other interfaces.
	IsTmplPageHeaderEvenContext()
}

type TmplPageHeaderEvenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageHeaderEvenContext() *TmplPageHeaderEvenContext {
	var p = new(TmplPageHeaderEvenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageHeaderEven
	return p
}

func (*TmplPageHeaderEvenContext) IsTmplPageHeaderEvenContext() {}

func NewTmplPageHeaderEvenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageHeaderEvenContext {
	var p = new(TmplPageHeaderEvenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageHeaderEven

	return p
}

func (s *TmplPageHeaderEvenContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageHeaderEvenContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageHeaderEvenContext) AllPosition() []IPositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPositionContext)(nil)).Elem())
	var tst = make([]IPositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPositionContext)
		}
	}

	return tst
}

func (s *TmplPageHeaderEvenContext) Position(i int) IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *TmplPageHeaderEvenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageHeaderEvenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageHeaderEvenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageHeaderEven(s)
	}
}

func (s *TmplPageHeaderEvenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageHeaderEven(s)
	}
}

func (s *TmplPageHeaderEvenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageHeaderEven(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageHeaderEven() (localctx ITmplPageHeaderEvenContext) {
	localctx = NewTmplPageHeaderEvenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, LMLParserRULE_tmplPageHeaderEven)
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
		p.SetState(303)
		p.Match(LMLParserT__15)
	}
	{
		p.SetState(304)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(305)
			p.Position()
		}

		p.SetState(308)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmplPageFooterEvenContext is an interface to support dynamic dispatch.
type ITmplPageFooterEvenContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageFooterEvenContext differentiates from other interfaces.
	IsTmplPageFooterEvenContext()
}

type TmplPageFooterEvenContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageFooterEvenContext() *TmplPageFooterEvenContext {
	var p = new(TmplPageFooterEvenContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageFooterEven
	return p
}

func (*TmplPageFooterEvenContext) IsTmplPageFooterEvenContext() {}

func NewTmplPageFooterEvenContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageFooterEvenContext {
	var p = new(TmplPageFooterEvenContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageFooterEven

	return p
}

func (s *TmplPageFooterEvenContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageFooterEvenContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageFooterEvenContext) AllPosition() []IPositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPositionContext)(nil)).Elem())
	var tst = make([]IPositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPositionContext)
		}
	}

	return tst
}

func (s *TmplPageFooterEvenContext) Position(i int) IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *TmplPageFooterEvenContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageFooterEvenContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageFooterEvenContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageFooterEven(s)
	}
}

func (s *TmplPageFooterEvenContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageFooterEven(s)
	}
}

func (s *TmplPageFooterEvenContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageFooterEven(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageFooterEven() (localctx ITmplPageFooterEvenContext) {
	localctx = NewTmplPageFooterEvenContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, LMLParserRULE_tmplPageFooterEven)
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
		p.SetState(310)
		p.Match(LMLParserT__16)
	}
	{
		p.SetState(311)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(312)
			p.Position()
		}

		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmplPageHeaderOddContext is an interface to support dynamic dispatch.
type ITmplPageHeaderOddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageHeaderOddContext differentiates from other interfaces.
	IsTmplPageHeaderOddContext()
}

type TmplPageHeaderOddContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageHeaderOddContext() *TmplPageHeaderOddContext {
	var p = new(TmplPageHeaderOddContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageHeaderOdd
	return p
}

func (*TmplPageHeaderOddContext) IsTmplPageHeaderOddContext() {}

func NewTmplPageHeaderOddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageHeaderOddContext {
	var p = new(TmplPageHeaderOddContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageHeaderOdd

	return p
}

func (s *TmplPageHeaderOddContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageHeaderOddContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageHeaderOddContext) AllPosition() []IPositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPositionContext)(nil)).Elem())
	var tst = make([]IPositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPositionContext)
		}
	}

	return tst
}

func (s *TmplPageHeaderOddContext) Position(i int) IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *TmplPageHeaderOddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageHeaderOddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageHeaderOddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageHeaderOdd(s)
	}
}

func (s *TmplPageHeaderOddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageHeaderOdd(s)
	}
}

func (s *TmplPageHeaderOddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageHeaderOdd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageHeaderOdd() (localctx ITmplPageHeaderOddContext) {
	localctx = NewTmplPageHeaderOddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, LMLParserRULE_tmplPageHeaderOdd)
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
		p.SetState(317)
		p.Match(LMLParserT__17)
	}
	{
		p.SetState(318)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(320)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(319)
			p.Position()
		}

		p.SetState(322)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmplPageFooterOddContext is an interface to support dynamic dispatch.
type ITmplPageFooterOddContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageFooterOddContext differentiates from other interfaces.
	IsTmplPageFooterOddContext()
}

type TmplPageFooterOddContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageFooterOddContext() *TmplPageFooterOddContext {
	var p = new(TmplPageFooterOddContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageFooterOdd
	return p
}

func (*TmplPageFooterOddContext) IsTmplPageFooterOddContext() {}

func NewTmplPageFooterOddContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageFooterOddContext {
	var p = new(TmplPageFooterOddContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageFooterOdd

	return p
}

func (s *TmplPageFooterOddContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageFooterOddContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageFooterOddContext) AllPosition() []IPositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPositionContext)(nil)).Elem())
	var tst = make([]IPositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPositionContext)
		}
	}

	return tst
}

func (s *TmplPageFooterOddContext) Position(i int) IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *TmplPageFooterOddContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageFooterOddContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageFooterOddContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageFooterOdd(s)
	}
}

func (s *TmplPageFooterOddContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageFooterOdd(s)
	}
}

func (s *TmplPageFooterOddContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageFooterOdd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageFooterOdd() (localctx ITmplPageFooterOddContext) {
	localctx = NewTmplPageFooterOddContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, LMLParserRULE_tmplPageFooterOdd)
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
		p.SetState(324)
		p.Match(LMLParserT__18)
	}
	{
		p.SetState(325)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(327)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(326)
			p.Position()
		}

		p.SetState(329)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmplPageHeaderContext is an interface to support dynamic dispatch.
type ITmplPageHeaderContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageHeaderContext differentiates from other interfaces.
	IsTmplPageHeaderContext()
}

type TmplPageHeaderContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageHeaderContext() *TmplPageHeaderContext {
	var p = new(TmplPageHeaderContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageHeader
	return p
}

func (*TmplPageHeaderContext) IsTmplPageHeaderContext() {}

func NewTmplPageHeaderContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageHeaderContext {
	var p = new(TmplPageHeaderContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageHeader

	return p
}

func (s *TmplPageHeaderContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageHeaderContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageHeaderContext) AllPosition() []IPositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPositionContext)(nil)).Elem())
	var tst = make([]IPositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPositionContext)
		}
	}

	return tst
}

func (s *TmplPageHeaderContext) Position(i int) IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *TmplPageHeaderContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageHeaderContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageHeaderContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageHeader(s)
	}
}

func (s *TmplPageHeaderContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageHeader(s)
	}
}

func (s *TmplPageHeaderContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageHeader(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageHeader() (localctx ITmplPageHeaderContext) {
	localctx = NewTmplPageHeaderContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, LMLParserRULE_tmplPageHeader)
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
		p.SetState(331)
		p.Match(LMLParserT__19)
	}
	{
		p.SetState(332)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(333)
			p.Position()
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmplPageFooterContext is an interface to support dynamic dispatch.
type ITmplPageFooterContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageFooterContext differentiates from other interfaces.
	IsTmplPageFooterContext()
}

type TmplPageFooterContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageFooterContext() *TmplPageFooterContext {
	var p = new(TmplPageFooterContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageFooter
	return p
}

func (*TmplPageFooterContext) IsTmplPageFooterContext() {}

func NewTmplPageFooterContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageFooterContext {
	var p = new(TmplPageFooterContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageFooter

	return p
}

func (s *TmplPageFooterContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageFooterContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageFooterContext) AllPosition() []IPositionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IPositionContext)(nil)).Elem())
	var tst = make([]IPositionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IPositionContext)
		}
	}

	return tst
}

func (s *TmplPageFooterContext) Position(i int) IPositionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPositionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IPositionContext)
}

func (s *TmplPageFooterContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageFooterContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageFooterContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageFooter(s)
	}
}

func (s *TmplPageFooterContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageFooter(s)
	}
}

func (s *TmplPageFooterContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageFooter(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageFooter() (localctx ITmplPageFooterContext) {
	localctx = NewTmplPageFooterContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, LMLParserRULE_tmplPageFooter)
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
		p.SetState(338)
		p.Match(LMLParserT__20)
	}
	{
		p.SetState(339)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(341)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(340)
			p.Position()
		}

		p.SetState(343)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ITmplPageNumberContext is an interface to support dynamic dispatch.
type ITmplPageNumberContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplPageNumberContext differentiates from other interfaces.
	IsTmplPageNumberContext()
}

type TmplPageNumberContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplPageNumberContext() *TmplPageNumberContext {
	var p = new(TmplPageNumberContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplPageNumber
	return p
}

func (*TmplPageNumberContext) IsTmplPageNumberContext() {}

func NewTmplPageNumberContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplPageNumberContext {
	var p = new(TmplPageNumberContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplPageNumber

	return p
}

func (s *TmplPageNumberContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplPageNumberContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplPageNumberContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *TmplPageNumberContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplPageNumberContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplPageNumberContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplPageNumber(s)
	}
}

func (s *TmplPageNumberContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplPageNumber(s)
	}
}

func (s *TmplPageNumberContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplPageNumber(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplPageNumber() (localctx ITmplPageNumberContext) {
	localctx = NewTmplPageNumberContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, LMLParserRULE_tmplPageNumber)

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
		p.SetState(345)
		p.Match(LMLParserT__21)
	}
	{
		p.SetState(346)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(347)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// ITmplStatusContext is an interface to support dynamic dispatch.
type ITmplStatusContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplStatusContext differentiates from other interfaces.
	IsTmplStatusContext()
}

type TmplStatusContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplStatusContext() *TmplStatusContext {
	var p = new(TmplStatusContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplStatus
	return p
}

func (*TmplStatusContext) IsTmplStatusContext() {}

func NewTmplStatusContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplStatusContext {
	var p = new(TmplStatusContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplStatus

	return p
}

func (s *TmplStatusContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplStatusContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplStatusContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplStatusContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplStatusContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplStatusContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplStatus(s)
	}
}

func (s *TmplStatusContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplStatus(s)
	}
}

func (s *TmplStatusContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplStatus(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplStatus() (localctx ITmplStatusContext) {
	localctx = NewTmplStatusContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, LMLParserRULE_tmplStatus)

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
		p.SetState(349)
		p.Match(LMLParserT__22)
	}
	{
		p.SetState(350)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(351)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplTitleContext is an interface to support dynamic dispatch.
type ITmplTitleContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplTitleContext differentiates from other interfaces.
	IsTmplTitleContext()
}

type TmplTitleContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplTitleContext() *TmplTitleContext {
	var p = new(TmplTitleContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplTitle
	return p
}

func (*TmplTitleContext) IsTmplTitleContext() {}

func NewTmplTitleContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplTitleContext {
	var p = new(TmplTitleContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplTitle

	return p
}

func (s *TmplTitleContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplTitleContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplTitleContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplTitleContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplTitleContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplTitleContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplTitle(s)
	}
}

func (s *TmplTitleContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplTitle(s)
	}
}

func (s *TmplTitleContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplTitle(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplTitle() (localctx ITmplTitleContext) {
	localctx = NewTmplTitleContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, LMLParserRULE_tmplTitle)

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
		p.SetState(353)
		p.Match(LMLParserT__23)
	}
	{
		p.SetState(354)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(355)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplTypeContext is an interface to support dynamic dispatch.
type ITmplTypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplTypeContext differentiates from other interfaces.
	IsTmplTypeContext()
}

type TmplTypeContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplTypeContext() *TmplTypeContext {
	var p = new(TmplTypeContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplType
	return p
}

func (*TmplTypeContext) IsTmplTypeContext() {}

func NewTmplTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplTypeContext {
	var p = new(TmplTypeContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplType

	return p
}

func (s *TmplTypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplTypeContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplTypeContext) STRING() antlr.TerminalNode {
	return s.GetToken(LMLParserSTRING, 0)
}

func (s *TmplTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplTypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplType(s)
	}
}

func (s *TmplTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplType(s)
	}
}

func (s *TmplTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplType() (localctx ITmplTypeContext) {
	localctx = NewTmplTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, LMLParserRULE_tmplType)

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
		p.SetState(357)
		p.Match(LMLParserT__24)
	}
	{
		p.SetState(358)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(359)
		p.Match(LMLParserSTRING)
	}

	return localctx
}

// ITmplYearContext is an interface to support dynamic dispatch.
type ITmplYearContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsTmplYearContext differentiates from other interfaces.
	IsTmplYearContext()
}

type TmplYearContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTmplYearContext() *TmplYearContext {
	var p = new(TmplYearContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_tmplYear
	return p
}

func (*TmplYearContext) IsTmplYearContext() {}

func NewTmplYearContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TmplYearContext {
	var p = new(TmplYearContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_tmplYear

	return p
}

func (s *TmplYearContext) GetParser() antlr.Parser { return s.parser }

func (s *TmplYearContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *TmplYearContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *TmplYearContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TmplYearContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TmplYearContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterTmplYear(s)
	}
}

func (s *TmplYearContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitTmplYear(s)
	}
}

func (s *TmplYearContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitTmplYear(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) TmplYear() (localctx ITmplYearContext) {
	localctx = NewTmplYearContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, LMLParserRULE_tmplYear)

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
		p.SetState(361)
		p.Match(LMLParserT__25)
	}
	{
		p.SetState(362)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(363)
		p.Match(LMLParserINTEGER)
	}

	return localctx
}

// IBlockContext is an interface to support dynamic dispatch.
type IBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBlockContext differentiates from other interfaces.
	IsBlockContext()
}

type BlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlockContext() *BlockContext {
	var p = new(BlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_block
	return p
}

func (*BlockContext) IsBlockContext() {}

func NewBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlockContext {
	var p = new(BlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_block

	return p
}

func (s *BlockContext) GetParser() antlr.Parser { return s.parser }

func (s *BlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserLBRACE, 0)
}

func (s *BlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserRBRACE, 0)
}

func (s *BlockContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *BlockContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *BlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterBlock(s)
	}
}

func (s *BlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitBlock(s)
	}
}

func (s *BlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Block() (localctx IBlockContext) {
	localctx = NewBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, LMLParserRULE_block)
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
		p.SetState(365)
		p.Match(LMLParserLBRACE)
	}
	p.SetState(369)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
		{
			p.SetState(366)
			p.Statement()
		}

		p.SetState(371)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(372)
		p.Match(LMLParserRBRACE)
	}

	return localctx
}

// ISwitchBlockContext is an interface to support dynamic dispatch.
type ISwitchBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockContext differentiates from other interfaces.
	IsSwitchBlockContext()
}

type SwitchBlockContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockContext() *SwitchBlockContext {
	var p = new(SwitchBlockContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_switchBlock
	return p
}

func (*SwitchBlockContext) IsSwitchBlockContext() {}

func NewSwitchBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockContext {
	var p = new(SwitchBlockContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_switchBlock

	return p
}

func (s *SwitchBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserLBRACE, 0)
}

func (s *SwitchBlockContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserRBRACE, 0)
}

func (s *SwitchBlockContext) AllSwitchBlockStatementGroup() []ISwitchBlockStatementGroupContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem())
	var tst = make([]ISwitchBlockStatementGroupContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchBlockStatementGroupContext)
		}
	}

	return tst
}

func (s *SwitchBlockContext) SwitchBlockStatementGroup(i int) ISwitchBlockStatementGroupContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchBlockStatementGroupContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchBlockStatementGroupContext)
}

func (s *SwitchBlockContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *SwitchBlockContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *SwitchBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterSwitchBlock(s)
	}
}

func (s *SwitchBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitSwitchBlock(s)
	}
}

func (s *SwitchBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitSwitchBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) SwitchBlock() (localctx ISwitchBlockContext) {
	localctx = NewSwitchBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, LMLParserRULE_switchBlock)
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
	{
		p.SetState(374)
		p.Match(LMLParserLBRACE)
	}
	p.SetState(378)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(375)
				p.SwitchBlockStatementGroup()
			}

		}
		p.SetState(380)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext())
	}
	p.SetState(384)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__26 || _la == LMLParserCASE {
		{
			p.SetState(381)
			p.SwitchLabel()
		}

		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(387)
		p.Match(LMLParserRBRACE)
	}

	return localctx
}

// ISwitchBlockStatementGroupContext is an interface to support dynamic dispatch.
type ISwitchBlockStatementGroupContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchBlockStatementGroupContext differentiates from other interfaces.
	IsSwitchBlockStatementGroupContext()
}

type SwitchBlockStatementGroupContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchBlockStatementGroupContext() *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_switchBlockStatementGroup
	return p
}

func (*SwitchBlockStatementGroupContext) IsSwitchBlockStatementGroupContext() {}

func NewSwitchBlockStatementGroupContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchBlockStatementGroupContext {
	var p = new(SwitchBlockStatementGroupContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_switchBlockStatementGroup

	return p
}

func (s *SwitchBlockStatementGroupContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchBlockStatementGroupContext) AllSwitchLabel() []ISwitchLabelContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem())
	var tst = make([]ISwitchLabelContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISwitchLabelContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) SwitchLabel(i int) ISwitchLabelContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISwitchLabelContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISwitchLabelContext)
}

func (s *SwitchBlockStatementGroupContext) AllStatement() []IStatementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IStatementContext)(nil)).Elem())
	var tst = make([]IStatementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IStatementContext)
		}
	}

	return tst
}

func (s *SwitchBlockStatementGroupContext) Statement(i int) IStatementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStatementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IStatementContext)
}

func (s *SwitchBlockStatementGroupContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserLBRACE, 0)
}

func (s *SwitchBlockStatementGroupContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(LMLParserRBRACE, 0)
}

func (s *SwitchBlockStatementGroupContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchBlockStatementGroupContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchBlockStatementGroupContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitSwitchBlockStatementGroup(s)
	}
}

func (s *SwitchBlockStatementGroupContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitSwitchBlockStatementGroup(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) SwitchBlockStatementGroup() (localctx ISwitchBlockStatementGroupContext) {
	localctx = NewSwitchBlockStatementGroupContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, LMLParserRULE_switchBlockStatementGroup)
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

	p.SetState(414)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(389)
					p.SwitchLabel()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(392)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
		}
		p.SetState(397)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
			{
				p.SetState(394)
				p.Statement()
			}

			p.SetState(399)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(401)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LMLParserT__26 || _la == LMLParserCASE {
			{
				p.SetState(400)
				p.SwitchLabel()
			}

			p.SetState(403)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(405)
			p.Match(LMLParserLBRACE)
		}
		p.SetState(409)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
			{
				p.SetState(406)
				p.Statement()
			}

			p.SetState(411)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(412)
			p.Match(LMLParserRBRACE)
		}

	}

	return localctx
}

// ISwitchLabelContext is an interface to support dynamic dispatch.
type ISwitchLabelContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSwitchLabelContext differentiates from other interfaces.
	IsSwitchLabelContext()
}

type SwitchLabelContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySwitchLabelContext() *SwitchLabelContext {
	var p = new(SwitchLabelContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_switchLabel
	return p
}

func (*SwitchLabelContext) IsSwitchLabelContext() {}

func NewSwitchLabelContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SwitchLabelContext {
	var p = new(SwitchLabelContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_switchLabel

	return p
}

func (s *SwitchLabelContext) GetParser() antlr.Parser { return s.parser }

func (s *SwitchLabelContext) CASE() antlr.TerminalNode {
	return s.GetToken(LMLParserCASE, 0)
}

func (s *SwitchLabelContext) IntegerExpression() IIntegerExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerExpressionContext)
}

func (s *SwitchLabelContext) COLON() antlr.TerminalNode {
	return s.GetToken(LMLParserCOLON, 0)
}

func (s *SwitchLabelContext) DowExpression() IDowExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDowExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDowExpressionContext)
}

func (s *SwitchLabelContext) MonthName() IMonthNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthNameContext)
}

func (s *SwitchLabelContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SwitchLabelContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SwitchLabelContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitSwitchLabel(s)
	}
}

func (s *SwitchLabelContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitSwitchLabel(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) SwitchLabel() (localctx ISwitchLabelContext) {
	localctx = NewSwitchLabelContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, LMLParserRULE_switchLabel)

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

	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(416)
			p.Match(LMLParserCASE)
		}
		{
			p.SetState(417)
			p.IntegerExpression()
		}
		{
			p.SetState(418)
			p.Match(LMLParserCOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(420)
			p.Match(LMLParserCASE)
		}
		{
			p.SetState(421)
			p.DowExpression()
		}
		{
			p.SetState(422)
			p.Match(LMLParserCOLON)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(424)
			p.Match(LMLParserCASE)
		}
		{
			p.SetState(425)
			p.MonthName()
		}
		{
			p.SetState(426)
			p.IntegerExpression()
		}
		{
			p.SetState(427)
			p.Match(LMLParserCOLON)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(429)
			p.Match(LMLParserT__26)
		}
		{
			p.SetState(430)
			p.Match(LMLParserCOLON)
		}

	}

	return localctx
}

// IIntegerExpressionContext is an interface to support dynamic dispatch.
type IIntegerExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerExpressionContext differentiates from other interfaces.
	IsIntegerExpressionContext()
}

type IntegerExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerExpressionContext() *IntegerExpressionContext {
	var p = new(IntegerExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_integerExpression
	return p
}

func (*IntegerExpressionContext) IsIntegerExpressionContext() {}

func NewIntegerExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerExpressionContext {
	var p = new(IntegerExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_integerExpression

	return p
}

func (s *IntegerExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerExpressionContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(LMLParserINTEGER)
}

func (s *IntegerExpressionContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, i)
}

func (s *IntegerExpressionContext) THRU() antlr.TerminalNode {
	return s.GetToken(LMLParserTHRU, 0)
}

func (s *IntegerExpressionContext) IntegerList() IIntegerListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIntegerListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIntegerListContext)
}

func (s *IntegerExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterIntegerExpression(s)
	}
}

func (s *IntegerExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitIntegerExpression(s)
	}
}

func (s *IntegerExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitIntegerExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) IntegerExpression() (localctx IIntegerExpressionContext) {
	localctx = NewIntegerExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, LMLParserRULE_integerExpression)

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

	p.SetState(438)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 40, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(433)
			p.Match(LMLParserINTEGER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(434)
			p.Match(LMLParserINTEGER)
		}
		{
			p.SetState(435)
			p.Match(LMLParserTHRU)
		}
		{
			p.SetState(436)
			p.Match(LMLParserINTEGER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(437)
			p.IntegerList()
		}

	}

	return localctx
}

// IIntegerListContext is an interface to support dynamic dispatch.
type IIntegerListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsIntegerListContext differentiates from other interfaces.
	IsIntegerListContext()
}

type IntegerListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIntegerListContext() *IntegerListContext {
	var p = new(IntegerListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_integerList
	return p
}

func (*IntegerListContext) IsIntegerListContext() {}

func NewIntegerListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IntegerListContext {
	var p = new(IntegerListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_integerList

	return p
}

func (s *IntegerListContext) GetParser() antlr.Parser { return s.parser }

func (s *IntegerListContext) AllINTEGER() []antlr.TerminalNode {
	return s.GetTokens(LMLParserINTEGER)
}

func (s *IntegerListContext) INTEGER(i int) antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, i)
}

func (s *IntegerListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntegerListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IntegerListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterIntegerList(s)
	}
}

func (s *IntegerListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitIntegerList(s)
	}
}

func (s *IntegerListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitIntegerList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) IntegerList() (localctx IIntegerListContext) {
	localctx = NewIntegerListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, LMLParserRULE_integerList)
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
		p.SetState(440)
		p.Match(LMLParserINTEGER)
	}
	p.SetState(445)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__27 {
		{
			p.SetState(441)
			p.Match(LMLParserT__27)
		}
		{
			p.SetState(442)
			p.Match(LMLParserINTEGER)
		}

		p.SetState(447)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IDowExpressionContext is an interface to support dynamic dispatch.
type IDowExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDowExpressionContext differentiates from other interfaces.
	IsDowExpressionContext()
}

type DowExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDowExpressionContext() *DowExpressionContext {
	var p = new(DowExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_dowExpression
	return p
}

func (*DowExpressionContext) IsDowExpressionContext() {}

func NewDowExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DowExpressionContext {
	var p = new(DowExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_dowExpression

	return p
}

func (s *DowExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *DowExpressionContext) AllDowName() []IDowNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDowNameContext)(nil)).Elem())
	var tst = make([]IDowNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDowNameContext)
		}
	}

	return tst
}

func (s *DowExpressionContext) DowName(i int) IDowNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDowNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDowNameContext)
}

func (s *DowExpressionContext) THRU() antlr.TerminalNode {
	return s.GetToken(LMLParserTHRU, 0)
}

func (s *DowExpressionContext) DowList() IDowListContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDowListContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDowListContext)
}

func (s *DowExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DowExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DowExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterDowExpression(s)
	}
}

func (s *DowExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitDowExpression(s)
	}
}

func (s *DowExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitDowExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) DowExpression() (localctx IDowExpressionContext) {
	localctx = NewDowExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, LMLParserRULE_dowExpression)

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

	p.SetState(454)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 42, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(448)
			p.DowName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(449)
			p.DowName()
		}
		{
			p.SetState(450)
			p.Match(LMLParserTHRU)
		}
		{
			p.SetState(451)
			p.DowName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(453)
			p.DowList()
		}

	}

	return localctx
}

// IDowListContext is an interface to support dynamic dispatch.
type IDowListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDowListContext differentiates from other interfaces.
	IsDowListContext()
}

type DowListContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDowListContext() *DowListContext {
	var p = new(DowListContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_dowList
	return p
}

func (*DowListContext) IsDowListContext() {}

func NewDowListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DowListContext {
	var p = new(DowListContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_dowList

	return p
}

func (s *DowListContext) GetParser() antlr.Parser { return s.parser }

func (s *DowListContext) AllDowName() []IDowNameContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IDowNameContext)(nil)).Elem())
	var tst = make([]IDowNameContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IDowNameContext)
		}
	}

	return tst
}

func (s *DowListContext) DowName(i int) IDowNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDowNameContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IDowNameContext)
}

func (s *DowListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DowListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DowListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterDowList(s)
	}
}

func (s *DowListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitDowList(s)
	}
}

func (s *DowListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitDowList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) DowList() (localctx IDowListContext) {
	localctx = NewDowListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, LMLParserRULE_dowList)
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
		p.SetState(456)
		p.DowName()
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__27 {
		{
			p.SetState(457)
			p.Match(LMLParserT__27)
		}
		{
			p.SetState(458)
			p.DowName()
		}

		p.SetState(463)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ILdpIntContext is an interface to support dynamic dispatch.
type ILdpIntContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLdpIntContext differentiates from other interfaces.
	IsLdpIntContext()
}

type LdpIntContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLdpIntContext() *LdpIntContext {
	var p = new(LdpIntContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_ldpInt
	return p
}

func (*LdpIntContext) IsLdpIntContext() {}

func NewLdpIntContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LdpIntContext {
	var p = new(LdpIntContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_ldpInt

	return p
}

func (s *LdpIntContext) GetParser() antlr.Parser { return s.parser }

func (s *LdpIntContext) MODE_OF_WEEK() antlr.TerminalNode {
	return s.GetToken(LMLParserMODE_OF_WEEK, 0)
}

func (s *LdpIntContext) MOVABLE_CYCLE_DAY() antlr.TerminalNode {
	return s.GetToken(LMLParserMOVABLE_CYCLE_DAY, 0)
}

func (s *LdpIntContext) LUKAN_CYCLE_DAY() antlr.TerminalNode {
	return s.GetToken(LMLParserLUKAN_CYCLE_DAY, 0)
}

func (s *LdpIntContext) SUNDAY_AFTER_ELEVATION_OF_CROSS() antlr.TerminalNode {
	return s.GetToken(LMLParserSUNDAY_AFTER_ELEVATION_OF_CROSS, 0)
}

func (s *LdpIntContext) SUNDAYS_BEFORE_TRIODION() antlr.TerminalNode {
	return s.GetToken(LMLParserSUNDAYS_BEFORE_TRIODION, 0)
}

func (s *LdpIntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LdpIntContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LdpIntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterLdpInt(s)
	}
}

func (s *LdpIntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitLdpInt(s)
	}
}

func (s *LdpIntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitLdpInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) LdpInt() (localctx ILdpIntContext) {
	localctx = NewLdpIntContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, LMLParserRULE_ldpInt)
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
		p.SetState(464)
		_la = p.GetTokenStream().LA(1)

		if !(((_la-49)&-(0x1f+1)) == 0 && ((1<<uint((_la-49)))&((1<<(LMLParserMODE_OF_WEEK-49))|(1<<(LMLParserMOVABLE_CYCLE_DAY-49))|(1<<(LMLParserLUKAN_CYCLE_DAY-49))|(1<<(LMLParserSUNDAY_AFTER_ELEVATION_OF_CROSS-49))|(1<<(LMLParserSUNDAYS_BEFORE_TRIODION-49)))) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

	return localctx
}

// IMonthDayContext is an interface to support dynamic dispatch.
type IMonthDayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMonthDayContext differentiates from other interfaces.
	IsMonthDayContext()
}

type MonthDayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMonthDayContext() *MonthDayContext {
	var p = new(MonthDayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = LMLParserRULE_monthDay
	return p
}

func (*MonthDayContext) IsMonthDayContext() {}

func NewMonthDayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MonthDayContext {
	var p = new(MonthDayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_monthDay

	return p
}

func (s *MonthDayContext) GetParser() antlr.Parser { return s.parser }

func (s *MonthDayContext) MonthName() IMonthNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthNameContext)
}

func (s *MonthDayContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *MonthDayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MonthDayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MonthDayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterMonthDay(s)
	}
}

func (s *MonthDayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitMonthDay(s)
	}
}

func (s *MonthDayContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitMonthDay(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) MonthDay() (localctx IMonthDayContext) {
	localctx = NewMonthDayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, LMLParserRULE_monthDay)

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
		p.SetState(466)
		p.MonthName()
	}
	{
		p.SetState(467)
		p.Match(LMLParserINTEGER)
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
	p.RuleIndex = LMLParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = LMLParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) LdpInt() ILdpIntContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILdpIntContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILdpIntContext)
}

func (s *ExpressionContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(LMLParserINTEGER, 0)
}

func (s *ExpressionContext) LT() antlr.TerminalNode {
	return s.GetToken(LMLParserLT, 0)
}

func (s *ExpressionContext) ASSIGNMENT() antlr.TerminalNode {
	return s.GetToken(LMLParserASSIGNMENT, 0)
}

func (s *ExpressionContext) GT() antlr.TerminalNode {
	return s.GetToken(LMLParserGT, 0)
}

func (s *ExpressionContext) EQUALITY() antlr.TerminalNode {
	return s.GetToken(LMLParserEQUALITY, 0)
}

func (s *ExpressionContext) INEQUALITY() antlr.TerminalNode {
	return s.GetToken(LMLParserINEQUALITY, 0)
}

func (s *ExpressionContext) DATE() antlr.TerminalNode {
	return s.GetToken(LMLParserDATE, 0)
}

func (s *ExpressionContext) MonthDay() IMonthDayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMonthDayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMonthDayContext)
}

func (s *ExpressionContext) EXISTS() antlr.TerminalNode {
	return s.GetToken(LMLParserEXISTS, 0)
}

func (s *ExpressionContext) Rid() IRidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRidContext)
}

func (s *ExpressionContext) NAME_OF_DAY() antlr.TerminalNode {
	return s.GetToken(LMLParserNAME_OF_DAY, 0)
}

func (s *ExpressionContext) DowName() IDowNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDowNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDowNameContext)
}

func (s *ExpressionContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(LMLParserLPAREN, 0)
}

func (s *ExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *ExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ExpressionContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(LMLParserRPAREN, 0)
}

func (s *ExpressionContext) AND() antlr.TerminalNode {
	return s.GetToken(LMLParserAND, 0)
}

func (s *ExpressionContext) OR() antlr.TerminalNode {
	return s.GetToken(LMLParserOR, 0)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(LMLListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case LMLVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *LMLParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *LMLParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 96
	p.EnterRecursionRule(localctx, 96, LMLParserRULE_expression, _p)
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
	p.SetState(507)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(470)
			p.LdpInt()
		}
		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(471)
				p.Match(LMLParserLT)
			}
			{
				p.SetState(472)
				p.Match(LMLParserASSIGNMENT)
			}

		case 2:
			{
				p.SetState(473)
				p.Match(LMLParserGT)
			}
			{
				p.SetState(474)
				p.Match(LMLParserASSIGNMENT)
			}

		case 3:
			{
				p.SetState(475)
				p.Match(LMLParserGT)
			}

		case 4:
			{
				p.SetState(476)
				p.Match(LMLParserLT)
			}

		}
		{
			p.SetState(479)
			p.Match(LMLParserINTEGER)
		}

	case 2:
		{
			p.SetState(481)
			p.LdpInt()
		}
		{
			p.SetState(482)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LMLParserEQUALITY || _la == LMLParserINEQUALITY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(483)
			p.Match(LMLParserINTEGER)
		}

	case 3:
		{
			p.SetState(485)
			p.Match(LMLParserDATE)
		}
		{
			p.SetState(486)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LMLParserEQUALITY || _la == LMLParserINEQUALITY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(487)
			p.MonthDay()
		}

	case 4:
		{
			p.SetState(488)
			p.Match(LMLParserDATE)
		}
		p.SetState(495)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(489)
				p.Match(LMLParserLT)
			}
			{
				p.SetState(490)
				p.Match(LMLParserASSIGNMENT)
			}

		case 2:
			{
				p.SetState(491)
				p.Match(LMLParserGT)
			}
			{
				p.SetState(492)
				p.Match(LMLParserASSIGNMENT)
			}

		case 3:
			{
				p.SetState(493)
				p.Match(LMLParserGT)
			}

		case 4:
			{
				p.SetState(494)
				p.Match(LMLParserLT)
			}

		}
		{
			p.SetState(497)
			p.MonthDay()
		}

	case 5:
		{
			p.SetState(498)
			p.Match(LMLParserEXISTS)
		}
		{
			p.SetState(499)
			p.Rid()
		}

	case 6:
		{
			p.SetState(500)
			p.Match(LMLParserNAME_OF_DAY)
		}
		{
			p.SetState(501)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LMLParserEQUALITY || _la == LMLParserINEQUALITY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(502)
			p.DowName()
		}

	case 7:
		{
			p.SetState(503)
			p.Match(LMLParserLPAREN)
		}
		{
			p.SetState(504)
			p.expression(0)
		}
		{
			p.SetState(505)
			p.Match(LMLParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(517)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(515)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LMLParserRULE_expression)
				p.SetState(509)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(510)
					p.Match(LMLParserAND)
				}
				{
					p.SetState(511)
					p.expression(4)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LMLParserRULE_expression)
				p.SetState(512)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(513)
					p.Match(LMLParserOR)
				}
				{
					p.SetState(514)
					p.expression(3)
				}

			}

		}
		p.SetState(519)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 48, p.GetParserRuleContext())
	}

	return localctx
}

func (p *LMLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 48:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *LMLParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
