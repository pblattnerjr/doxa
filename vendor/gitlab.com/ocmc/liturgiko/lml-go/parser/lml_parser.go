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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 92, 512,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4, 39, 9,
	39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44, 9, 44,
	4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9, 49, 3,
	2, 3, 2, 3, 2, 3, 2, 3, 2, 6, 2, 104, 10, 2, 13, 2, 14, 2, 105, 3, 2, 3,
	2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 5, 3, 125, 10, 3, 3, 4, 7, 4, 128, 10, 4, 12, 4, 14,
	4, 131, 11, 4, 3, 4, 3, 4, 7, 4, 135, 10, 4, 12, 4, 14, 4, 138, 11, 4,
	3, 4, 5, 4, 141, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5,
	6, 5, 151, 10, 5, 13, 5, 14, 5, 152, 5, 5, 155, 10, 5, 3, 5, 3, 5, 3, 5,
	3, 5, 3, 5, 6, 5, 162, 10, 5, 13, 5, 14, 5, 163, 5, 5, 166, 10, 5, 3, 5,
	3, 5, 3, 5, 3, 5, 5, 5, 172, 10, 5, 3, 5, 3, 5, 5, 5, 176, 10, 5, 3, 6,
	3, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 6, 8, 188, 10, 8,
	13, 8, 14, 8, 189, 3, 8, 5, 8, 193, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9,
	6, 9, 200, 10, 9, 13, 9, 14, 9, 201, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11,
	3, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 5, 12, 222, 10, 12, 3, 13, 3, 13, 6, 13, 226, 10, 13,
	13, 13, 14, 13, 227, 3, 14, 3, 14, 3, 15, 3, 15, 3, 15, 3, 15, 5, 15, 236,
	10, 15, 3, 16, 3, 16, 3, 16, 6, 16, 241, 10, 16, 13, 16, 14, 16, 242, 3,
	16, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 7, 17, 251, 10, 17, 12, 17, 14,
	17, 254, 11, 17, 3, 18, 3, 18, 5, 18, 258, 10, 18, 3, 19, 3, 19, 3, 19,
	3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 5, 21, 269, 10, 21, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3,
	27, 3, 27, 3, 28, 3, 28, 3, 28, 6, 28, 298, 10, 28, 13, 28, 14, 28, 299,
	3, 29, 3, 29, 3, 29, 6, 29, 305, 10, 29, 13, 29, 14, 29, 306, 3, 30, 3,
	30, 3, 30, 6, 30, 312, 10, 30, 13, 30, 14, 30, 313, 3, 31, 3, 31, 3, 31,
	6, 31, 319, 10, 31, 13, 31, 14, 31, 320, 3, 32, 3, 32, 3, 32, 6, 32, 326,
	10, 32, 13, 32, 14, 32, 327, 3, 33, 3, 33, 3, 33, 6, 33, 333, 10, 33, 13,
	33, 14, 33, 334, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35, 3, 35, 3, 35,
	3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3,
	38, 3, 38, 3, 39, 3, 39, 7, 39, 359, 10, 39, 12, 39, 14, 39, 362, 11, 39,
	3, 39, 3, 39, 3, 40, 3, 40, 7, 40, 368, 10, 40, 12, 40, 14, 40, 371, 11,
	40, 3, 40, 7, 40, 374, 10, 40, 12, 40, 14, 40, 377, 11, 40, 3, 40, 3, 40,
	3, 41, 6, 41, 382, 10, 41, 13, 41, 14, 41, 383, 3, 41, 7, 41, 387, 10,
	41, 12, 41, 14, 41, 390, 11, 41, 3, 41, 6, 41, 393, 10, 41, 13, 41, 14,
	41, 394, 3, 41, 3, 41, 7, 41, 399, 10, 41, 12, 41, 14, 41, 402, 11, 41,
	3, 41, 3, 41, 5, 41, 406, 10, 41, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3,
	42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 3, 42, 5, 42,
	423, 10, 42, 3, 43, 3, 43, 3, 43, 3, 43, 3, 43, 5, 43, 430, 10, 43, 3,
	44, 3, 44, 3, 44, 7, 44, 435, 10, 44, 12, 44, 14, 44, 438, 11, 44, 3, 45,
	3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 5, 45, 446, 10, 45, 3, 46, 3, 46, 3,
	46, 7, 46, 451, 10, 46, 12, 46, 14, 46, 454, 11, 46, 3, 47, 3, 47, 3, 48,
	3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 5,
	49, 469, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 5, 49, 487, 10,
	49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	5, 49, 499, 10, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 7, 49, 507,
	10, 49, 12, 49, 14, 49, 510, 11, 49, 3, 49, 2, 3, 96, 50, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82,
	84, 86, 88, 90, 92, 94, 96, 2, 6, 3, 2, 71, 77, 3, 2, 78, 80, 4, 2, 51,
	52, 67, 69, 3, 2, 31, 32, 2, 557, 2, 98, 3, 2, 2, 2, 4, 124, 3, 2, 2, 2,
	6, 140, 3, 2, 2, 2, 8, 175, 3, 2, 2, 2, 10, 177, 3, 2, 2, 2, 12, 179, 3,
	2, 2, 2, 14, 182, 3, 2, 2, 2, 16, 194, 3, 2, 2, 2, 18, 203, 3, 2, 2, 2,
	20, 206, 3, 2, 2, 2, 22, 221, 3, 2, 2, 2, 24, 223, 3, 2, 2, 2, 26, 229,
	3, 2, 2, 2, 28, 235, 3, 2, 2, 2, 30, 237, 3, 2, 2, 2, 32, 247, 3, 2, 2,
	2, 34, 257, 3, 2, 2, 2, 36, 259, 3, 2, 2, 2, 38, 262, 3, 2, 2, 2, 40, 265,
	3, 2, 2, 2, 42, 270, 3, 2, 2, 2, 44, 274, 3, 2, 2, 2, 46, 278, 3, 2, 2,
	2, 48, 282, 3, 2, 2, 2, 50, 286, 3, 2, 2, 2, 52, 290, 3, 2, 2, 2, 54, 294,
	3, 2, 2, 2, 56, 301, 3, 2, 2, 2, 58, 308, 3, 2, 2, 2, 60, 315, 3, 2, 2,
	2, 62, 322, 3, 2, 2, 2, 64, 329, 3, 2, 2, 2, 66, 336, 3, 2, 2, 2, 68, 340,
	3, 2, 2, 2, 70, 344, 3, 2, 2, 2, 72, 348, 3, 2, 2, 2, 74, 352, 3, 2, 2,
	2, 76, 356, 3, 2, 2, 2, 78, 365, 3, 2, 2, 2, 80, 405, 3, 2, 2, 2, 82, 422,
	3, 2, 2, 2, 84, 429, 3, 2, 2, 2, 86, 431, 3, 2, 2, 2, 88, 445, 3, 2, 2,
	2, 90, 447, 3, 2, 2, 2, 92, 455, 3, 2, 2, 2, 94, 457, 3, 2, 2, 2, 96, 498,
	3, 2, 2, 2, 98, 99, 5, 50, 26, 2, 99, 100, 5, 72, 37, 2, 100, 101, 5, 68,
	35, 2, 101, 103, 5, 6, 4, 2, 102, 104, 5, 8, 5, 2, 103, 102, 3, 2, 2, 2,
	104, 105, 3, 2, 2, 2, 105, 103, 3, 2, 2, 2, 105, 106, 3, 2, 2, 2, 106,
	107, 3, 2, 2, 2, 107, 108, 7, 2, 2, 3, 108, 3, 3, 2, 2, 2, 109, 125, 5,
	48, 25, 2, 110, 125, 5, 52, 27, 2, 111, 125, 5, 74, 38, 2, 112, 125, 5,
	44, 23, 2, 113, 125, 5, 46, 24, 2, 114, 125, 5, 70, 36, 2, 115, 125, 5,
	48, 25, 2, 116, 125, 5, 42, 22, 2, 117, 125, 5, 66, 34, 2, 118, 125, 5,
	54, 28, 2, 119, 125, 5, 58, 30, 2, 120, 125, 5, 56, 29, 2, 121, 125, 5,
	60, 31, 2, 122, 125, 5, 62, 32, 2, 123, 125, 5, 64, 33, 2, 124, 109, 3,
	2, 2, 2, 124, 110, 3, 2, 2, 2, 124, 111, 3, 2, 2, 2, 124, 112, 3, 2, 2,
	2, 124, 113, 3, 2, 2, 2, 124, 114, 3, 2, 2, 2, 124, 115, 3, 2, 2, 2, 124,
	116, 3, 2, 2, 2, 124, 117, 3, 2, 2, 2, 124, 118, 3, 2, 2, 2, 124, 119,
	3, 2, 2, 2, 124, 120, 3, 2, 2, 2, 124, 121, 3, 2, 2, 2, 124, 122, 3, 2,
	2, 2, 124, 123, 3, 2, 2, 2, 125, 5, 3, 2, 2, 2, 126, 128, 5, 4, 3, 2, 127,
	126, 3, 2, 2, 2, 128, 131, 3, 2, 2, 2, 129, 127, 3, 2, 2, 2, 129, 130,
	3, 2, 2, 2, 130, 141, 3, 2, 2, 2, 131, 129, 3, 2, 2, 2, 132, 136, 7, 34,
	2, 2, 133, 135, 5, 4, 3, 2, 134, 133, 3, 2, 2, 2, 135, 138, 3, 2, 2, 2,
	136, 134, 3, 2, 2, 2, 136, 137, 3, 2, 2, 2, 137, 139, 3, 2, 2, 2, 138,
	136, 3, 2, 2, 2, 139, 141, 7, 35, 2, 2, 140, 129, 3, 2, 2, 2, 140, 132,
	3, 2, 2, 2, 141, 7, 3, 2, 2, 2, 142, 176, 5, 18, 10, 2, 143, 176, 5, 14,
	8, 2, 144, 176, 5, 12, 7, 2, 145, 146, 7, 45, 2, 2, 146, 147, 5, 96, 49,
	2, 147, 154, 5, 76, 39, 2, 148, 150, 7, 46, 2, 2, 149, 151, 5, 76, 39,
	2, 150, 149, 3, 2, 2, 2, 151, 152, 3, 2, 2, 2, 152, 150, 3, 2, 2, 2, 152,
	153, 3, 2, 2, 2, 153, 155, 3, 2, 2, 2, 154, 148, 3, 2, 2, 2, 154, 155,
	3, 2, 2, 2, 155, 176, 3, 2, 2, 2, 156, 157, 7, 45, 2, 2, 157, 158, 5, 96,
	49, 2, 158, 165, 5, 76, 39, 2, 159, 161, 7, 3, 2, 2, 160, 162, 5, 76, 39,
	2, 161, 160, 3, 2, 2, 2, 162, 163, 3, 2, 2, 2, 163, 161, 3, 2, 2, 2, 163,
	164, 3, 2, 2, 2, 164, 166, 3, 2, 2, 2, 165, 159, 3, 2, 2, 2, 165, 166,
	3, 2, 2, 2, 166, 176, 3, 2, 2, 2, 167, 171, 7, 39, 2, 2, 168, 172, 5, 92,
	47, 2, 169, 172, 7, 65, 2, 2, 170, 172, 7, 66, 2, 2, 171, 168, 3, 2, 2,
	2, 171, 169, 3, 2, 2, 2, 171, 170, 3, 2, 2, 2, 172, 173, 3, 2, 2, 2, 173,
	176, 5, 78, 40, 2, 174, 176, 5, 76, 39, 2, 175, 142, 3, 2, 2, 2, 175, 143,
	3, 2, 2, 2, 175, 144, 3, 2, 2, 2, 175, 145, 3, 2, 2, 2, 175, 156, 3, 2,
	2, 2, 175, 167, 3, 2, 2, 2, 175, 174, 3, 2, 2, 2, 176, 9, 3, 2, 2, 2, 177,
	178, 9, 2, 2, 2, 178, 11, 3, 2, 2, 2, 179, 180, 7, 4, 2, 2, 180, 181, 7,
	87, 2, 2, 181, 13, 3, 2, 2, 2, 182, 187, 7, 81, 2, 2, 183, 188, 5, 20,
	11, 2, 184, 188, 5, 32, 17, 2, 185, 188, 5, 40, 21, 2, 186, 188, 5, 16,
	9, 2, 187, 183, 3, 2, 2, 2, 187, 184, 3, 2, 2, 2, 187, 185, 3, 2, 2, 2,
	187, 186, 3, 2, 2, 2, 188, 189, 3, 2, 2, 2, 189, 187, 3, 2, 2, 2, 189,
	190, 3, 2, 2, 2, 190, 192, 3, 2, 2, 2, 191, 193, 7, 86, 2, 2, 192, 191,
	3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 15, 3, 2, 2, 2, 194, 199, 7, 82,
	2, 2, 195, 200, 5, 20, 11, 2, 196, 200, 5, 32, 17, 2, 197, 200, 5, 40,
	21, 2, 198, 200, 5, 16, 9, 2, 199, 195, 3, 2, 2, 2, 199, 196, 3, 2, 2,
	2, 199, 197, 3, 2, 2, 2, 199, 198, 3, 2, 2, 2, 200, 201, 3, 2, 2, 2, 201,
	199, 3, 2, 2, 2, 201, 202, 3, 2, 2, 2, 202, 17, 3, 2, 2, 2, 203, 204, 7,
	5, 2, 2, 204, 205, 5, 40, 21, 2, 205, 19, 3, 2, 2, 2, 206, 207, 7, 6, 2,
	2, 207, 208, 7, 87, 2, 2, 208, 21, 3, 2, 2, 2, 209, 222, 7, 53, 2, 2, 210,
	222, 7, 54, 2, 2, 211, 222, 7, 55, 2, 2, 212, 213, 7, 56, 2, 2, 213, 222,
	7, 57, 2, 2, 214, 222, 7, 58, 2, 2, 215, 222, 7, 59, 2, 2, 216, 217, 7,
	60, 2, 2, 217, 222, 7, 61, 2, 2, 218, 222, 7, 62, 2, 2, 219, 222, 7, 63,
	2, 2, 220, 222, 7, 64, 2, 2, 221, 209, 3, 2, 2, 2, 221, 210, 3, 2, 2, 2,
	221, 211, 3, 2, 2, 2, 221, 212, 3, 2, 2, 2, 221, 214, 3, 2, 2, 2, 221,
	215, 3, 2, 2, 2, 221, 216, 3, 2, 2, 2, 221, 218, 3, 2, 2, 2, 221, 219,
	3, 2, 2, 2, 221, 220, 3, 2, 2, 2, 222, 23, 3, 2, 2, 2, 223, 225, 5, 26,
	14, 2, 224, 226, 5, 28, 15, 2, 225, 224, 3, 2, 2, 2, 226, 227, 3, 2, 2,
	2, 227, 225, 3, 2, 2, 2, 227, 228, 3, 2, 2, 2, 228, 25, 3, 2, 2, 2, 229,
	230, 9, 3, 2, 2, 230, 27, 3, 2, 2, 2, 231, 236, 5, 30, 16, 2, 232, 236,
	7, 83, 2, 2, 233, 236, 7, 85, 2, 2, 234, 236, 5, 20, 11, 2, 235, 231, 3,
	2, 2, 2, 235, 232, 3, 2, 2, 2, 235, 233, 3, 2, 2, 2, 235, 234, 3, 2, 2,
	2, 236, 29, 3, 2, 2, 2, 237, 240, 7, 84, 2, 2, 238, 241, 5, 40, 21, 2,
	239, 241, 5, 32, 17, 2, 240, 238, 3, 2, 2, 2, 240, 239, 3, 2, 2, 2, 241,
	242, 3, 2, 2, 2, 242, 240, 3, 2, 2, 2, 242, 243, 3, 2, 2, 2, 243, 244,
	3, 2, 2, 2, 244, 245, 7, 7, 2, 2, 245, 246, 7, 70, 2, 2, 246, 31, 3, 2,
	2, 2, 247, 248, 7, 8, 2, 2, 248, 252, 7, 87, 2, 2, 249, 251, 5, 34, 18,
	2, 250, 249, 3, 2, 2, 2, 251, 254, 3, 2, 2, 2, 252, 250, 3, 2, 2, 2, 252,
	253, 3, 2, 2, 2, 253, 33, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 255, 258, 5,
	36, 19, 2, 256, 258, 5, 38, 20, 2, 257, 255, 3, 2, 2, 2, 257, 256, 3, 2,
	2, 2, 258, 35, 3, 2, 2, 2, 259, 260, 7, 9, 2, 2, 260, 261, 7, 70, 2, 2,
	261, 37, 3, 2, 2, 2, 262, 263, 7, 10, 2, 2, 263, 264, 7, 70, 2, 2, 264,
	39, 3, 2, 2, 2, 265, 266, 7, 11, 2, 2, 266, 268, 7, 87, 2, 2, 267, 269,
	7, 86, 2, 2, 268, 267, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 41, 3, 2,
	2, 2, 270, 271, 7, 12, 2, 2, 271, 272, 7, 33, 2, 2, 272, 273, 7, 87, 2,
	2, 273, 43, 3, 2, 2, 2, 274, 275, 7, 13, 2, 2, 275, 276, 7, 33, 2, 2, 276,
	277, 7, 87, 2, 2, 277, 45, 3, 2, 2, 2, 278, 279, 7, 14, 2, 2, 279, 280,
	7, 33, 2, 2, 280, 281, 7, 87, 2, 2, 281, 47, 3, 2, 2, 2, 282, 283, 7, 15,
	2, 2, 283, 284, 7, 33, 2, 2, 284, 285, 7, 70, 2, 2, 285, 49, 3, 2, 2, 2,
	286, 287, 7, 16, 2, 2, 287, 288, 7, 33, 2, 2, 288, 289, 7, 87, 2, 2, 289,
	51, 3, 2, 2, 2, 290, 291, 7, 17, 2, 2, 291, 292, 7, 33, 2, 2, 292, 293,
	7, 70, 2, 2, 293, 53, 3, 2, 2, 2, 294, 295, 7, 18, 2, 2, 295, 297, 7, 33,
	2, 2, 296, 298, 5, 24, 13, 2, 297, 296, 3, 2, 2, 2, 298, 299, 3, 2, 2,
	2, 299, 297, 3, 2, 2, 2, 299, 300, 3, 2, 2, 2, 300, 55, 3, 2, 2, 2, 301,
	302, 7, 19, 2, 2, 302, 304, 7, 33, 2, 2, 303, 305, 5, 24, 13, 2, 304, 303,
	3, 2, 2, 2, 305, 306, 3, 2, 2, 2, 306, 304, 3, 2, 2, 2, 306, 307, 3, 2,
	2, 2, 307, 57, 3, 2, 2, 2, 308, 309, 7, 20, 2, 2, 309, 311, 7, 33, 2, 2,
	310, 312, 5, 24, 13, 2, 311, 310, 3, 2, 2, 2, 312, 313, 3, 2, 2, 2, 313,
	311, 3, 2, 2, 2, 313, 314, 3, 2, 2, 2, 314, 59, 3, 2, 2, 2, 315, 316, 7,
	21, 2, 2, 316, 318, 7, 33, 2, 2, 317, 319, 5, 24, 13, 2, 318, 317, 3, 2,
	2, 2, 319, 320, 3, 2, 2, 2, 320, 318, 3, 2, 2, 2, 320, 321, 3, 2, 2, 2,
	321, 61, 3, 2, 2, 2, 322, 323, 7, 22, 2, 2, 323, 325, 7, 33, 2, 2, 324,
	326, 5, 24, 13, 2, 325, 324, 3, 2, 2, 2, 326, 327, 3, 2, 2, 2, 327, 325,
	3, 2, 2, 2, 327, 328, 3, 2, 2, 2, 328, 63, 3, 2, 2, 2, 329, 330, 7, 23,
	2, 2, 330, 332, 7, 33, 2, 2, 331, 333, 5, 24, 13, 2, 332, 331, 3, 2, 2,
	2, 333, 334, 3, 2, 2, 2, 334, 332, 3, 2, 2, 2, 334, 335, 3, 2, 2, 2, 335,
	65, 3, 2, 2, 2, 336, 337, 7, 24, 2, 2, 337, 338, 7, 33, 2, 2, 338, 339,
	7, 70, 2, 2, 339, 67, 3, 2, 2, 2, 340, 341, 7, 25, 2, 2, 341, 342, 7, 33,
	2, 2, 342, 343, 7, 87, 2, 2, 343, 69, 3, 2, 2, 2, 344, 345, 7, 26, 2, 2,
	345, 346, 7, 33, 2, 2, 346, 347, 7, 87, 2, 2, 347, 71, 3, 2, 2, 2, 348,
	349, 7, 27, 2, 2, 349, 350, 7, 33, 2, 2, 350, 351, 7, 87, 2, 2, 351, 73,
	3, 2, 2, 2, 352, 353, 7, 28, 2, 2, 353, 354, 7, 33, 2, 2, 354, 355, 7,
	70, 2, 2, 355, 75, 3, 2, 2, 2, 356, 360, 7, 34, 2, 2, 357, 359, 5, 8, 5,
	2, 358, 357, 3, 2, 2, 2, 359, 362, 3, 2, 2, 2, 360, 358, 3, 2, 2, 2, 360,
	361, 3, 2, 2, 2, 361, 363, 3, 2, 2, 2, 362, 360, 3, 2, 2, 2, 363, 364,
	7, 35, 2, 2, 364, 77, 3, 2, 2, 2, 365, 369, 7, 34, 2, 2, 366, 368, 5, 80,
	41, 2, 367, 366, 3, 2, 2, 2, 368, 371, 3, 2, 2, 2, 369, 367, 3, 2, 2, 2,
	369, 370, 3, 2, 2, 2, 370, 375, 3, 2, 2, 2, 371, 369, 3, 2, 2, 2, 372,
	374, 5, 82, 42, 2, 373, 372, 3, 2, 2, 2, 374, 377, 3, 2, 2, 2, 375, 373,
	3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 378, 3, 2, 2, 2, 377, 375, 3, 2,
	2, 2, 378, 379, 7, 35, 2, 2, 379, 79, 3, 2, 2, 2, 380, 382, 5, 82, 42,
	2, 381, 380, 3, 2, 2, 2, 382, 383, 3, 2, 2, 2, 383, 381, 3, 2, 2, 2, 383,
	384, 3, 2, 2, 2, 384, 388, 3, 2, 2, 2, 385, 387, 5, 8, 5, 2, 386, 385,
	3, 2, 2, 2, 387, 390, 3, 2, 2, 2, 388, 386, 3, 2, 2, 2, 388, 389, 3, 2,
	2, 2, 389, 406, 3, 2, 2, 2, 390, 388, 3, 2, 2, 2, 391, 393, 5, 82, 42,
	2, 392, 391, 3, 2, 2, 2, 393, 394, 3, 2, 2, 2, 394, 392, 3, 2, 2, 2, 394,
	395, 3, 2, 2, 2, 395, 396, 3, 2, 2, 2, 396, 400, 7, 34, 2, 2, 397, 399,
	5, 8, 5, 2, 398, 397, 3, 2, 2, 2, 399, 402, 3, 2, 2, 2, 400, 398, 3, 2,
	2, 2, 400, 401, 3, 2, 2, 2, 401, 403, 3, 2, 2, 2, 402, 400, 3, 2, 2, 2,
	403, 404, 7, 35, 2, 2, 404, 406, 3, 2, 2, 2, 405, 381, 3, 2, 2, 2, 405,
	392, 3, 2, 2, 2, 406, 81, 3, 2, 2, 2, 407, 408, 7, 47, 2, 2, 408, 409,
	5, 84, 43, 2, 409, 410, 7, 38, 2, 2, 410, 423, 3, 2, 2, 2, 411, 412, 7,
	47, 2, 2, 412, 413, 5, 88, 45, 2, 413, 414, 7, 38, 2, 2, 414, 423, 3, 2,
	2, 2, 415, 416, 7, 47, 2, 2, 416, 417, 5, 22, 12, 2, 417, 418, 5, 84, 43,
	2, 418, 419, 7, 38, 2, 2, 419, 423, 3, 2, 2, 2, 420, 421, 7, 29, 2, 2,
	421, 423, 7, 38, 2, 2, 422, 407, 3, 2, 2, 2, 422, 411, 3, 2, 2, 2, 422,
	415, 3, 2, 2, 2, 422, 420, 3, 2, 2, 2, 423, 83, 3, 2, 2, 2, 424, 430, 7,
	70, 2, 2, 425, 426, 7, 70, 2, 2, 426, 427, 7, 48, 2, 2, 427, 430, 7, 70,
	2, 2, 428, 430, 5, 86, 44, 2, 429, 424, 3, 2, 2, 2, 429, 425, 3, 2, 2,
	2, 429, 428, 3, 2, 2, 2, 430, 85, 3, 2, 2, 2, 431, 436, 7, 70, 2, 2, 432,
	433, 7, 30, 2, 2, 433, 435, 7, 70, 2, 2, 434, 432, 3, 2, 2, 2, 435, 438,
	3, 2, 2, 2, 436, 434, 3, 2, 2, 2, 436, 437, 3, 2, 2, 2, 437, 87, 3, 2,
	2, 2, 438, 436, 3, 2, 2, 2, 439, 446, 5, 10, 6, 2, 440, 441, 5, 10, 6,
	2, 441, 442, 7, 48, 2, 2, 442, 443, 5, 10, 6, 2, 443, 446, 3, 2, 2, 2,
	444, 446, 5, 90, 46, 2, 445, 439, 3, 2, 2, 2, 445, 440, 3, 2, 2, 2, 445,
	444, 3, 2, 2, 2, 446, 89, 3, 2, 2, 2, 447, 452, 5, 10, 6, 2, 448, 449,
	7, 30, 2, 2, 449, 451, 5, 10, 6, 2, 450, 448, 3, 2, 2, 2, 451, 454, 3,
	2, 2, 2, 452, 450, 3, 2, 2, 2, 452, 453, 3, 2, 2, 2, 453, 91, 3, 2, 2,
	2, 454, 452, 3, 2, 2, 2, 455, 456, 9, 4, 2, 2, 456, 93, 3, 2, 2, 2, 457,
	458, 5, 22, 12, 2, 458, 459, 7, 70, 2, 2, 459, 95, 3, 2, 2, 2, 460, 461,
	8, 49, 1, 2, 461, 468, 5, 92, 47, 2, 462, 463, 7, 43, 2, 2, 463, 469, 7,
	33, 2, 2, 464, 465, 7, 42, 2, 2, 465, 469, 7, 33, 2, 2, 466, 469, 7, 42,
	2, 2, 467, 469, 7, 43, 2, 2, 468, 462, 3, 2, 2, 2, 468, 464, 3, 2, 2, 2,
	468, 466, 3, 2, 2, 2, 468, 467, 3, 2, 2, 2, 469, 470, 3, 2, 2, 2, 470,
	471, 7, 70, 2, 2, 471, 499, 3, 2, 2, 2, 472, 473, 5, 92, 47, 2, 473, 474,
	9, 5, 2, 2, 474, 475, 7, 70, 2, 2, 475, 499, 3, 2, 2, 2, 476, 477, 7, 65,
	2, 2, 477, 478, 9, 5, 2, 2, 478, 499, 5, 94, 48, 2, 479, 486, 7, 65, 2,
	2, 480, 481, 7, 43, 2, 2, 481, 487, 7, 33, 2, 2, 482, 483, 7, 42, 2, 2,
	483, 487, 7, 33, 2, 2, 484, 487, 7, 42, 2, 2, 485, 487, 7, 43, 2, 2, 486,
	480, 3, 2, 2, 2, 486, 482, 3, 2, 2, 2, 486, 484, 3, 2, 2, 2, 486, 485,
	3, 2, 2, 2, 487, 488, 3, 2, 2, 2, 488, 499, 5, 94, 48, 2, 489, 490, 7,
	50, 2, 2, 490, 499, 5, 32, 17, 2, 491, 492, 7, 66, 2, 2, 492, 493, 9, 5,
	2, 2, 493, 499, 5, 10, 6, 2, 494, 495, 7, 36, 2, 2, 495, 496, 5, 96, 49,
	2, 496, 497, 7, 37, 2, 2, 497, 499, 3, 2, 2, 2, 498, 460, 3, 2, 2, 2, 498,
	472, 3, 2, 2, 2, 498, 476, 3, 2, 2, 2, 498, 479, 3, 2, 2, 2, 498, 489,
	3, 2, 2, 2, 498, 491, 3, 2, 2, 2, 498, 494, 3, 2, 2, 2, 499, 508, 3, 2,
	2, 2, 500, 501, 12, 5, 2, 2, 501, 502, 7, 40, 2, 2, 502, 507, 5, 96, 49,
	6, 503, 504, 12, 4, 2, 2, 504, 505, 7, 41, 2, 2, 505, 507, 5, 96, 49, 5,
	506, 500, 3, 2, 2, 2, 506, 503, 3, 2, 2, 2, 507, 510, 3, 2, 2, 2, 508,
	506, 3, 2, 2, 2, 508, 509, 3, 2, 2, 2, 509, 97, 3, 2, 2, 2, 510, 508, 3,
	2, 2, 2, 50, 105, 124, 129, 136, 140, 152, 154, 163, 165, 171, 175, 187,
	189, 192, 199, 201, 221, 227, 235, 240, 242, 252, 257, 268, 299, 306, 313,
	320, 327, 334, 360, 369, 375, 383, 388, 394, 400, 405, 422, 429, 436, 445,
	452, 468, 486, 498, 506, 508,
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
	"para", "span", "media", "nid", "monthName", "position", "positionType",
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
	LMLParserRULE_span                      = 7
	LMLParserRULE_media                     = 8
	LMLParserRULE_nid                       = 9
	LMLParserRULE_monthName                 = 10
	LMLParserRULE_position                  = 11
	LMLParserRULE_positionType              = 12
	LMLParserRULE_directive                 = 13
	LMLParserRULE_lookup                    = 14
	LMLParserRULE_rid                       = 15
	LMLParserRULE_override                  = 16
	LMLParserRULE_overrideMode              = 17
	LMLParserRULE_overrideDay               = 18
	LMLParserRULE_sid                       = 19
	LMLParserRULE_tmplCalendar              = 20
	LMLParserRULE_tmplHtmlCss               = 21
	LMLParserRULE_tmplPdfCss                = 22
	LMLParserRULE_tmplDay                   = 23
	LMLParserRULE_tmplID                    = 24
	LMLParserRULE_tmplMonth                 = 25
	LMLParserRULE_tmplPageHeaderEven        = 26
	LMLParserRULE_tmplPageFooterEven        = 27
	LMLParserRULE_tmplPageHeaderOdd         = 28
	LMLParserRULE_tmplPageFooterOdd         = 29
	LMLParserRULE_tmplPageHeader            = 30
	LMLParserRULE_tmplPageFooter            = 31
	LMLParserRULE_tmplPageNumber            = 32
	LMLParserRULE_tmplStatus                = 33
	LMLParserRULE_tmplTitle                 = 34
	LMLParserRULE_tmplType                  = 35
	LMLParserRULE_tmplYear                  = 36
	LMLParserRULE_block                     = 37
	LMLParserRULE_switchBlock               = 38
	LMLParserRULE_switchBlockStatementGroup = 39
	LMLParserRULE_switchLabel               = 40
	LMLParserRULE_integerExpression         = 41
	LMLParserRULE_integerList               = 42
	LMLParserRULE_dowExpression             = 43
	LMLParserRULE_dowList                   = 44
	LMLParserRULE_ldpInt                    = 45
	LMLParserRULE_monthDay                  = 46
	LMLParserRULE_expression                = 47
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
		p.SetState(96)
		p.TmplID()
	}
	{
		p.SetState(97)
		p.TmplType()
	}
	{
		p.SetState(98)
		p.TmplStatus()
	}
	{
		p.SetState(99)
		p.PropertyBlock()
	}
	p.SetState(101)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
		{
			p.SetState(100)
			p.Statement()
		}

		p.SetState(103)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(105)
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

	p.SetState(122)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(107)
			p.TmplDay()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(108)
			p.TmplMonth()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(109)
			p.TmplYear()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(110)
			p.TmplHtmlCss()
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(111)
			p.TmplPdfCss()
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(112)
			p.TmplTitle()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(113)
			p.TmplDay()
		}

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(114)
			p.TmplCalendar()
		}

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(115)
			p.TmplPageNumber()
		}

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(116)
			p.TmplPageHeaderEven()
		}

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(117)
			p.TmplPageHeaderOdd()
		}

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(118)
			p.TmplPageFooterEven()
		}

	case 13:
		p.EnterOuterAlt(localctx, 13)
		{
			p.SetState(119)
			p.TmplPageFooterOdd()
		}

	case 14:
		p.EnterOuterAlt(localctx, 14)
		{
			p.SetState(120)
			p.TmplPageHeader()
		}

	case 15:
		p.EnterOuterAlt(localctx, 15)
		{
			p.SetState(121)
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

	p.SetState(138)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(127)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LMLParserT__9)|(1<<LMLParserT__10)|(1<<LMLParserT__11)|(1<<LMLParserT__12)|(1<<LMLParserT__14)|(1<<LMLParserT__15)|(1<<LMLParserT__16)|(1<<LMLParserT__17)|(1<<LMLParserT__18)|(1<<LMLParserT__19)|(1<<LMLParserT__20)|(1<<LMLParserT__21)|(1<<LMLParserT__23)|(1<<LMLParserT__25))) != 0 {
			{
				p.SetState(124)
				p.Property()
			}

			p.SetState(129)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(130)
			p.Match(LMLParserLBRACE)
		}
		p.SetState(134)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LMLParserT__9)|(1<<LMLParserT__10)|(1<<LMLParserT__11)|(1<<LMLParserT__12)|(1<<LMLParserT__14)|(1<<LMLParserT__15)|(1<<LMLParserT__16)|(1<<LMLParserT__17)|(1<<LMLParserT__18)|(1<<LMLParserT__19)|(1<<LMLParserT__20)|(1<<LMLParserT__21)|(1<<LMLParserT__23)|(1<<LMLParserT__25))) != 0 {
			{
				p.SetState(131)
				p.Property()
			}

			p.SetState(136)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(137)
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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(140)
			p.Media()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(141)
			p.Para()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(142)
			p.Insert()
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(143)
			p.Match(LMLParserIF)
		}
		{
			p.SetState(144)
			p.expression(0)
		}
		{
			p.SetState(145)
			p.Block()
		}
		p.SetState(152)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LMLParserELSE {
			{
				p.SetState(146)
				p.Match(LMLParserELSE)
			}
			p.SetState(148)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(147)
						p.Block()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(150)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext())
			}

		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(154)
			p.Match(LMLParserIF)
		}
		{
			p.SetState(155)
			p.expression(0)
		}
		{
			p.SetState(156)
			p.Block()
		}
		p.SetState(163)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == LMLParserT__0 {
			{
				p.SetState(157)
				p.Match(LMLParserT__0)
			}
			p.SetState(159)
			p.GetErrorHandler().Sync(p)
			_alt = 1
			for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
				switch _alt {
				case 1:
					{
						p.SetState(158)
						p.Block()
					}

				default:
					panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				}

				p.SetState(161)
				p.GetErrorHandler().Sync(p)
				_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext())
			}

		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(165)
			p.Match(LMLParserSWITCH)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LMLParserMODE_OF_WEEK, LMLParserMOVABLE_CYCLE_DAY, LMLParserLUKAN_CYCLE_DAY, LMLParserSUNDAY_AFTER_ELEVATION_OF_CROSS, LMLParserSUNDAYS_BEFORE_TRIODION:
			{
				p.SetState(166)
				p.LdpInt()
			}

		case LMLParserDATE:
			{
				p.SetState(167)
				p.Match(LMLParserDATE)
			}

		case LMLParserNAME_OF_DAY:
			{
				p.SetState(168)
				p.Match(LMLParserNAME_OF_DAY)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}
		{
			p.SetState(171)
			p.SwitchBlock()
		}

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(172)
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
		p.SetState(175)
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
		p.SetState(177)
		p.Match(LMLParserT__1)
	}
	{
		p.SetState(178)
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
		p.SetState(180)
		p.Match(LMLParserPARA_STYLE)
	}
	p.SetState(185)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<LMLParserT__3)|(1<<LMLParserT__5)|(1<<LMLParserT__8))) != 0) || _la == LMLParserSPAN_STYLE {
		p.SetState(185)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LMLParserT__3:
			{
				p.SetState(181)
				p.Nid()
			}

		case LMLParserT__5:
			{
				p.SetState(182)
				p.Rid()
			}

		case LMLParserT__8:
			{
				p.SetState(183)
				p.Sid()
			}

		case LMLParserSPAN_STYLE:
			{
				p.SetState(184)
				p.Span()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(187)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(190)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == LMLParserINSERT_VER {
		{
			p.SetState(189)
			p.Match(LMLParserINSERT_VER)
		}

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

func (s *SpanContext) AllSpan() []ISpanContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISpanContext)(nil)).Elem())
	var tst = make([]ISpanContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISpanContext)
		}
	}

	return tst
}

func (s *SpanContext) Span(i int) ISpanContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISpanContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISpanContext)
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
	p.EnterRule(localctx, 14, LMLParserRULE_span)

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
		p.SetState(192)
		p.Match(LMLParserSPAN_STYLE)
	}
	p.SetState(197)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			p.SetState(197)
			p.GetErrorHandler().Sync(p)

			switch p.GetTokenStream().LA(1) {
			case LMLParserT__3:
				{
					p.SetState(193)
					p.Nid()
				}

			case LMLParserT__5:
				{
					p.SetState(194)
					p.Rid()
				}

			case LMLParserT__8:
				{
					p.SetState(195)
					p.Sid()
				}

			case LMLParserSPAN_STYLE:
				{
					p.SetState(196)
					p.Span()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(199)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 16, LMLParserRULE_media)

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
		p.SetState(201)
		p.Match(LMLParserT__2)
	}
	{
		p.SetState(202)
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
	p.EnterRule(localctx, 18, LMLParserRULE_nid)

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
		p.SetState(204)
		p.Match(LMLParserT__3)
	}
	{
		p.SetState(205)
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
	p.EnterRule(localctx, 20, LMLParserRULE_monthName)

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

	p.SetState(219)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LMLParserJAN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(207)
			p.Match(LMLParserJAN)
		}

	case LMLParserFEB:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(208)
			p.Match(LMLParserFEB)
		}

	case LMLParserMAR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(209)
			p.Match(LMLParserMAR)
		}

	case LMLParserAPR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(210)
			p.Match(LMLParserAPR)
		}
		{
			p.SetState(211)
			p.Match(LMLParserMAY)
		}

	case LMLParserJUN:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(212)
			p.Match(LMLParserJUN)
		}

	case LMLParserJUL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(213)
			p.Match(LMLParserJUL)
		}

	case LMLParserAUG:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(214)
			p.Match(LMLParserAUG)
		}
		{
			p.SetState(215)
			p.Match(LMLParserSEP)
		}

	case LMLParserOCT:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(216)
			p.Match(LMLParserOCT)
		}

	case LMLParserNOV:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(217)
			p.Match(LMLParserNOV)
		}

	case LMLParserDEC:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(218)
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
	p.EnterRule(localctx, 22, LMLParserRULE_position)
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
		p.SetState(221)
		p.PositionType()
	}
	p.SetState(223)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserT__3 || (((_la-81)&-(0x1f+1)) == 0 && ((1<<uint((_la-81)))&((1<<(LMLParserINSERT_DATE-81))|(1<<(LMLParserINSERT_LOOKUP-81))|(1<<(LMLParserINSERT_PAGE_NUMBER-81)))) != 0) {
		{
			p.SetState(222)
			p.Directive()
		}

		p.SetState(225)
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
	p.EnterRule(localctx, 24, LMLParserRULE_positionType)
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
		p.SetState(227)
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
	p.EnterRule(localctx, 26, LMLParserRULE_directive)

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

	p.SetState(233)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LMLParserINSERT_LOOKUP:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(229)
			p.Lookup()
		}

	case LMLParserINSERT_DATE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(230)
			p.Match(LMLParserINSERT_DATE)
		}

	case LMLParserINSERT_PAGE_NUMBER:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(231)
			p.Match(LMLParserINSERT_PAGE_NUMBER)
		}

	case LMLParserT__3:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(232)
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
	p.EnterRule(localctx, 28, LMLParserRULE_lookup)
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
		p.SetState(235)
		p.Match(LMLParserINSERT_LOOKUP)
	}
	p.SetState(238)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = _la == LMLParserT__5 || _la == LMLParserT__8 {
		p.SetState(238)
		p.GetErrorHandler().Sync(p)

		switch p.GetTokenStream().LA(1) {
		case LMLParserT__8:
			{
				p.SetState(236)
				p.Sid()
			}

		case LMLParserT__5:
			{
				p.SetState(237)
				p.Rid()
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(240)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(242)
		p.Match(LMLParserT__4)
	}
	{
		p.SetState(243)
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
	p.EnterRule(localctx, 30, LMLParserRULE_rid)

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
		p.SetState(245)
		p.Match(LMLParserT__5)
	}
	{
		p.SetState(246)
		p.Match(LMLParserSTRING)
	}
	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(247)
				p.Override()
			}

		}
		p.SetState(252)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 32, LMLParserRULE_override)

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

	p.SetState(255)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case LMLParserT__6:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(253)
			p.OverrideMode()
		}

	case LMLParserT__7:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(254)
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
	p.EnterRule(localctx, 34, LMLParserRULE_overrideMode)

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
		p.SetState(257)
		p.Match(LMLParserT__6)
	}
	{
		p.SetState(258)
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
	p.EnterRule(localctx, 36, LMLParserRULE_overrideDay)

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
		p.SetState(260)
		p.Match(LMLParserT__7)
	}
	{
		p.SetState(261)
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
	p.EnterRule(localctx, 38, LMLParserRULE_sid)

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
		p.SetState(263)
		p.Match(LMLParserT__8)
	}
	{
		p.SetState(264)
		p.Match(LMLParserSTRING)
	}
	p.SetState(266)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 23, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(265)
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
	p.EnterRule(localctx, 40, LMLParserRULE_tmplCalendar)

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
		p.Match(LMLParserT__9)
	}
	{
		p.SetState(269)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(270)
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
	p.EnterRule(localctx, 42, LMLParserRULE_tmplHtmlCss)

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
		p.SetState(272)
		p.Match(LMLParserT__10)
	}
	{
		p.SetState(273)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(274)
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
	p.EnterRule(localctx, 44, LMLParserRULE_tmplPdfCss)

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
		p.SetState(276)
		p.Match(LMLParserT__11)
	}
	{
		p.SetState(277)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(278)
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
	p.EnterRule(localctx, 46, LMLParserRULE_tmplDay)

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
		p.SetState(280)
		p.Match(LMLParserT__12)
	}
	{
		p.SetState(281)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(282)
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
	p.EnterRule(localctx, 48, LMLParserRULE_tmplID)

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
		p.SetState(284)
		p.Match(LMLParserT__13)
	}
	{
		p.SetState(285)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(286)
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
	p.EnterRule(localctx, 50, LMLParserRULE_tmplMonth)

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
		p.SetState(288)
		p.Match(LMLParserT__14)
	}
	{
		p.SetState(289)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(290)
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
	p.EnterRule(localctx, 52, LMLParserRULE_tmplPageHeaderEven)
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
		p.SetState(292)
		p.Match(LMLParserT__15)
	}
	{
		p.SetState(293)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(294)
			p.Position()
		}

		p.SetState(297)
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
	p.EnterRule(localctx, 54, LMLParserRULE_tmplPageFooterEven)
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
		p.SetState(299)
		p.Match(LMLParserT__16)
	}
	{
		p.SetState(300)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(302)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(301)
			p.Position()
		}

		p.SetState(304)
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
	p.EnterRule(localctx, 56, LMLParserRULE_tmplPageHeaderOdd)
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
		p.SetState(306)
		p.Match(LMLParserT__17)
	}
	{
		p.SetState(307)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(309)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(308)
			p.Position()
		}

		p.SetState(311)
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
	p.EnterRule(localctx, 58, LMLParserRULE_tmplPageFooterOdd)
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
		p.SetState(313)
		p.Match(LMLParserT__18)
	}
	{
		p.SetState(314)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(315)
			p.Position()
		}

		p.SetState(318)
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
	p.EnterRule(localctx, 60, LMLParserRULE_tmplPageHeader)
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
		p.SetState(320)
		p.Match(LMLParserT__19)
	}
	{
		p.SetState(321)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(323)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(322)
			p.Position()
		}

		p.SetState(325)
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
	p.EnterRule(localctx, 62, LMLParserRULE_tmplPageFooter)
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
		p.SetState(327)
		p.Match(LMLParserT__20)
	}
	{
		p.SetState(328)
		p.Match(LMLParserASSIGNMENT)
	}
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = (((_la-76)&-(0x1f+1)) == 0 && ((1<<uint((_la-76)))&((1<<(LMLParserLEFT-76))|(1<<(LMLParserCENTER-76))|(1<<(LMLParserRIGHT-76)))) != 0) {
		{
			p.SetState(329)
			p.Position()
		}

		p.SetState(332)
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
	p.EnterRule(localctx, 64, LMLParserRULE_tmplPageNumber)

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
		p.SetState(334)
		p.Match(LMLParserT__21)
	}
	{
		p.SetState(335)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(336)
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
	p.EnterRule(localctx, 66, LMLParserRULE_tmplStatus)

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
		p.Match(LMLParserT__22)
	}
	{
		p.SetState(339)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(340)
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
	p.EnterRule(localctx, 68, LMLParserRULE_tmplTitle)

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
		p.Match(LMLParserT__23)
	}
	{
		p.SetState(343)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(344)
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
	p.EnterRule(localctx, 70, LMLParserRULE_tmplType)

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
		p.SetState(346)
		p.Match(LMLParserT__24)
	}
	{
		p.SetState(347)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(348)
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
	p.EnterRule(localctx, 72, LMLParserRULE_tmplYear)

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
		p.SetState(350)
		p.Match(LMLParserT__25)
	}
	{
		p.SetState(351)
		p.Match(LMLParserASSIGNMENT)
	}
	{
		p.SetState(352)
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
	p.EnterRule(localctx, 74, LMLParserRULE_block)
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
		p.SetState(354)
		p.Match(LMLParserLBRACE)
	}
	p.SetState(358)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
		{
			p.SetState(355)
			p.Statement()
		}

		p.SetState(360)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(361)
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
	p.EnterRule(localctx, 76, LMLParserRULE_switchBlock)
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
		p.SetState(363)
		p.Match(LMLParserLBRACE)
	}
	p.SetState(367)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			{
				p.SetState(364)
				p.SwitchBlockStatementGroup()
			}

		}
		p.SetState(369)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 31, p.GetParserRuleContext())
	}
	p.SetState(373)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__26 || _la == LMLParserCASE {
		{
			p.SetState(370)
			p.SwitchLabel()
		}

		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(376)
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
	p.EnterRule(localctx, 78, LMLParserRULE_switchBlockStatementGroup)
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

	p.SetState(403)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 37, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(379)
		p.GetErrorHandler().Sync(p)
		_alt = 1
		for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			switch _alt {
			case 1:
				{
					p.SetState(378)
					p.SwitchLabel()
				}

			default:
				panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			}

			p.SetState(381)
			p.GetErrorHandler().Sync(p)
			_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext())
		}
		p.SetState(386)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
			{
				p.SetState(383)
				p.Statement()
			}

			p.SetState(388)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		p.SetState(390)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = _la == LMLParserT__26 || _la == LMLParserCASE {
			{
				p.SetState(389)
				p.SwitchLabel()
			}

			p.SetState(392)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(394)
			p.Match(LMLParserLBRACE)
		}
		p.SetState(398)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == LMLParserT__1 || _la == LMLParserT__2 || (((_la-32)&-(0x1f+1)) == 0 && ((1<<uint((_la-32)))&((1<<(LMLParserLBRACE-32))|(1<<(LMLParserSWITCH-32))|(1<<(LMLParserIF-32)))) != 0) || _la == LMLParserPARA_STYLE {
			{
				p.SetState(395)
				p.Statement()
			}

			p.SetState(400)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(401)
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
	p.EnterRule(localctx, 80, LMLParserRULE_switchLabel)

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

	p.SetState(420)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 38, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(405)
			p.Match(LMLParserCASE)
		}
		{
			p.SetState(406)
			p.IntegerExpression()
		}
		{
			p.SetState(407)
			p.Match(LMLParserCOLON)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(409)
			p.Match(LMLParserCASE)
		}
		{
			p.SetState(410)
			p.DowExpression()
		}
		{
			p.SetState(411)
			p.Match(LMLParserCOLON)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(413)
			p.Match(LMLParserCASE)
		}
		{
			p.SetState(414)
			p.MonthName()
		}
		{
			p.SetState(415)
			p.IntegerExpression()
		}
		{
			p.SetState(416)
			p.Match(LMLParserCOLON)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(418)
			p.Match(LMLParserT__26)
		}
		{
			p.SetState(419)
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
	p.EnterRule(localctx, 82, LMLParserRULE_integerExpression)

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

	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 39, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(422)
			p.Match(LMLParserINTEGER)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(423)
			p.Match(LMLParserINTEGER)
		}
		{
			p.SetState(424)
			p.Match(LMLParserTHRU)
		}
		{
			p.SetState(425)
			p.Match(LMLParserINTEGER)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(426)
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
	p.EnterRule(localctx, 84, LMLParserRULE_integerList)
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
		p.SetState(429)
		p.Match(LMLParserINTEGER)
	}
	p.SetState(434)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__27 {
		{
			p.SetState(430)
			p.Match(LMLParserT__27)
		}
		{
			p.SetState(431)
			p.Match(LMLParserINTEGER)
		}

		p.SetState(436)
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
	p.EnterRule(localctx, 86, LMLParserRULE_dowExpression)

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

	p.SetState(443)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 41, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(437)
			p.DowName()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(438)
			p.DowName()
		}
		{
			p.SetState(439)
			p.Match(LMLParserTHRU)
		}
		{
			p.SetState(440)
			p.DowName()
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(442)
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
	p.EnterRule(localctx, 88, LMLParserRULE_dowList)
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
		p.SetState(445)
		p.DowName()
	}
	p.SetState(450)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == LMLParserT__27 {
		{
			p.SetState(446)
			p.Match(LMLParserT__27)
		}
		{
			p.SetState(447)
			p.DowName()
		}

		p.SetState(452)
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
	p.EnterRule(localctx, 90, LMLParserRULE_ldpInt)
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
		p.SetState(453)
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
	p.EnterRule(localctx, 92, LMLParserRULE_monthDay)

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
		p.SetState(455)
		p.MonthName()
	}
	{
		p.SetState(456)
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
	_startState := 94
	p.EnterRecursionRule(localctx, 94, LMLParserRULE_expression, _p)
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
	p.SetState(496)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 45, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(459)
			p.LdpInt()
		}
		p.SetState(466)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 43, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(460)
				p.Match(LMLParserLT)
			}
			{
				p.SetState(461)
				p.Match(LMLParserASSIGNMENT)
			}

		case 2:
			{
				p.SetState(462)
				p.Match(LMLParserGT)
			}
			{
				p.SetState(463)
				p.Match(LMLParserASSIGNMENT)
			}

		case 3:
			{
				p.SetState(464)
				p.Match(LMLParserGT)
			}

		case 4:
			{
				p.SetState(465)
				p.Match(LMLParserLT)
			}

		}
		{
			p.SetState(468)
			p.Match(LMLParserINTEGER)
		}

	case 2:
		{
			p.SetState(470)
			p.LdpInt()
		}
		{
			p.SetState(471)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LMLParserEQUALITY || _la == LMLParserINEQUALITY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(472)
			p.Match(LMLParserINTEGER)
		}

	case 3:
		{
			p.SetState(474)
			p.Match(LMLParserDATE)
		}
		{
			p.SetState(475)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LMLParserEQUALITY || _la == LMLParserINEQUALITY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(476)
			p.MonthDay()
		}

	case 4:
		{
			p.SetState(477)
			p.Match(LMLParserDATE)
		}
		p.SetState(484)
		p.GetErrorHandler().Sync(p)
		switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 44, p.GetParserRuleContext()) {
		case 1:
			{
				p.SetState(478)
				p.Match(LMLParserLT)
			}
			{
				p.SetState(479)
				p.Match(LMLParserASSIGNMENT)
			}

		case 2:
			{
				p.SetState(480)
				p.Match(LMLParserGT)
			}
			{
				p.SetState(481)
				p.Match(LMLParserASSIGNMENT)
			}

		case 3:
			{
				p.SetState(482)
				p.Match(LMLParserGT)
			}

		case 4:
			{
				p.SetState(483)
				p.Match(LMLParserLT)
			}

		}
		{
			p.SetState(486)
			p.MonthDay()
		}

	case 5:
		{
			p.SetState(487)
			p.Match(LMLParserEXISTS)
		}
		{
			p.SetState(488)
			p.Rid()
		}

	case 6:
		{
			p.SetState(489)
			p.Match(LMLParserNAME_OF_DAY)
		}
		{
			p.SetState(490)
			_la = p.GetTokenStream().LA(1)

			if !(_la == LMLParserEQUALITY || _la == LMLParserINEQUALITY) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(491)
			p.DowName()
		}

	case 7:
		{
			p.SetState(492)
			p.Match(LMLParserLPAREN)
		}
		{
			p.SetState(493)
			p.expression(0)
		}
		{
			p.SetState(494)
			p.Match(LMLParserRPAREN)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(506)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(504)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 46, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LMLParserRULE_expression)
				p.SetState(498)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(499)
					p.Match(LMLParserAND)
				}
				{
					p.SetState(500)
					p.expression(4)
				}

			case 2:
				localctx = NewExpressionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, LMLParserRULE_expression)
				p.SetState(501)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(502)
					p.Match(LMLParserOR)
				}
				{
					p.SetState(503)
					p.expression(3)
				}

			}

		}
		p.SetState(508)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 47, p.GetParserRuleContext())
	}

	return localctx
}

func (p *LMLParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 47:
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
