// Generated from SqlParser.g4 by ANTLR 4.7.

package parser // SqlParser

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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 239, 321,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 3, 2, 3, 2, 3,
	2, 3, 2, 3, 2, 5, 2, 60, 10, 2, 3, 2, 5, 2, 63, 10, 2, 3, 2, 5, 2, 66,
	10, 2, 3, 2, 5, 2, 69, 10, 2, 3, 2, 5, 2, 72, 10, 2, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 7, 3, 81, 10, 3, 12, 3, 14, 3, 84, 11, 3, 3, 4,
	3, 4, 5, 4, 88, 10, 4, 3, 5, 3, 5, 3, 5, 7, 5, 93, 10, 5, 12, 5, 14, 5,
	96, 11, 5, 3, 6, 3, 6, 3, 6, 5, 6, 101, 10, 6, 3, 6, 3, 6, 3, 6, 3, 6,
	5, 6, 107, 10, 6, 5, 6, 109, 10, 6, 3, 7, 3, 7, 3, 7, 3, 8, 3, 8, 3, 8,
	3, 8, 3, 8, 7, 8, 119, 10, 8, 12, 8, 14, 8, 122, 11, 8, 3, 9, 3, 9, 3,
	9, 3, 10, 3, 10, 3, 10, 3, 10, 5, 10, 131, 10, 10, 3, 10, 3, 10, 3, 10,
	3, 10, 3, 10, 5, 10, 138, 10, 10, 3, 11, 3, 11, 3, 12, 3, 12, 3, 12, 5,
	12, 145, 10, 12, 3, 12, 3, 12, 3, 13, 3, 13, 3, 14, 3, 14, 3, 15, 6, 15,
	154, 10, 15, 13, 15, 14, 15, 155, 3, 16, 3, 16, 3, 17, 3, 17, 5, 17, 162,
	10, 17, 3, 17, 3, 17, 5, 17, 166, 10, 17, 3, 18, 3, 18, 3, 18, 3, 18, 5,
	18, 172, 10, 18, 3, 18, 3, 18, 5, 18, 176, 10, 18, 3, 19, 3, 19, 3, 19,
	5, 19, 181, 10, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3,
	19, 3, 19, 3, 19, 5, 19, 193, 10, 19, 3, 19, 3, 19, 3, 19, 5, 19, 198,
	10, 19, 3, 20, 3, 20, 3, 20, 7, 20, 203, 10, 20, 12, 20, 14, 20, 206, 11,
	20, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 3, 21, 7, 21, 215, 10, 21,
	12, 21, 14, 21, 218, 11, 21, 3, 22, 3, 22, 5, 22, 222, 10, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 7, 22, 229, 10, 22, 12, 22, 14, 22, 232, 11, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 239, 10, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 249, 10, 22, 3, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 258, 10, 22, 3, 22, 3, 22, 3,
	22, 3, 22, 5, 22, 264, 10, 22, 3, 22, 3, 22, 5, 22, 268, 10, 22, 3, 22,
	3, 22, 3, 22, 3, 22, 5, 22, 274, 10, 22, 3, 23, 3, 23, 3, 23, 3, 23, 3,
	23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 5, 23, 287, 10, 23, 3, 23,
	3, 23, 3, 23, 3, 23, 7, 23, 293, 10, 23, 12, 23, 14, 23, 296, 11, 23, 3,
	24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 25,
	5, 25, 309, 10, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 5, 26, 317,
	10, 26, 3, 27, 3, 27, 3, 27, 2, 4, 40, 44, 28, 2, 4, 6, 8, 10, 12, 14,
	16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50,
	52, 2, 10, 4, 2, 14, 14, 43, 43, 4, 2, 218, 220, 228, 228, 4, 2, 58, 58,
	162, 162, 4, 2, 180, 180, 182, 184, 4, 2, 8, 8, 47, 47, 4, 2, 125, 125,
	136, 136, 6, 2, 104, 104, 199, 199, 201, 201, 207, 208, 4, 2, 196, 199,
	201, 201, 2, 347, 2, 54, 3, 2, 2, 2, 4, 75, 3, 2, 2, 2, 6, 85, 3, 2, 2,
	2, 8, 89, 3, 2, 2, 2, 10, 108, 3, 2, 2, 2, 12, 110, 3, 2, 2, 2, 14, 113,
	3, 2, 2, 2, 16, 123, 3, 2, 2, 2, 18, 126, 3, 2, 2, 2, 20, 139, 3, 2, 2,
	2, 22, 144, 3, 2, 2, 2, 24, 148, 3, 2, 2, 2, 26, 150, 3, 2, 2, 2, 28, 153,
	3, 2, 2, 2, 30, 157, 3, 2, 2, 2, 32, 165, 3, 2, 2, 2, 34, 175, 3, 2, 2,
	2, 36, 197, 3, 2, 2, 2, 38, 199, 3, 2, 2, 2, 40, 207, 3, 2, 2, 2, 42, 273,
	3, 2, 2, 2, 44, 286, 3, 2, 2, 2, 46, 297, 3, 2, 2, 2, 48, 308, 3, 2, 2,
	2, 50, 316, 3, 2, 2, 2, 52, 318, 3, 2, 2, 2, 54, 55, 7, 139, 2, 2, 55,
	56, 5, 8, 5, 2, 56, 57, 7, 63, 2, 2, 57, 59, 5, 24, 13, 2, 58, 60, 5, 12,
	7, 2, 59, 58, 3, 2, 2, 2, 59, 60, 3, 2, 2, 2, 60, 62, 3, 2, 2, 2, 61, 63,
	5, 14, 8, 2, 62, 61, 3, 2, 2, 2, 62, 63, 3, 2, 2, 2, 63, 65, 3, 2, 2, 2,
	64, 66, 5, 16, 9, 2, 65, 64, 3, 2, 2, 2, 65, 66, 3, 2, 2, 2, 66, 68, 3,
	2, 2, 2, 67, 69, 5, 4, 3, 2, 68, 67, 3, 2, 2, 2, 68, 69, 3, 2, 2, 2, 69,
	71, 3, 2, 2, 2, 70, 72, 5, 18, 10, 2, 71, 70, 3, 2, 2, 2, 71, 72, 3, 2,
	2, 2, 72, 73, 3, 2, 2, 2, 73, 74, 7, 2, 2, 3, 74, 3, 3, 2, 2, 2, 75, 76,
	7, 113, 2, 2, 76, 77, 7, 18, 2, 2, 77, 82, 5, 6, 4, 2, 78, 79, 7, 215,
	2, 2, 79, 81, 5, 6, 4, 2, 80, 78, 3, 2, 2, 2, 81, 84, 3, 2, 2, 2, 82, 80,
	3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 5, 3, 2, 2, 2, 84, 82, 3, 2, 2, 2,
	85, 87, 5, 40, 21, 2, 86, 88, 9, 2, 2, 2, 87, 86, 3, 2, 2, 2, 87, 88, 3,
	2, 2, 2, 88, 7, 3, 2, 2, 2, 89, 94, 5, 10, 6, 2, 90, 91, 7, 215, 2, 2,
	91, 93, 5, 10, 6, 2, 92, 90, 3, 2, 2, 2, 93, 96, 3, 2, 2, 2, 94, 92, 3,
	2, 2, 2, 94, 95, 3, 2, 2, 2, 95, 9, 3, 2, 2, 2, 96, 94, 3, 2, 2, 2, 97,
	98, 5, 24, 13, 2, 98, 99, 7, 212, 2, 2, 99, 101, 3, 2, 2, 2, 100, 97, 3,
	2, 2, 2, 100, 101, 3, 2, 2, 2, 101, 102, 3, 2, 2, 2, 102, 109, 7, 196,
	2, 2, 103, 106, 5, 44, 23, 2, 104, 105, 7, 13, 2, 2, 105, 107, 5, 24, 13,
	2, 106, 104, 3, 2, 2, 2, 106, 107, 3, 2, 2, 2, 107, 109, 3, 2, 2, 2, 108,
	100, 3, 2, 2, 2, 108, 103, 3, 2, 2, 2, 109, 11, 3, 2, 2, 2, 110, 111, 7,
	174, 2, 2, 111, 112, 5, 40, 21, 2, 112, 13, 3, 2, 2, 2, 113, 114, 7, 68,
	2, 2, 114, 115, 7, 18, 2, 2, 115, 120, 5, 40, 21, 2, 116, 117, 7, 215,
	2, 2, 117, 119, 5, 40, 21, 2, 118, 116, 3, 2, 2, 2, 119, 122, 3, 2, 2,
	2, 120, 118, 3, 2, 2, 2, 120, 121, 3, 2, 2, 2, 121, 15, 3, 2, 2, 2, 122,
	120, 3, 2, 2, 2, 123, 124, 7, 69, 2, 2, 124, 125, 5, 40, 21, 2, 125, 17,
	3, 2, 2, 2, 126, 137, 7, 91, 2, 2, 127, 128, 5, 20, 11, 2, 128, 129, 7,
	215, 2, 2, 129, 131, 3, 2, 2, 2, 130, 127, 3, 2, 2, 2, 130, 131, 3, 2,
	2, 2, 131, 132, 3, 2, 2, 2, 132, 138, 5, 20, 11, 2, 133, 134, 5, 20, 11,
	2, 134, 135, 7, 185, 2, 2, 135, 136, 5, 20, 11, 2, 136, 138, 3, 2, 2, 2,
	137, 130, 3, 2, 2, 2, 137, 133, 3, 2, 2, 2, 138, 19, 3, 2, 2, 2, 139, 140,
	5, 26, 14, 2, 140, 21, 3, 2, 2, 2, 141, 142, 5, 24, 13, 2, 142, 143, 7,
	212, 2, 2, 143, 145, 3, 2, 2, 2, 144, 141, 3, 2, 2, 2, 144, 145, 3, 2,
	2, 2, 145, 146, 3, 2, 2, 2, 146, 147, 5, 24, 13, 2, 147, 23, 3, 2, 2, 2,
	148, 149, 7, 233, 2, 2, 149, 25, 3, 2, 2, 2, 150, 151, 9, 3, 2, 2, 151,
	27, 3, 2, 2, 2, 152, 154, 7, 227, 2, 2, 153, 152, 3, 2, 2, 2, 154, 155,
	3, 2, 2, 2, 155, 153, 3, 2, 2, 2, 155, 156, 3, 2, 2, 2, 156, 29, 3, 2,
	2, 2, 157, 158, 9, 4, 2, 2, 158, 31, 3, 2, 2, 2, 159, 166, 5, 28, 15, 2,
	160, 162, 7, 201, 2, 2, 161, 160, 3, 2, 2, 2, 161, 162, 3, 2, 2, 2, 162,
	163, 3, 2, 2, 2, 163, 166, 5, 26, 14, 2, 164, 166, 5, 30, 16, 2, 165, 159,
	3, 2, 2, 2, 165, 161, 3, 2, 2, 2, 165, 164, 3, 2, 2, 2, 166, 33, 3, 2,
	2, 2, 167, 176, 5, 36, 19, 2, 168, 169, 5, 24, 13, 2, 169, 171, 7, 213,
	2, 2, 170, 172, 5, 38, 20, 2, 171, 170, 3, 2, 2, 2, 171, 172, 3, 2, 2,
	2, 172, 173, 3, 2, 2, 2, 173, 174, 7, 214, 2, 2, 174, 176, 3, 2, 2, 2,
	175, 167, 3, 2, 2, 2, 175, 168, 3, 2, 2, 2, 176, 35, 3, 2, 2, 2, 177, 178,
	9, 5, 2, 2, 178, 180, 7, 213, 2, 2, 179, 181, 9, 6, 2, 2, 180, 179, 3,
	2, 2, 2, 180, 181, 3, 2, 2, 2, 181, 182, 3, 2, 2, 2, 182, 183, 5, 40, 21,
	2, 183, 184, 7, 214, 2, 2, 184, 198, 3, 2, 2, 2, 185, 186, 7, 181, 2, 2,
	186, 187, 7, 213, 2, 2, 187, 188, 7, 196, 2, 2, 188, 198, 7, 214, 2, 2,
	189, 190, 7, 181, 2, 2, 190, 192, 7, 213, 2, 2, 191, 193, 9, 6, 2, 2, 192,
	191, 3, 2, 2, 2, 192, 193, 3, 2, 2, 2, 193, 194, 3, 2, 2, 2, 194, 195,
	5, 38, 20, 2, 195, 196, 7, 214, 2, 2, 196, 198, 3, 2, 2, 2, 197, 177, 3,
	2, 2, 2, 197, 185, 3, 2, 2, 2, 197, 189, 3, 2, 2, 2, 198, 37, 3, 2, 2,
	2, 199, 204, 5, 40, 21, 2, 200, 201, 7, 215, 2, 2, 201, 203, 5, 40, 21,
	2, 202, 200, 3, 2, 2, 2, 203, 206, 3, 2, 2, 2, 204, 202, 3, 2, 2, 2, 204,
	205, 3, 2, 2, 2, 205, 39, 3, 2, 2, 2, 206, 204, 3, 2, 2, 2, 207, 208, 8,
	21, 1, 2, 208, 209, 5, 42, 22, 2, 209, 216, 3, 2, 2, 2, 210, 211, 12, 4,
	2, 2, 211, 212, 5, 50, 26, 2, 212, 213, 5, 40, 21, 5, 213, 215, 3, 2, 2,
	2, 214, 210, 3, 2, 2, 2, 215, 218, 3, 2, 2, 2, 216, 214, 3, 2, 2, 2, 216,
	217, 3, 2, 2, 2, 217, 41, 3, 2, 2, 2, 218, 216, 3, 2, 2, 2, 219, 221, 5,
	44, 23, 2, 220, 222, 7, 104, 2, 2, 221, 220, 3, 2, 2, 2, 221, 222, 3, 2,
	2, 2, 222, 223, 3, 2, 2, 2, 223, 224, 7, 73, 2, 2, 224, 225, 7, 213, 2,
	2, 225, 230, 5, 44, 23, 2, 226, 227, 7, 215, 2, 2, 227, 229, 5, 44, 23,
	2, 228, 226, 3, 2, 2, 2, 229, 232, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 230,
	231, 3, 2, 2, 2, 231, 233, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 233, 234,
	7, 214, 2, 2, 234, 274, 3, 2, 2, 2, 235, 236, 5, 44, 23, 2, 236, 238, 7,
	81, 2, 2, 237, 239, 7, 104, 2, 2, 238, 237, 3, 2, 2, 2, 238, 239, 3, 2,
	2, 2, 239, 240, 3, 2, 2, 2, 240, 241, 7, 106, 2, 2, 241, 274, 3, 2, 2,
	2, 242, 243, 5, 44, 23, 2, 243, 244, 5, 48, 25, 2, 244, 245, 5, 44, 23,
	2, 245, 274, 3, 2, 2, 2, 246, 248, 5, 44, 23, 2, 247, 249, 7, 104, 2, 2,
	248, 247, 3, 2, 2, 2, 248, 249, 3, 2, 2, 2, 249, 250, 3, 2, 2, 2, 250,
	251, 7, 16, 2, 2, 251, 252, 5, 44, 23, 2, 252, 253, 7, 12, 2, 2, 253, 254,
	5, 44, 23, 2, 254, 274, 3, 2, 2, 2, 255, 257, 5, 44, 23, 2, 256, 258, 7,
	104, 2, 2, 257, 256, 3, 2, 2, 2, 257, 258, 3, 2, 2, 2, 258, 259, 3, 2,
	2, 2, 259, 260, 7, 90, 2, 2, 260, 263, 5, 44, 23, 2, 261, 262, 7, 186,
	2, 2, 262, 264, 7, 227, 2, 2, 263, 261, 3, 2, 2, 2, 263, 264, 3, 2, 2,
	2, 264, 274, 3, 2, 2, 2, 265, 267, 5, 44, 23, 2, 266, 268, 7, 104, 2, 2,
	267, 266, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269,
	270, 9, 7, 2, 2, 270, 271, 5, 44, 23, 2, 271, 274, 3, 2, 2, 2, 272, 274,
	5, 44, 23, 2, 273, 219, 3, 2, 2, 2, 273, 235, 3, 2, 2, 2, 273, 242, 3,
	2, 2, 2, 273, 246, 3, 2, 2, 2, 273, 255, 3, 2, 2, 2, 273, 265, 3, 2, 2,
	2, 273, 272, 3, 2, 2, 2, 274, 43, 3, 2, 2, 2, 275, 276, 8, 23, 1, 2, 276,
	287, 5, 32, 17, 2, 277, 287, 5, 22, 12, 2, 278, 287, 5, 34, 18, 2, 279,
	280, 5, 46, 24, 2, 280, 281, 5, 44, 23, 5, 281, 287, 3, 2, 2, 2, 282, 283,
	7, 213, 2, 2, 283, 284, 5, 40, 21, 2, 284, 285, 7, 214, 2, 2, 285, 287,
	3, 2, 2, 2, 286, 275, 3, 2, 2, 2, 286, 277, 3, 2, 2, 2, 286, 278, 3, 2,
	2, 2, 286, 279, 3, 2, 2, 2, 286, 282, 3, 2, 2, 2, 287, 294, 3, 2, 2, 2,
	288, 289, 12, 3, 2, 2, 289, 290, 5, 52, 27, 2, 290, 291, 5, 44, 23, 4,
	291, 293, 3, 2, 2, 2, 292, 288, 3, 2, 2, 2, 293, 296, 3, 2, 2, 2, 294,
	292, 3, 2, 2, 2, 294, 295, 3, 2, 2, 2, 295, 45, 3, 2, 2, 2, 296, 294, 3,
	2, 2, 2, 297, 298, 9, 8, 2, 2, 298, 47, 3, 2, 2, 2, 299, 309, 7, 204, 2,
	2, 300, 309, 7, 205, 2, 2, 301, 309, 7, 206, 2, 2, 302, 303, 7, 206, 2,
	2, 303, 309, 7, 204, 2, 2, 304, 305, 7, 205, 2, 2, 305, 309, 7, 204, 2,
	2, 306, 307, 7, 207, 2, 2, 307, 309, 7, 204, 2, 2, 308, 299, 3, 2, 2, 2,
	308, 300, 3, 2, 2, 2, 308, 301, 3, 2, 2, 2, 308, 302, 3, 2, 2, 2, 308,
	304, 3, 2, 2, 2, 308, 306, 3, 2, 2, 2, 309, 49, 3, 2, 2, 2, 310, 317, 7,
	12, 2, 2, 311, 312, 7, 210, 2, 2, 312, 317, 7, 210, 2, 2, 313, 317, 7,
	112, 2, 2, 314, 315, 7, 209, 2, 2, 315, 317, 7, 209, 2, 2, 316, 310, 3,
	2, 2, 2, 316, 311, 3, 2, 2, 2, 316, 313, 3, 2, 2, 2, 316, 314, 3, 2, 2,
	2, 317, 51, 3, 2, 2, 2, 318, 319, 9, 9, 2, 2, 319, 53, 3, 2, 2, 2, 39,
	59, 62, 65, 68, 71, 82, 87, 94, 100, 106, 108, 120, 130, 137, 144, 155,
	161, 165, 171, 175, 180, 192, 197, 204, 216, 221, 230, 238, 248, 257, 263,
	267, 273, 286, 294, 308, 316,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "", "", "", "", "'ADD'", "'ALL'", "'ALTER'", "'ALWAYS'", "'ANALYZE'",
	"'AND'", "'AS'", "'ASC'", "'BEFORE'", "'BETWEEN'", "'BOTH'", "'BY'", "'CALL'",
	"'CASCADE'", "'CASE'", "'CAST'", "'CHANGE'", "'CHARACTER'", "'CHECK'",
	"'COLLATE'", "'COLUMN'", "'CONDITION'", "'CONSTRAINT'", "'CONTINUE'", "'CONVERT'",
	"'CREATE'", "'CROSS'", "'CURRENT'", "'CURRENT_USER'", "'CURSOR'", "'DATABASE'",
	"'DATABASES'", "'DECLARE'", "'DEFAULT'", "'DELAYED'", "'DELETE'", "'DESC'",
	"'DESCRIBE'", "'DETERMINISTIC'", "'DIAGNOSTICS'", "'DISTINCT'", "'DISTINCTROW'",
	"'DROP'", "'EACH'", "'ELSE'", "'ELSEIF'", "'ENCLOSED'", "'ESCAPED'", "'EXISTS'",
	"'EXIT'", "'EXPLAIN'", "'FALSE'", "'FETCH'", "'FOR'", "'FORCE'", "'FOREIGN'",
	"'FROM'", "'FULLTEXT'", "'GENERATED'", "'GET'", "'GRANT'", "'GROUP'", "'HAVING'",
	"'HIGH_PRIORITY'", "'IF'", "'IGNORE'", "'IN'", "'INDEX'", "'INFILE'", "'INNER'",
	"'INOUT'", "'INSERT'", "'INTERVAL'", "'INTO'", "'IS'", "'ITERATE'", "'JOIN'",
	"'KEY'", "'KEYS'", "'KILL'", "'LEADING'", "'LEAVE'", "'LEFT'", "'LIKE'",
	"'LIMIT'", "'LINEAR'", "'LINES'", "'LOAD'", "'LOCK'", "'LOOP'", "'LOW_PRIORITY'",
	"'MASTER_BIND'", "'MASTER_SSL_VERIFY_SERVER_CERT'", "'MATCH'", "'MAXVALUE'",
	"'MODIFIES'", "'NATURAL'", "'NOT'", "'NO_WRITE_TO_BINLOG'", "'NULL'", "'NUMBER'",
	"'ON'", "'OPTIMIZE'", "'OPTION'", "'OPTIONALLY'", "'OR'", "'ORDER'", "'OUT'",
	"'OUTER'", "'OUTFILE'", "'PARTITION'", "'PRIMARY'", "'PROCEDURE'", "'PURGE'",
	"'RANGE'", "'READ'", "'READS'", "'REFERENCES'", "'REGEXP'", "'RELEASE'",
	"'RENAME'", "'REPEAT'", "'REPLACE'", "'REQUIRE'", "'RESIGNAL'", "'RESTRICT'",
	"'RETURN'", "'REVOKE'", "'RIGHT'", "'RLIKE'", "'SCHEMA'", "'SCHEMAS'",
	"'SELECT'", "'SET'", "'SEPARATOR'", "'SHOW'", "'SIGNAL'", "'SPATIAL'",
	"'SQL'", "'SQLEXCEPTION'", "'SQLSTATE'", "'SQLWARNING'", "'SQL_BIG_RESULT'",
	"'SQL_CALC_FOUND_ROWS'", "'SQL_SMALL_RESULT'", "'SSL'", "'STACKED'", "'STARTING'",
	"'STRAIGHT_JOIN'", "'TABLE'", "'TERMINATED'", "'THEN'", "'TO'", "'TRAILING'",
	"'TRIGGER'", "'TRUE'", "'UNDO'", "'UNION'", "'UNIQUE'", "'UNLOCK'", "'UNSIGNED'",
	"'UPDATE'", "'USAGE'", "'USE'", "'USING'", "'VALUES'", "'WHEN'", "'WHERE'",
	"'WHILE'", "'WITH'", "'WRITE'", "'XOR'", "'ZEROFILL'", "'AVG'", "'COUNT'",
	"'MAX'", "'MIN'", "'SUM'", "'OFFSET'", "'ESCAPE'", "':='", "'+='", "'-='",
	"'*='", "'/='", "'%='", "'&='", "'^='", "'|='", "'*'", "'/'", "'%'", "'+'",
	"'--'", "'-'", "'DIV'", "'MOD'", "'='", "'>'", "'<'", "'!'", "'~'", "'|'",
	"'&'", "'^'", "'.'", "'('", "')'", "','", "';'", "'@'", "'0'", "'1'", "'2'",
	"'''", "'\"'", "'`'", "':'",
}
var symbolicNames = []string{
	"", "SPACE", "SPEC_MYSQL_COMMENT", "COMMENT_INPUT", "LINE_COMMENT", "ADD",
	"ALL", "ALTER", "ALWAYS", "ANALYZE", "AND", "AS", "ASC", "BEFORE", "BETWEEN",
	"BOTH", "BY", "CALL", "CASCADE", "CASE", "CAST", "CHANGE", "CHARACTER",
	"CHECK", "COLLATE", "COLUMN", "CONDITION", "CONSTRAINT", "CONTINUE", "CONVERT",
	"CREATE", "CROSS", "CURRENT", "CURRENT_USER", "CURSOR", "DATABASE", "DATABASES",
	"DECLARE", "DEFAULT", "DELAYED", "DELETE", "DESC", "DESCRIBE", "DETERMINISTIC",
	"DIAGNOSTICS", "DISTINCT", "DISTINCTROW", "DROP", "EACH", "ELSE", "ELSEIF",
	"ENCLOSED", "ESCAPED", "EXISTS", "EXIT", "EXPLAIN", "FALSE", "FETCH", "FOR",
	"FORCE", "FOREIGN", "FROM", "FULLTEXT", "GENERATED", "GET", "GRANT", "GROUP",
	"HAVING", "HIGH_PRIORITY", "IF", "IGNORE", "IN", "INDEX", "INFILE", "INNER",
	"INOUT", "INSERT", "INTERVAL", "INTO", "IS", "ITERATE", "JOIN", "KEY",
	"KEYS", "KILL", "LEADING", "LEAVE", "LEFT", "LIKE", "LIMIT", "LINEAR",
	"LINES", "LOAD", "LOCK", "LOOP", "LOW_PRIORITY", "MASTER_BIND", "MASTER_SSL_VERIFY_SERVER_CERT",
	"MATCH", "MAXVALUE", "MODIFIES", "NATURAL", "NOT", "NO_WRITE_TO_BINLOG",
	"NULL_LITERAL", "NUMBER", "ON", "OPTIMIZE", "OPTION", "OPTIONALLY", "OR",
	"ORDER", "OUT", "OUTER", "OUTFILE", "PARTITION", "PRIMARY", "PROCEDURE",
	"PURGE", "RANGE", "READ", "READS", "REFERENCES", "REGEXP", "RELEASE", "RENAME",
	"REPEAT", "REPLACE", "REQUIRE", "RESIGNAL", "RESTRICT", "RETURN", "REVOKE",
	"RIGHT", "RLIKE", "SCHEMA", "SCHEMAS", "SELECT", "SET", "SEPARATOR", "SHOW",
	"SIGNAL", "SPATIAL", "SQL", "SQLEXCEPTION", "SQLSTATE", "SQLWARNING", "SQL_BIG_RESULT",
	"SQL_CALC_FOUND_ROWS", "SQL_SMALL_RESULT", "SSL", "STACKED", "STARTING",
	"STRAIGHT_JOIN", "TABLE", "TERMINATED", "THEN", "TO", "TRAILING", "TRIGGER",
	"TRUE", "UNDO", "UNION", "UNIQUE", "UNLOCK", "UNSIGNED", "UPDATE", "USAGE",
	"USE", "USING", "VALUES", "WHEN", "WHERE", "WHILE", "WITH", "WRITE", "XOR",
	"ZEROFILL", "AVG", "COUNT", "MAX", "MIN", "SUM", "OFFSET", "ESCAPE", "VAR_ASSIGN",
	"PLUS_ASSIGN", "MINUS_ASSIGN", "MULT_ASSIGN", "DIV_ASSIGN", "MOD_ASSIGN",
	"AND_ASSIGN", "XOR_ASSIGN", "OR_ASSIGN", "STAR", "DIVIDE", "MODULE", "PLUS",
	"MINUSMINUS", "MINUS", "DIV", "MOD", "EQUAL_SYMBOL", "GREATER_SYMBOL",
	"LESS_SYMBOL", "EXCLAMATION_SYMBOL", "BIT_NOT_OP", "BIT_OR_OP", "BIT_AND_OP",
	"BIT_XOR_OP", "DOT", "LR_BRACKET", "RR_BRACKET", "COMMA", "SEMI", "AT_SIGN",
	"ZERO_DECIMAL", "ONE_DECIMAL", "TWO_DECIMAL", "SINGLE_QUOTE_SYMB", "DOUBLE_QUOTE_SYMB",
	"REVERSE_QUOTE_SYMB", "COLON_SYMB", "FILESIZE_LITERAL", "START_NATIONAL_STRING_LITERAL",
	"STRING_LITERAL", "DECIMAL_LITERAL", "HEXADECIMAL_LITERAL", "REAL_LITERAL",
	"NULL_SPEC_LITERAL", "DOT_ID", "ID", "REVERSE_QUOTE_ID", "STRING_USER_NAME",
	"LOCAL_ID", "GLOBAL_ID", "ID_LITERAL", "ERROR_RECONGNIGION",
}

var ruleNames = []string{
	"selectStatement", "orderByClause", "orderByExpression", "selectElements",
	"selectElement", "whereClause", "groupbyClause", "havingClause", "limitClause",
	"limitClauseAtom", "columnName", "uid", "decimalLiteral", "stringLiteral",
	"booleanLiteral", "constant", "functionCall", "aggregateWindowedFunction",
	"functionArgs", "expression", "predicate", "expressionAtom", "unaryOperator",
	"comparisonOperator", "logicalOperator", "mathOperator",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type SqlParser struct {
	*antlr.BaseParser
}

func NewSqlParser(input antlr.TokenStream) *SqlParser {
	this := new(SqlParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "SqlParser.g4"

	return this
}

// SqlParser tokens.
const (
	SqlParserEOF                           = antlr.TokenEOF
	SqlParserSPACE                         = 1
	SqlParserSPEC_MYSQL_COMMENT            = 2
	SqlParserCOMMENT_INPUT                 = 3
	SqlParserLINE_COMMENT                  = 4
	SqlParserADD                           = 5
	SqlParserALL                           = 6
	SqlParserALTER                         = 7
	SqlParserALWAYS                        = 8
	SqlParserANALYZE                       = 9
	SqlParserAND                           = 10
	SqlParserAS                            = 11
	SqlParserASC                           = 12
	SqlParserBEFORE                        = 13
	SqlParserBETWEEN                       = 14
	SqlParserBOTH                          = 15
	SqlParserBY                            = 16
	SqlParserCALL                          = 17
	SqlParserCASCADE                       = 18
	SqlParserCASE                          = 19
	SqlParserCAST                          = 20
	SqlParserCHANGE                        = 21
	SqlParserCHARACTER                     = 22
	SqlParserCHECK                         = 23
	SqlParserCOLLATE                       = 24
	SqlParserCOLUMN                        = 25
	SqlParserCONDITION                     = 26
	SqlParserCONSTRAINT                    = 27
	SqlParserCONTINUE                      = 28
	SqlParserCONVERT                       = 29
	SqlParserCREATE                        = 30
	SqlParserCROSS                         = 31
	SqlParserCURRENT                       = 32
	SqlParserCURRENT_USER                  = 33
	SqlParserCURSOR                        = 34
	SqlParserDATABASE                      = 35
	SqlParserDATABASES                     = 36
	SqlParserDECLARE                       = 37
	SqlParserDEFAULT                       = 38
	SqlParserDELAYED                       = 39
	SqlParserDELETE                        = 40
	SqlParserDESC                          = 41
	SqlParserDESCRIBE                      = 42
	SqlParserDETERMINISTIC                 = 43
	SqlParserDIAGNOSTICS                   = 44
	SqlParserDISTINCT                      = 45
	SqlParserDISTINCTROW                   = 46
	SqlParserDROP                          = 47
	SqlParserEACH                          = 48
	SqlParserELSE                          = 49
	SqlParserELSEIF                        = 50
	SqlParserENCLOSED                      = 51
	SqlParserESCAPED                       = 52
	SqlParserEXISTS                        = 53
	SqlParserEXIT                          = 54
	SqlParserEXPLAIN                       = 55
	SqlParserFALSE                         = 56
	SqlParserFETCH                         = 57
	SqlParserFOR                           = 58
	SqlParserFORCE                         = 59
	SqlParserFOREIGN                       = 60
	SqlParserFROM                          = 61
	SqlParserFULLTEXT                      = 62
	SqlParserGENERATED                     = 63
	SqlParserGET                           = 64
	SqlParserGRANT                         = 65
	SqlParserGROUP                         = 66
	SqlParserHAVING                        = 67
	SqlParserHIGH_PRIORITY                 = 68
	SqlParserIF                            = 69
	SqlParserIGNORE                        = 70
	SqlParserIN                            = 71
	SqlParserINDEX                         = 72
	SqlParserINFILE                        = 73
	SqlParserINNER                         = 74
	SqlParserINOUT                         = 75
	SqlParserINSERT                        = 76
	SqlParserINTERVAL                      = 77
	SqlParserINTO                          = 78
	SqlParserIS                            = 79
	SqlParserITERATE                       = 80
	SqlParserJOIN                          = 81
	SqlParserKEY                           = 82
	SqlParserKEYS                          = 83
	SqlParserKILL                          = 84
	SqlParserLEADING                       = 85
	SqlParserLEAVE                         = 86
	SqlParserLEFT                          = 87
	SqlParserLIKE                          = 88
	SqlParserLIMIT                         = 89
	SqlParserLINEAR                        = 90
	SqlParserLINES                         = 91
	SqlParserLOAD                          = 92
	SqlParserLOCK                          = 93
	SqlParserLOOP                          = 94
	SqlParserLOW_PRIORITY                  = 95
	SqlParserMASTER_BIND                   = 96
	SqlParserMASTER_SSL_VERIFY_SERVER_CERT = 97
	SqlParserMATCH                         = 98
	SqlParserMAXVALUE                      = 99
	SqlParserMODIFIES                      = 100
	SqlParserNATURAL                       = 101
	SqlParserNOT                           = 102
	SqlParserNO_WRITE_TO_BINLOG            = 103
	SqlParserNULL_LITERAL                  = 104
	SqlParserNUMBER                        = 105
	SqlParserON                            = 106
	SqlParserOPTIMIZE                      = 107
	SqlParserOPTION                        = 108
	SqlParserOPTIONALLY                    = 109
	SqlParserOR                            = 110
	SqlParserORDER                         = 111
	SqlParserOUT                           = 112
	SqlParserOUTER                         = 113
	SqlParserOUTFILE                       = 114
	SqlParserPARTITION                     = 115
	SqlParserPRIMARY                       = 116
	SqlParserPROCEDURE                     = 117
	SqlParserPURGE                         = 118
	SqlParserRANGE                         = 119
	SqlParserREAD                          = 120
	SqlParserREADS                         = 121
	SqlParserREFERENCES                    = 122
	SqlParserREGEXP                        = 123
	SqlParserRELEASE                       = 124
	SqlParserRENAME                        = 125
	SqlParserREPEAT                        = 126
	SqlParserREPLACE                       = 127
	SqlParserREQUIRE                       = 128
	SqlParserRESIGNAL                      = 129
	SqlParserRESTRICT                      = 130
	SqlParserRETURN                        = 131
	SqlParserREVOKE                        = 132
	SqlParserRIGHT                         = 133
	SqlParserRLIKE                         = 134
	SqlParserSCHEMA                        = 135
	SqlParserSCHEMAS                       = 136
	SqlParserSELECT                        = 137
	SqlParserSET                           = 138
	SqlParserSEPARATOR                     = 139
	SqlParserSHOW                          = 140
	SqlParserSIGNAL                        = 141
	SqlParserSPATIAL                       = 142
	SqlParserSQL                           = 143
	SqlParserSQLEXCEPTION                  = 144
	SqlParserSQLSTATE                      = 145
	SqlParserSQLWARNING                    = 146
	SqlParserSQL_BIG_RESULT                = 147
	SqlParserSQL_CALC_FOUND_ROWS           = 148
	SqlParserSQL_SMALL_RESULT              = 149
	SqlParserSSL                           = 150
	SqlParserSTACKED                       = 151
	SqlParserSTARTING                      = 152
	SqlParserSTRAIGHT_JOIN                 = 153
	SqlParserTABLE                         = 154
	SqlParserTERMINATED                    = 155
	SqlParserTHEN                          = 156
	SqlParserTO                            = 157
	SqlParserTRAILING                      = 158
	SqlParserTRIGGER                       = 159
	SqlParserTRUE                          = 160
	SqlParserUNDO                          = 161
	SqlParserUNION                         = 162
	SqlParserUNIQUE                        = 163
	SqlParserUNLOCK                        = 164
	SqlParserUNSIGNED                      = 165
	SqlParserUPDATE                        = 166
	SqlParserUSAGE                         = 167
	SqlParserUSE                           = 168
	SqlParserUSING                         = 169
	SqlParserVALUES                        = 170
	SqlParserWHEN                          = 171
	SqlParserWHERE                         = 172
	SqlParserWHILE                         = 173
	SqlParserWITH                          = 174
	SqlParserWRITE                         = 175
	SqlParserXOR                           = 176
	SqlParserZEROFILL                      = 177
	SqlParserAVG                           = 178
	SqlParserCOUNT                         = 179
	SqlParserMAX                           = 180
	SqlParserMIN                           = 181
	SqlParserSUM                           = 182
	SqlParserOFFSET                        = 183
	SqlParserESCAPE                        = 184
	SqlParserVAR_ASSIGN                    = 185
	SqlParserPLUS_ASSIGN                   = 186
	SqlParserMINUS_ASSIGN                  = 187
	SqlParserMULT_ASSIGN                   = 188
	SqlParserDIV_ASSIGN                    = 189
	SqlParserMOD_ASSIGN                    = 190
	SqlParserAND_ASSIGN                    = 191
	SqlParserXOR_ASSIGN                    = 192
	SqlParserOR_ASSIGN                     = 193
	SqlParserSTAR                          = 194
	SqlParserDIVIDE                        = 195
	SqlParserMODULE                        = 196
	SqlParserPLUS                          = 197
	SqlParserMINUSMINUS                    = 198
	SqlParserMINUS                         = 199
	SqlParserDIV                           = 200
	SqlParserMOD                           = 201
	SqlParserEQUAL_SYMBOL                  = 202
	SqlParserGREATER_SYMBOL                = 203
	SqlParserLESS_SYMBOL                   = 204
	SqlParserEXCLAMATION_SYMBOL            = 205
	SqlParserBIT_NOT_OP                    = 206
	SqlParserBIT_OR_OP                     = 207
	SqlParserBIT_AND_OP                    = 208
	SqlParserBIT_XOR_OP                    = 209
	SqlParserDOT                           = 210
	SqlParserLR_BRACKET                    = 211
	SqlParserRR_BRACKET                    = 212
	SqlParserCOMMA                         = 213
	SqlParserSEMI                          = 214
	SqlParserAT_SIGN                       = 215
	SqlParserZERO_DECIMAL                  = 216
	SqlParserONE_DECIMAL                   = 217
	SqlParserTWO_DECIMAL                   = 218
	SqlParserSINGLE_QUOTE_SYMB             = 219
	SqlParserDOUBLE_QUOTE_SYMB             = 220
	SqlParserREVERSE_QUOTE_SYMB            = 221
	SqlParserCOLON_SYMB                    = 222
	SqlParserFILESIZE_LITERAL              = 223
	SqlParserSTART_NATIONAL_STRING_LITERAL = 224
	SqlParserSTRING_LITERAL                = 225
	SqlParserDECIMAL_LITERAL               = 226
	SqlParserHEXADECIMAL_LITERAL           = 227
	SqlParserREAL_LITERAL                  = 228
	SqlParserNULL_SPEC_LITERAL             = 229
	SqlParserDOT_ID                        = 230
	SqlParserID                            = 231
	SqlParserREVERSE_QUOTE_ID              = 232
	SqlParserSTRING_USER_NAME              = 233
	SqlParserLOCAL_ID                      = 234
	SqlParserGLOBAL_ID                     = 235
	SqlParserID_LITERAL                    = 236
	SqlParserERROR_RECONGNIGION            = 237
)

// SqlParser rules.
const (
	SqlParserRULE_selectStatement           = 0
	SqlParserRULE_orderByClause             = 1
	SqlParserRULE_orderByExpression         = 2
	SqlParserRULE_selectElements            = 3
	SqlParserRULE_selectElement             = 4
	SqlParserRULE_whereClause               = 5
	SqlParserRULE_groupbyClause             = 6
	SqlParserRULE_havingClause              = 7
	SqlParserRULE_limitClause               = 8
	SqlParserRULE_limitClauseAtom           = 9
	SqlParserRULE_columnName                = 10
	SqlParserRULE_uid                       = 11
	SqlParserRULE_decimalLiteral            = 12
	SqlParserRULE_stringLiteral             = 13
	SqlParserRULE_booleanLiteral            = 14
	SqlParserRULE_constant                  = 15
	SqlParserRULE_functionCall              = 16
	SqlParserRULE_aggregateWindowedFunction = 17
	SqlParserRULE_functionArgs              = 18
	SqlParserRULE_expression                = 19
	SqlParserRULE_predicate                 = 20
	SqlParserRULE_expressionAtom            = 21
	SqlParserRULE_unaryOperator             = 22
	SqlParserRULE_comparisonOperator        = 23
	SqlParserRULE_logicalOperator           = 24
	SqlParserRULE_mathOperator              = 25
)

// ISelectStatementContext is an interface to support dynamic dispatch.
type ISelectStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTableName returns the tableName rule contexts.
	GetTableName() IUidContext

	// SetTableName sets the tableName rule contexts.
	SetTableName(IUidContext)

	// IsSelectStatementContext differentiates from other interfaces.
	IsSelectStatementContext()
}

type SelectStatementContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	tableName IUidContext
}

func NewEmptySelectStatementContext() *SelectStatementContext {
	var p = new(SelectStatementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectStatement
	return p
}

func (*SelectStatementContext) IsSelectStatementContext() {}

func NewSelectStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectStatementContext {
	var p = new(SelectStatementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectStatement

	return p
}

func (s *SelectStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectStatementContext) GetTableName() IUidContext { return s.tableName }

func (s *SelectStatementContext) SetTableName(v IUidContext) { s.tableName = v }

func (s *SelectStatementContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SqlParserSELECT, 0)
}

func (s *SelectStatementContext) SelectElements() ISelectElementsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISelectElementsContext)
}

func (s *SelectStatementContext) FROM() antlr.TerminalNode {
	return s.GetToken(SqlParserFROM, 0)
}

func (s *SelectStatementContext) EOF() antlr.TerminalNode {
	return s.GetToken(SqlParserEOF, 0)
}

func (s *SelectStatementContext) Uid() IUidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *SelectStatementContext) WhereClause() IWhereClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IWhereClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectStatementContext) GroupbyClause() IGroupbyClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IGroupbyClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IGroupbyClauseContext)
}

func (s *SelectStatementContext) HavingClause() IHavingClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IHavingClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IHavingClauseContext)
}

func (s *SelectStatementContext) OrderByClause() IOrderByClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IOrderByClauseContext)
}

func (s *SelectStatementContext) LimitClause() ILimitClauseContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *SelectStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectStatement(s)
	}
}

func (s *SelectStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectStatement(s)
	}
}

func (s *SelectStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectStatement() (localctx ISelectStatementContext) {
	localctx = NewSelectStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SqlParserRULE_selectStatement)
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
		p.SetState(52)
		p.Match(SqlParserSELECT)
	}
	{
		p.SetState(53)
		p.SelectElements()
	}
	{
		p.SetState(54)
		p.Match(SqlParserFROM)
	}
	{
		p.SetState(55)

		var _x = p.Uid()

		localctx.(*SelectStatementContext).tableName = _x
	}
	p.SetState(57)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserWHERE {
		{
			p.SetState(56)
			p.WhereClause()
		}

	}
	p.SetState(60)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserGROUP {
		{
			p.SetState(59)
			p.GroupbyClause()
		}

	}
	p.SetState(63)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserHAVING {
		{
			p.SetState(62)
			p.HavingClause()
		}

	}
	p.SetState(66)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserORDER {
		{
			p.SetState(65)
			p.OrderByClause()
		}

	}
	p.SetState(69)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserLIMIT {
		{
			p.SetState(68)
			p.LimitClause()
		}

	}
	{
		p.SetState(71)
		p.Match(SqlParserEOF)
	}

	return localctx
}

// IOrderByClauseContext is an interface to support dynamic dispatch.
type IOrderByClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsOrderByClauseContext differentiates from other interfaces.
	IsOrderByClauseContext()
}

type OrderByClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderByClauseContext() *OrderByClauseContext {
	var p = new(OrderByClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByClause
	return p
}

func (*OrderByClauseContext) IsOrderByClauseContext() {}

func NewOrderByClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByClauseContext {
	var p = new(OrderByClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByClause

	return p
}

func (s *OrderByClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SqlParserORDER, 0)
}

func (s *OrderByClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *OrderByClauseContext) AllOrderByExpression() []IOrderByExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IOrderByExpressionContext)(nil)).Elem())
	var tst = make([]IOrderByExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IOrderByExpressionContext)
		}
	}

	return tst
}

func (s *OrderByClauseContext) OrderByExpression(i int) IOrderByExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IOrderByExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IOrderByExpressionContext)
}

func (s *OrderByClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterOrderByClause(s)
	}
}

func (s *OrderByClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitOrderByClause(s)
	}
}

func (s *OrderByClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitOrderByClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OrderByClause() (localctx IOrderByClauseContext) {
	localctx = NewOrderByClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SqlParserRULE_orderByClause)
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
		p.SetState(73)
		p.Match(SqlParserORDER)
	}
	{
		p.SetState(74)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(75)
		p.OrderByExpression()
	}
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(76)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(77)
			p.OrderByExpression()
		}

		p.SetState(82)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IOrderByExpressionContext is an interface to support dynamic dispatch.
type IOrderByExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOrder returns the order token.
	GetOrder() antlr.Token

	// SetOrder sets the order token.
	SetOrder(antlr.Token)

	// IsOrderByExpressionContext differentiates from other interfaces.
	IsOrderByExpressionContext()
}

type OrderByExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	order  antlr.Token
}

func NewEmptyOrderByExpressionContext() *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_orderByExpression
	return p
}

func (*OrderByExpressionContext) IsOrderByExpressionContext() {}

func NewOrderByExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderByExpressionContext {
	var p = new(OrderByExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_orderByExpression

	return p
}

func (s *OrderByExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderByExpressionContext) GetOrder() antlr.Token { return s.order }

func (s *OrderByExpressionContext) SetOrder(v antlr.Token) { s.order = v }

func (s *OrderByExpressionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *OrderByExpressionContext) ASC() antlr.TerminalNode {
	return s.GetToken(SqlParserASC, 0)
}

func (s *OrderByExpressionContext) DESC() antlr.TerminalNode {
	return s.GetToken(SqlParserDESC, 0)
}

func (s *OrderByExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderByExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderByExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitOrderByExpression(s)
	}
}

func (s *OrderByExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitOrderByExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) OrderByExpression() (localctx IOrderByExpressionContext) {
	localctx = NewOrderByExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SqlParserRULE_orderByExpression)
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
		p.SetState(83)
		p.expression(0)
	}
	p.SetState(85)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	if _la == SqlParserASC || _la == SqlParserDESC {
		p.SetState(84)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*OrderByExpressionContext).order = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserASC || _la == SqlParserDESC) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*OrderByExpressionContext).order = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}

	}

	return localctx
}

// ISelectElementsContext is an interface to support dynamic dispatch.
type ISelectElementsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectElementsContext differentiates from other interfaces.
	IsSelectElementsContext()
}

type SelectElementsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementsContext() *SelectElementsContext {
	var p = new(SelectElementsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectElements
	return p
}

func (*SelectElementsContext) IsSelectElementsContext() {}

func NewSelectElementsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementsContext {
	var p = new(SelectElementsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElements

	return p
}

func (s *SelectElementsContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementsContext) AllSelectElement() []ISelectElementContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ISelectElementContext)(nil)).Elem())
	var tst = make([]ISelectElementContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ISelectElementContext)
		}
	}

	return tst
}

func (s *SelectElementsContext) SelectElement(i int) ISelectElementContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISelectElementContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ISelectElementContext)
}

func (s *SelectElementsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectElementsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectElements(s)
	}
}

func (s *SelectElementsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectElements(s)
	}
}

func (s *SelectElementsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectElements(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectElements() (localctx ISelectElementsContext) {
	localctx = NewSelectElementsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SqlParserRULE_selectElements)
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
		p.SetState(87)
		p.SelectElement()
	}
	p.SetState(92)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(88)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(89)
			p.SelectElement()
		}

		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// ISelectElementContext is an interface to support dynamic dispatch.
type ISelectElementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsSelectElementContext differentiates from other interfaces.
	IsSelectElementContext()
}

type SelectElementContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectElementContext() *SelectElementContext {
	var p = new(SelectElementContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_selectElement
	return p
}

func (*SelectElementContext) IsSelectElementContext() {}

func NewSelectElementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectElementContext {
	var p = new(SelectElementContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_selectElement

	return p
}

func (s *SelectElementContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectElementContext) CopyFrom(ctx *SelectElementContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *SelectElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectElementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type SelectExprElementContext struct {
	*SelectElementContext
}

func NewSelectExprElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectExprElementContext {
	var p = new(SelectExprElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectExprElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectExprElementContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *SelectExprElementContext) AS() antlr.TerminalNode {
	return s.GetToken(SqlParserAS, 0)
}

func (s *SelectExprElementContext) Uid() IUidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *SelectExprElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectExprElement(s)
	}
}

func (s *SelectExprElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectExprElement(s)
	}
}

func (s *SelectExprElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectExprElement(s)

	default:
		return t.VisitChildren(s)
	}
}

type SelectStarElementContext struct {
	*SelectElementContext
}

func NewSelectStarElementContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *SelectStarElementContext {
	var p = new(SelectStarElementContext)

	p.SelectElementContext = NewEmptySelectElementContext()
	p.parser = parser
	p.CopyFrom(ctx.(*SelectElementContext))

	return p
}

func (s *SelectStarElementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectStarElementContext) Uid() IUidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *SelectStarElementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterSelectStarElement(s)
	}
}

func (s *SelectStarElementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitSelectStarElement(s)
	}
}

func (s *SelectStarElementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitSelectStarElement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) SelectElement() (localctx ISelectElementContext) {
	localctx = NewSelectElementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SqlParserRULE_selectElement)
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

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext()) {
	case 1:
		localctx = NewSelectStarElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserID {
			{
				p.SetState(95)
				p.Uid()
			}
			{
				p.SetState(96)
				p.Match(SqlParserDOT)
			}

		}
		{
			p.SetState(100)
			p.Match(SqlParserSTAR)
		}

	case 2:
		localctx = NewSelectExprElementContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(101)
			p.expressionAtom(0)
		}
		p.SetState(104)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserAS {
			{
				p.SetState(102)
				p.Match(SqlParserAS)
			}
			{
				p.SetState(103)
				p.Uid()
			}

		}

	}

	return localctx
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_whereClause
	return p
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SqlParserWHERE, 0)
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
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SqlParserRULE_whereClause)

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
		p.SetState(108)
		p.Match(SqlParserWHERE)
	}
	{
		p.SetState(109)
		p.expression(0)
	}

	return localctx
}

// IGroupbyClauseContext is an interface to support dynamic dispatch.
type IGroupbyClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsGroupbyClauseContext differentiates from other interfaces.
	IsGroupbyClauseContext()
}

type GroupbyClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupbyClauseContext() *GroupbyClauseContext {
	var p = new(GroupbyClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_groupbyClause
	return p
}

func (*GroupbyClauseContext) IsGroupbyClauseContext() {}

func NewGroupbyClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupbyClauseContext {
	var p = new(GroupbyClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_groupbyClause

	return p
}

func (s *GroupbyClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupbyClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SqlParserGROUP, 0)
}

func (s *GroupbyClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SqlParserBY, 0)
}

func (s *GroupbyClauseContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *GroupbyClauseContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GroupbyClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupbyClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupbyClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterGroupbyClause(s)
	}
}

func (s *GroupbyClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitGroupbyClause(s)
	}
}

func (s *GroupbyClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitGroupbyClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) GroupbyClause() (localctx IGroupbyClauseContext) {
	localctx = NewGroupbyClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SqlParserRULE_groupbyClause)
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
		p.SetState(111)
		p.Match(SqlParserGROUP)
	}
	{
		p.SetState(112)
		p.Match(SqlParserBY)
	}
	{
		p.SetState(113)
		p.expression(0)
	}
	p.SetState(118)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(114)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(115)
			p.expression(0)
		}

		p.SetState(120)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	return localctx
}

// IHavingClauseContext is an interface to support dynamic dispatch.
type IHavingClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetHavingExpr returns the havingExpr rule contexts.
	GetHavingExpr() IExpressionContext

	// SetHavingExpr sets the havingExpr rule contexts.
	SetHavingExpr(IExpressionContext)

	// IsHavingClauseContext differentiates from other interfaces.
	IsHavingClauseContext()
}

type HavingClauseContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	havingExpr IExpressionContext
}

func NewEmptyHavingClauseContext() *HavingClauseContext {
	var p = new(HavingClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_havingClause
	return p
}

func (*HavingClauseContext) IsHavingClauseContext() {}

func NewHavingClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingClauseContext {
	var p = new(HavingClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_havingClause

	return p
}

func (s *HavingClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingClauseContext) GetHavingExpr() IExpressionContext { return s.havingExpr }

func (s *HavingClauseContext) SetHavingExpr(v IExpressionContext) { s.havingExpr = v }

func (s *HavingClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SqlParserHAVING, 0)
}

func (s *HavingClauseContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *HavingClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterHavingClause(s)
	}
}

func (s *HavingClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitHavingClause(s)
	}
}

func (s *HavingClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitHavingClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) HavingClause() (localctx IHavingClauseContext) {
	localctx = NewHavingClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SqlParserRULE_havingClause)

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
		p.SetState(121)
		p.Match(SqlParserHAVING)
	}
	{
		p.SetState(122)

		var _x = p.expression(0)

		localctx.(*HavingClauseContext).havingExpr = _x
	}

	return localctx
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetOffset returns the offset rule contexts.
	GetOffset() ILimitClauseAtomContext

	// GetLimit returns the limit rule contexts.
	GetLimit() ILimitClauseAtomContext

	// SetOffset sets the offset rule contexts.
	SetOffset(ILimitClauseAtomContext)

	// SetLimit sets the limit rule contexts.
	SetLimit(ILimitClauseAtomContext)

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	offset ILimitClauseAtomContext
	limit  ILimitClauseAtomContext
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_limitClause
	return p
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) GetOffset() ILimitClauseAtomContext { return s.offset }

func (s *LimitClauseContext) GetLimit() ILimitClauseAtomContext { return s.limit }

func (s *LimitClauseContext) SetOffset(v ILimitClauseAtomContext) { s.offset = v }

func (s *LimitClauseContext) SetLimit(v ILimitClauseAtomContext) { s.limit = v }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SqlParserLIMIT, 0)
}

func (s *LimitClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SqlParserOFFSET, 0)
}

func (s *LimitClauseContext) AllLimitClauseAtom() []ILimitClauseAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*ILimitClauseAtomContext)(nil)).Elem())
	var tst = make([]ILimitClauseAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(ILimitClauseAtomContext)
		}
	}

	return tst
}

func (s *LimitClauseContext) LimitClauseAtom(i int) ILimitClauseAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILimitClauseAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(ILimitClauseAtomContext)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SqlParserRULE_limitClause)

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
		p.SetState(124)
		p.Match(SqlParserLIMIT)
	}
	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.SetState(128)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(125)

				var _x = p.LimitClauseAtom()

				localctx.(*LimitClauseContext).offset = _x
			}
			{
				p.SetState(126)
				p.Match(SqlParserCOMMA)
			}

		}
		{
			p.SetState(130)

			var _x = p.LimitClauseAtom()

			localctx.(*LimitClauseContext).limit = _x
		}

	case 2:
		{
			p.SetState(131)

			var _x = p.LimitClauseAtom()

			localctx.(*LimitClauseContext).limit = _x
		}
		{
			p.SetState(132)
			p.Match(SqlParserOFFSET)
		}
		{
			p.SetState(133)

			var _x = p.LimitClauseAtom()

			localctx.(*LimitClauseContext).offset = _x
		}

	}

	return localctx
}

// ILimitClauseAtomContext is an interface to support dynamic dispatch.
type ILimitClauseAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLimitClauseAtomContext differentiates from other interfaces.
	IsLimitClauseAtomContext()
}

type LimitClauseAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseAtomContext() *LimitClauseAtomContext {
	var p = new(LimitClauseAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_limitClauseAtom
	return p
}

func (*LimitClauseAtomContext) IsLimitClauseAtomContext() {}

func NewLimitClauseAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseAtomContext {
	var p = new(LimitClauseAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_limitClauseAtom

	return p
}

func (s *LimitClauseAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseAtomContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *LimitClauseAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLimitClauseAtom(s)
	}
}

func (s *LimitClauseAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLimitClauseAtom(s)
	}
}

func (s *LimitClauseAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLimitClauseAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LimitClauseAtom() (localctx ILimitClauseAtomContext) {
	localctx = NewLimitClauseAtomContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SqlParserRULE_limitClauseAtom)

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
		p.SetState(137)
		p.DecimalLiteral()
	}

	return localctx
}

// IColumnNameContext is an interface to support dynamic dispatch.
type IColumnNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetTn returns the tn rule contexts.
	GetTn() IUidContext

	// GetCn returns the cn rule contexts.
	GetCn() IUidContext

	// SetTn sets the tn rule contexts.
	SetTn(IUidContext)

	// SetCn sets the cn rule contexts.
	SetCn(IUidContext)

	// IsColumnNameContext differentiates from other interfaces.
	IsColumnNameContext()
}

type ColumnNameContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	tn     IUidContext
	cn     IUidContext
}

func NewEmptyColumnNameContext() *ColumnNameContext {
	var p = new(ColumnNameContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_columnName
	return p
}

func (*ColumnNameContext) IsColumnNameContext() {}

func NewColumnNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ColumnNameContext {
	var p = new(ColumnNameContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_columnName

	return p
}

func (s *ColumnNameContext) GetParser() antlr.Parser { return s.parser }

func (s *ColumnNameContext) GetTn() IUidContext { return s.tn }

func (s *ColumnNameContext) GetCn() IUidContext { return s.cn }

func (s *ColumnNameContext) SetTn(v IUidContext) { s.tn = v }

func (s *ColumnNameContext) SetCn(v IUidContext) { s.cn = v }

func (s *ColumnNameContext) AllUid() []IUidContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IUidContext)(nil)).Elem())
	var tst = make([]IUidContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IUidContext)
		}
	}

	return tst
}

func (s *ColumnNameContext) Uid(i int) IUidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *ColumnNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ColumnNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnName(s)
	}
}

func (s *ColumnNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnName(s)
	}
}

func (s *ColumnNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ColumnName() (localctx IColumnNameContext) {
	localctx = NewColumnNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SqlParserRULE_columnName)

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
	p.SetState(142)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 14, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(139)

			var _x = p.Uid()

			localctx.(*ColumnNameContext).tn = _x
		}
		{
			p.SetState(140)
			p.Match(SqlParserDOT)
		}

	}
	{
		p.SetState(144)

		var _x = p.Uid()

		localctx.(*ColumnNameContext).cn = _x
	}

	return localctx
}

// IUidContext is an interface to support dynamic dispatch.
type IUidContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUidContext differentiates from other interfaces.
	IsUidContext()
}

type UidContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUidContext() *UidContext {
	var p = new(UidContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_uid
	return p
}

func (*UidContext) IsUidContext() {}

func NewUidContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UidContext {
	var p = new(UidContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_uid

	return p
}

func (s *UidContext) GetParser() antlr.Parser { return s.parser }

func (s *UidContext) ID() antlr.TerminalNode {
	return s.GetToken(SqlParserID, 0)
}

func (s *UidContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UidContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UidContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUid(s)
	}
}

func (s *UidContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUid(s)
	}
}

func (s *UidContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUid(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Uid() (localctx IUidContext) {
	localctx = NewUidContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SqlParserRULE_uid)

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
		p.SetState(146)
		p.Match(SqlParserID)
	}

	return localctx
}

// IDecimalLiteralContext is an interface to support dynamic dispatch.
type IDecimalLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsDecimalLiteralContext differentiates from other interfaces.
	IsDecimalLiteralContext()
}

type DecimalLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDecimalLiteralContext() *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_decimalLiteral
	return p
}

func (*DecimalLiteralContext) IsDecimalLiteralContext() {}

func NewDecimalLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DecimalLiteralContext {
	var p = new(DecimalLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_decimalLiteral

	return p
}

func (s *DecimalLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *DecimalLiteralContext) DECIMAL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserDECIMAL_LITERAL, 0)
}

func (s *DecimalLiteralContext) ZERO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserZERO_DECIMAL, 0)
}

func (s *DecimalLiteralContext) ONE_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserONE_DECIMAL, 0)
}

func (s *DecimalLiteralContext) TWO_DECIMAL() antlr.TerminalNode {
	return s.GetToken(SqlParserTWO_DECIMAL, 0)
}

func (s *DecimalLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DecimalLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DecimalLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitDecimalLiteral(s)
	}
}

func (s *DecimalLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitDecimalLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) DecimalLiteral() (localctx IDecimalLiteralContext) {
	localctx = NewDecimalLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SqlParserRULE_decimalLiteral)
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
	p.SetState(148)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-216)&-(0x1f+1)) == 0 && ((1<<uint((_la-216)))&((1<<(SqlParserZERO_DECIMAL-216))|(1<<(SqlParserONE_DECIMAL-216))|(1<<(SqlParserTWO_DECIMAL-216))|(1<<(SqlParserDECIMAL_LITERAL-216)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IStringLiteralContext is an interface to support dynamic dispatch.
type IStringLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsStringLiteralContext differentiates from other interfaces.
	IsStringLiteralContext()
}

type StringLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStringLiteralContext() *StringLiteralContext {
	var p = new(StringLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_stringLiteral
	return p
}

func (*StringLiteralContext) IsStringLiteralContext() {}

func NewStringLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StringLiteralContext {
	var p = new(StringLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_stringLiteral

	return p
}

func (s *StringLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *StringLiteralContext) AllSTRING_LITERAL() []antlr.TerminalNode {
	return s.GetTokens(SqlParserSTRING_LITERAL)
}

func (s *StringLiteralContext) STRING_LITERAL(i int) antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING_LITERAL, i)
}

func (s *StringLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StringLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterStringLiteral(s)
	}
}

func (s *StringLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitStringLiteral(s)
	}
}

func (s *StringLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitStringLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) StringLiteral() (localctx IStringLiteralContext) {
	localctx = NewStringLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SqlParserRULE_stringLiteral)

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
	p.SetState(151)
	p.GetErrorHandler().Sync(p)
	_alt = 1
	for ok := true; ok; ok = _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		switch _alt {
		case 1:
			{
				p.SetState(150)
				p.Match(SqlParserSTRING_LITERAL)
			}

		default:
			panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		}

		p.SetState(153)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext())
	}

	return localctx
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_booleanLiteral
	return p
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SqlParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SqlParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SqlParserRULE_booleanLiteral)
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
	p.SetState(155)
	_la = p.GetTokenStream().LA(1)

	if !(_la == SqlParserFALSE || _la == SqlParserTRUE) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IConstantContext is an interface to support dynamic dispatch.
type IConstantContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsConstantContext differentiates from other interfaces.
	IsConstantContext()
}

type ConstantContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstantContext() *ConstantContext {
	var p = new(ConstantContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_constant
	return p
}

func (*ConstantContext) IsConstantContext() {}

func NewConstantContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstantContext {
	var p = new(ConstantContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_constant

	return p
}

func (s *ConstantContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstantContext) CopyFrom(ctx *ConstantContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type StringConstantContext struct {
	*ConstantContext
}

func NewStringConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *StringConstantContext {
	var p = new(StringConstantContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *StringConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StringConstantContext) StringLiteral() IStringLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IStringLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IStringLiteralContext)
}

func (s *StringConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterStringConstant(s)
	}
}

func (s *StringConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitStringConstant(s)
	}
}

func (s *StringConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitStringConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type BoolConstantContext struct {
	*ConstantContext
}

func NewBoolConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BoolConstantContext {
	var p = new(BoolConstantContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *BoolConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BoolConstantContext) BooleanLiteral() IBooleanLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBooleanLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *BoolConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBoolConstant(s)
	}
}

func (s *BoolConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBoolConstant(s)
	}
}

func (s *BoolConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBoolConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumberConstantContext struct {
	*ConstantContext
	minusFlag antlr.Token
}

func NewNumberConstantContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumberConstantContext {
	var p = new(NumberConstantContext)

	p.ConstantContext = NewEmptyConstantContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ConstantContext))

	return p
}

func (s *NumberConstantContext) GetMinusFlag() antlr.Token { return s.minusFlag }

func (s *NumberConstantContext) SetMinusFlag(v antlr.Token) { s.minusFlag = v }

func (s *NumberConstantContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumberConstantContext) DecimalLiteral() IDecimalLiteralContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDecimalLiteralContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDecimalLiteralContext)
}

func (s *NumberConstantContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterNumberConstant(s)
	}
}

func (s *NumberConstantContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitNumberConstant(s)
	}
}

func (s *NumberConstantContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitNumberConstant(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Constant() (localctx IConstantContext) {
	localctx = NewConstantContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SqlParserRULE_constant)
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

	p.SetState(163)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserSTRING_LITERAL:
		localctx = NewStringConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(157)
			p.StringLiteral()
		}

	case SqlParserMINUS, SqlParserZERO_DECIMAL, SqlParserONE_DECIMAL, SqlParserTWO_DECIMAL, SqlParserDECIMAL_LITERAL:
		localctx = NewNumberConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		p.SetState(159)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserMINUS {
			{
				p.SetState(158)

				var _m = p.Match(SqlParserMINUS)

				localctx.(*NumberConstantContext).minusFlag = _m
			}

		}
		{
			p.SetState(161)
			p.DecimalLiteral()
		}

	case SqlParserFALSE, SqlParserTRUE:
		localctx = NewBoolConstantContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(162)
			p.BooleanLiteral()
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionCall
	return p
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) CopyFrom(ctx *FunctionCallContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UdfFunctionCallContext struct {
	*FunctionCallContext
}

func NewUdfFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UdfFunctionCallContext {
	var p = new(UdfFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *UdfFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UdfFunctionCallContext) Uid() IUidContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUidContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUidContext)
}

func (s *UdfFunctionCallContext) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *UdfFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUdfFunctionCall(s)
	}
}

func (s *UdfFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUdfFunctionCall(s)
	}
}

func (s *UdfFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUdfFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

type AggregateFunctionCallContext struct {
	*FunctionCallContext
}

func NewAggregateFunctionCallContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AggregateFunctionCallContext {
	var p = new(AggregateFunctionCallContext)

	p.FunctionCallContext = NewEmptyFunctionCallContext()
	p.parser = parser
	p.CopyFrom(ctx.(*FunctionCallContext))

	return p
}

func (s *AggregateFunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunctionCallContext) AggregateWindowedFunction() IAggregateWindowedFunctionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAggregateWindowedFunctionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAggregateWindowedFunctionContext)
}

func (s *AggregateFunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterAggregateFunctionCall(s)
	}
}

func (s *AggregateFunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitAggregateFunctionCall(s)
	}
}

func (s *AggregateFunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitAggregateFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SqlParserRULE_functionCall)
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

	p.SetState(173)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserAVG, SqlParserCOUNT, SqlParserMAX, SqlParserMIN, SqlParserSUM:
		localctx = NewAggregateFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(165)
			p.AggregateWindowedFunction()
		}

	case SqlParserID:
		localctx = NewUdfFunctionCallContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(166)
			p.Uid()
		}
		{
			p.SetState(167)
			p.Match(SqlParserLR_BRACKET)
		}
		p.SetState(169)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserFALSE || _la == SqlParserNOT || (((_la-160)&-(0x1f+1)) == 0 && ((1<<uint((_la-160)))&((1<<(SqlParserTRUE-160))|(1<<(SqlParserAVG-160))|(1<<(SqlParserCOUNT-160))|(1<<(SqlParserMAX-160))|(1<<(SqlParserMIN-160))|(1<<(SqlParserSUM-160)))) != 0) || (((_la-197)&-(0x1f+1)) == 0 && ((1<<uint((_la-197)))&((1<<(SqlParserPLUS-197))|(1<<(SqlParserMINUS-197))|(1<<(SqlParserEXCLAMATION_SYMBOL-197))|(1<<(SqlParserBIT_NOT_OP-197))|(1<<(SqlParserLR_BRACKET-197))|(1<<(SqlParserZERO_DECIMAL-197))|(1<<(SqlParserONE_DECIMAL-197))|(1<<(SqlParserTWO_DECIMAL-197))|(1<<(SqlParserSTRING_LITERAL-197))|(1<<(SqlParserDECIMAL_LITERAL-197)))) != 0) || _la == SqlParserID {
			{
				p.SetState(168)
				p.FunctionArgs()
			}

		}
		{
			p.SetState(171)
			p.Match(SqlParserRR_BRACKET)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAggregateWindowedFunctionContext is an interface to support dynamic dispatch.
type IAggregateWindowedFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsAggregateWindowedFunctionContext differentiates from other interfaces.
	IsAggregateWindowedFunctionContext()
}

type AggregateWindowedFunctionContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateWindowedFunctionContext() *AggregateWindowedFunctionContext {
	var p = new(AggregateWindowedFunctionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_aggregateWindowedFunction
	return p
}

func (*AggregateWindowedFunctionContext) IsAggregateWindowedFunctionContext() {}

func NewAggregateWindowedFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateWindowedFunctionContext {
	var p = new(AggregateWindowedFunctionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_aggregateWindowedFunction

	return p
}

func (s *AggregateWindowedFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateWindowedFunctionContext) CopyFrom(ctx *AggregateWindowedFunctionContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AggregateWindowedFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateWindowedFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type CountFunc2Context struct {
	*AggregateWindowedFunctionContext
	aggregator antlr.Token
}

func NewCountFunc2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountFunc2Context {
	var p = new(CountFunc2Context)

	p.AggregateWindowedFunctionContext = NewEmptyAggregateWindowedFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AggregateWindowedFunctionContext))

	return p
}

func (s *CountFunc2Context) GetAggregator() antlr.Token { return s.aggregator }

func (s *CountFunc2Context) SetAggregator(v antlr.Token) { s.aggregator = v }

func (s *CountFunc2Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountFunc2Context) COUNT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOUNT, 0)
}

func (s *CountFunc2Context) FunctionArgs() IFunctionArgsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionArgsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionArgsContext)
}

func (s *CountFunc2Context) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTINCT, 0)
}

func (s *CountFunc2Context) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *CountFunc2Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterCountFunc2(s)
	}
}

func (s *CountFunc2Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitCountFunc2(s)
	}
}

func (s *CountFunc2Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitCountFunc2(s)

	default:
		return t.VisitChildren(s)
	}
}

type CountFunc1Context struct {
	*AggregateWindowedFunctionContext
	starArg antlr.Token
}

func NewCountFunc1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *CountFunc1Context {
	var p = new(CountFunc1Context)

	p.AggregateWindowedFunctionContext = NewEmptyAggregateWindowedFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AggregateWindowedFunctionContext))

	return p
}

func (s *CountFunc1Context) GetStarArg() antlr.Token { return s.starArg }

func (s *CountFunc1Context) SetStarArg(v antlr.Token) { s.starArg = v }

func (s *CountFunc1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CountFunc1Context) COUNT() antlr.TerminalNode {
	return s.GetToken(SqlParserCOUNT, 0)
}

func (s *CountFunc1Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterCountFunc1(s)
	}
}

func (s *CountFunc1Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitCountFunc1(s)
	}
}

func (s *CountFunc1Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitCountFunc1(s)

	default:
		return t.VisitChildren(s)
	}
}

type CommonAggregateFuncContext struct {
	*AggregateWindowedFunctionContext
	aggFunc    antlr.Token
	aggregator antlr.Token
}

func NewCommonAggregateFuncContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CommonAggregateFuncContext {
	var p = new(CommonAggregateFuncContext)

	p.AggregateWindowedFunctionContext = NewEmptyAggregateWindowedFunctionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AggregateWindowedFunctionContext))

	return p
}

func (s *CommonAggregateFuncContext) GetAggFunc() antlr.Token { return s.aggFunc }

func (s *CommonAggregateFuncContext) GetAggregator() antlr.Token { return s.aggregator }

func (s *CommonAggregateFuncContext) SetAggFunc(v antlr.Token) { s.aggFunc = v }

func (s *CommonAggregateFuncContext) SetAggregator(v antlr.Token) { s.aggregator = v }

func (s *CommonAggregateFuncContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CommonAggregateFuncContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CommonAggregateFuncContext) AVG() antlr.TerminalNode {
	return s.GetToken(SqlParserAVG, 0)
}

func (s *CommonAggregateFuncContext) MAX() antlr.TerminalNode {
	return s.GetToken(SqlParserMAX, 0)
}

func (s *CommonAggregateFuncContext) MIN() antlr.TerminalNode {
	return s.GetToken(SqlParserMIN, 0)
}

func (s *CommonAggregateFuncContext) SUM() antlr.TerminalNode {
	return s.GetToken(SqlParserSUM, 0)
}

func (s *CommonAggregateFuncContext) ALL() antlr.TerminalNode {
	return s.GetToken(SqlParserALL, 0)
}

func (s *CommonAggregateFuncContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SqlParserDISTINCT, 0)
}

func (s *CommonAggregateFuncContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterCommonAggregateFunc(s)
	}
}

func (s *CommonAggregateFuncContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitCommonAggregateFunc(s)
	}
}

func (s *CommonAggregateFuncContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitCommonAggregateFunc(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) AggregateWindowedFunction() (localctx IAggregateWindowedFunctionContext) {
	localctx = NewAggregateWindowedFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SqlParserRULE_aggregateWindowedFunction)
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

	p.SetState(195)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext()) {
	case 1:
		localctx = NewCommonAggregateFuncContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		p.SetState(175)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*CommonAggregateFuncContext).aggFunc = _lt

		_la = p.GetTokenStream().LA(1)

		if !(((_la-178)&-(0x1f+1)) == 0 && ((1<<uint((_la-178)))&((1<<(SqlParserAVG-178))|(1<<(SqlParserMAX-178))|(1<<(SqlParserMIN-178))|(1<<(SqlParserSUM-178)))) != 0) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*CommonAggregateFuncContext).aggFunc = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(176)
			p.Match(SqlParserLR_BRACKET)
		}
		p.SetState(178)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserALL || _la == SqlParserDISTINCT {
			p.SetState(177)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CommonAggregateFuncContext).aggregator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserALL || _la == SqlParserDISTINCT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CommonAggregateFuncContext).aggregator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		{
			p.SetState(180)
			p.expression(0)
		}
		{
			p.SetState(181)
			p.Match(SqlParserRR_BRACKET)
		}

	case 2:
		localctx = NewCountFunc1Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(183)
			p.Match(SqlParserCOUNT)
		}
		{
			p.SetState(184)
			p.Match(SqlParserLR_BRACKET)
		}
		{
			p.SetState(185)

			var _m = p.Match(SqlParserSTAR)

			localctx.(*CountFunc1Context).starArg = _m
		}
		{
			p.SetState(186)
			p.Match(SqlParserRR_BRACKET)
		}

	case 3:
		localctx = NewCountFunc2Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(187)
			p.Match(SqlParserCOUNT)
		}
		{
			p.SetState(188)
			p.Match(SqlParserLR_BRACKET)
		}
		p.SetState(190)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserALL || _la == SqlParserDISTINCT {
			p.SetState(189)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*CountFunc2Context).aggregator = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == SqlParserALL || _la == SqlParserDISTINCT) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*CountFunc2Context).aggregator = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}

		}
		{
			p.SetState(192)
			p.FunctionArgs()
		}
		{
			p.SetState(193)
			p.Match(SqlParserRR_BRACKET)
		}

	}

	return localctx
}

// IFunctionArgsContext is an interface to support dynamic dispatch.
type IFunctionArgsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsFunctionArgsContext differentiates from other interfaces.
	IsFunctionArgsContext()
}

type FunctionArgsContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionArgsContext() *FunctionArgsContext {
	var p = new(FunctionArgsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_functionArgs
	return p
}

func (*FunctionArgsContext) IsFunctionArgsContext() {}

func NewFunctionArgsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionArgsContext {
	var p = new(FunctionArgsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_functionArgs

	return p
}

func (s *FunctionArgsContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionArgsContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *FunctionArgsContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *FunctionArgsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionArgsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionArgsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitFunctionArgs(s)
	}
}

func (s *FunctionArgsContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitFunctionArgs(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) FunctionArgs() (localctx IFunctionArgsContext) {
	localctx = NewFunctionArgsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SqlParserRULE_functionArgs)
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
		p.SetState(197)
		p.expression(0)
	}
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for _la == SqlParserCOMMA {
		{
			p.SetState(198)
			p.Match(SqlParserCOMMA)
		}
		{
			p.SetState(199)
			p.expression(0)
		}

		p.SetState(204)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
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
	p.RuleIndex = SqlParserRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expression

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

type LogicalExpressionContext struct {
	*ExpressionContext
	left  IExpressionContext
	right IExpressionContext
}

func NewLogicalExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LogicalExpressionContext {
	var p = new(LogicalExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *LogicalExpressionContext) GetLeft() IExpressionContext { return s.left }

func (s *LogicalExpressionContext) GetRight() IExpressionContext { return s.right }

func (s *LogicalExpressionContext) SetLeft(v IExpressionContext) { s.left = v }

func (s *LogicalExpressionContext) SetRight(v IExpressionContext) { s.right = v }

func (s *LogicalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalExpressionContext) LogicalOperator() ILogicalOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILogicalOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILogicalOperatorContext)
}

func (s *LogicalExpressionContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *LogicalExpressionContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *LogicalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLogicalExpression(s)
	}
}

func (s *LogicalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLogicalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

type PredicateExpressionContext struct {
	*ExpressionContext
}

func NewPredicateExpressionContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PredicateExpressionContext {
	var p = new(PredicateExpressionContext)

	p.ExpressionContext = NewEmptyExpressionContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionContext))

	return p
}

func (s *PredicateExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateExpressionContext) Predicate() IPredicateContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPredicateContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPredicateContext)
}

func (s *PredicateExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitPredicateExpression(s)
	}
}

func (s *PredicateExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitPredicateExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Expression() (localctx IExpressionContext) {
	return p.expression(0)
}

func (p *SqlParser) expression(_p int) (localctx IExpressionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 38
	p.EnterRecursionRule(localctx, 38, SqlParserRULE_expression, _p)

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
	localctx = NewPredicateExpressionContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(206)
		p.Predicate()
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(214)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewLogicalExpressionContext(p, NewExpressionContext(p, _parentctx, _parentState))
			localctx.(*LogicalExpressionContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expression)
			p.SetState(208)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(209)
				p.LogicalOperator()
			}
			{
				p.SetState(210)

				var _x = p.expression(3)

				localctx.(*LogicalExpressionContext).right = _x
			}

		}
		p.SetState(216)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IPredicateContext is an interface to support dynamic dispatch.
type IPredicateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsPredicateContext differentiates from other interfaces.
	IsPredicateContext()
}

type PredicateContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPredicateContext() *PredicateContext {
	var p = new(PredicateContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_predicate
	return p
}

func (*PredicateContext) IsPredicateContext() {}

func NewPredicateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PredicateContext {
	var p = new(PredicateContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_predicate

	return p
}

func (s *PredicateContext) GetParser() antlr.Parser { return s.parser }

func (s *PredicateContext) CopyFrom(ctx *PredicateContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *PredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PredicateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type ExpressionAtomPredicateContext struct {
	*PredicateContext
}

func NewExpressionAtomPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExpressionAtomPredicateContext {
	var p = new(ExpressionAtomPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *ExpressionAtomPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomPredicateContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *ExpressionAtomPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitExpressionAtomPredicate(s)
	}
}

func (s *ExpressionAtomPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitExpressionAtomPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type InPredicateContext struct {
	*PredicateContext
	left IExpressionAtomContext
}

func NewInPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *InPredicateContext {
	var p = new(InPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *InPredicateContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *InPredicateContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *InPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InPredicateContext) IN() antlr.TerminalNode {
	return s.GetToken(SqlParserIN, 0)
}

func (s *InPredicateContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *InPredicateContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *InPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *InPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterInPredicate(s)
	}
}

func (s *InPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitInPredicate(s)
	}
}

func (s *InPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitInPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type BetweenPredicateContext struct {
	*PredicateContext
	left IExpressionAtomContext
	from IExpressionAtomContext
	to   IExpressionAtomContext
}

func NewBetweenPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BetweenPredicateContext {
	var p = new(BetweenPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *BetweenPredicateContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *BetweenPredicateContext) GetFrom() IExpressionAtomContext { return s.from }

func (s *BetweenPredicateContext) GetTo() IExpressionAtomContext { return s.to }

func (s *BetweenPredicateContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *BetweenPredicateContext) SetFrom(v IExpressionAtomContext) { s.from = v }

func (s *BetweenPredicateContext) SetTo(v IExpressionAtomContext) { s.to = v }

func (s *BetweenPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BetweenPredicateContext) BETWEEN() antlr.TerminalNode {
	return s.GetToken(SqlParserBETWEEN, 0)
}

func (s *BetweenPredicateContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *BetweenPredicateContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *BetweenPredicateContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *BetweenPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *BetweenPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBetweenPredicate(s)
	}
}

func (s *BetweenPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBetweenPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type IsNullPredicateContext struct {
	*PredicateContext
}

func NewIsNullPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IsNullPredicateContext {
	var p = new(IsNullPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *IsNullPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IsNullPredicateContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *IsNullPredicateContext) IS() antlr.TerminalNode {
	return s.GetToken(SqlParserIS, 0)
}

func (s *IsNullPredicateContext) NULL_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserNULL_LITERAL, 0)
}

func (s *IsNullPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *IsNullPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitIsNullPredicate(s)
	}
}

func (s *IsNullPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitIsNullPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type BinaryComparasionPredicateContext struct {
	*PredicateContext
	left  IExpressionAtomContext
	right IExpressionAtomContext
}

func NewBinaryComparasionPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BinaryComparasionPredicateContext {
	var p = new(BinaryComparasionPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *BinaryComparasionPredicateContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *BinaryComparasionPredicateContext) GetRight() IExpressionAtomContext { return s.right }

func (s *BinaryComparasionPredicateContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *BinaryComparasionPredicateContext) SetRight(v IExpressionAtomContext) { s.right = v }

func (s *BinaryComparasionPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BinaryComparasionPredicateContext) ComparisonOperator() IComparisonOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IComparisonOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IComparisonOperatorContext)
}

func (s *BinaryComparasionPredicateContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *BinaryComparasionPredicateContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *BinaryComparasionPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBinaryComparasionPredicate(s)
	}
}

func (s *BinaryComparasionPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBinaryComparasionPredicate(s)
	}
}

func (s *BinaryComparasionPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBinaryComparasionPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type LikePredicateContext struct {
	*PredicateContext
	left    IExpressionAtomContext
	likeStr IExpressionAtomContext
}

func NewLikePredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *LikePredicateContext {
	var p = new(LikePredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *LikePredicateContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *LikePredicateContext) GetLikeStr() IExpressionAtomContext { return s.likeStr }

func (s *LikePredicateContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *LikePredicateContext) SetLikeStr(v IExpressionAtomContext) { s.likeStr = v }

func (s *LikePredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LikePredicateContext) LIKE() antlr.TerminalNode {
	return s.GetToken(SqlParserLIKE, 0)
}

func (s *LikePredicateContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *LikePredicateContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *LikePredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *LikePredicateContext) ESCAPE() antlr.TerminalNode {
	return s.GetToken(SqlParserESCAPE, 0)
}

func (s *LikePredicateContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(SqlParserSTRING_LITERAL, 0)
}

func (s *LikePredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLikePredicate(s)
	}
}

func (s *LikePredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLikePredicate(s)
	}
}

func (s *LikePredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLikePredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

type RegexpPredicateContext struct {
	*PredicateContext
	left   IExpressionAtomContext
	regex  antlr.Token
	regStr IExpressionAtomContext
}

func NewRegexpPredicateContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *RegexpPredicateContext {
	var p = new(RegexpPredicateContext)

	p.PredicateContext = NewEmptyPredicateContext()
	p.parser = parser
	p.CopyFrom(ctx.(*PredicateContext))

	return p
}

func (s *RegexpPredicateContext) GetRegex() antlr.Token { return s.regex }

func (s *RegexpPredicateContext) SetRegex(v antlr.Token) { s.regex = v }

func (s *RegexpPredicateContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *RegexpPredicateContext) GetRegStr() IExpressionAtomContext { return s.regStr }

func (s *RegexpPredicateContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *RegexpPredicateContext) SetRegStr(v IExpressionAtomContext) { s.regStr = v }

func (s *RegexpPredicateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexpPredicateContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *RegexpPredicateContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *RegexpPredicateContext) REGEXP() antlr.TerminalNode {
	return s.GetToken(SqlParserREGEXP, 0)
}

func (s *RegexpPredicateContext) RLIKE() antlr.TerminalNode {
	return s.GetToken(SqlParserRLIKE, 0)
}

func (s *RegexpPredicateContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *RegexpPredicateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterRegexpPredicate(s)
	}
}

func (s *RegexpPredicateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitRegexpPredicate(s)
	}
}

func (s *RegexpPredicateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitRegexpPredicate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) Predicate() (localctx IPredicateContext) {
	localctx = NewPredicateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SqlParserRULE_predicate)
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

	p.SetState(271)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 32, p.GetParserRuleContext()) {
	case 1:
		localctx = NewInPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(217)

			var _x = p.expressionAtom(0)

			localctx.(*InPredicateContext).left = _x
		}
		p.SetState(219)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(218)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(221)
			p.Match(SqlParserIN)
		}
		{
			p.SetState(222)
			p.Match(SqlParserLR_BRACKET)
		}
		{
			p.SetState(223)
			p.expressionAtom(0)
		}
		p.SetState(228)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		for _la == SqlParserCOMMA {
			{
				p.SetState(224)
				p.Match(SqlParserCOMMA)
			}
			{
				p.SetState(225)
				p.expressionAtom(0)
			}

			p.SetState(230)
			p.GetErrorHandler().Sync(p)
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(231)
			p.Match(SqlParserRR_BRACKET)
		}

	case 2:
		localctx = NewIsNullPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(233)
			p.expressionAtom(0)
		}
		{
			p.SetState(234)
			p.Match(SqlParserIS)
		}
		p.SetState(236)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(235)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(238)
			p.Match(SqlParserNULL_LITERAL)
		}

	case 3:
		localctx = NewBinaryComparasionPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(240)

			var _x = p.expressionAtom(0)

			localctx.(*BinaryComparasionPredicateContext).left = _x
		}
		{
			p.SetState(241)
			p.ComparisonOperator()
		}
		{
			p.SetState(242)

			var _x = p.expressionAtom(0)

			localctx.(*BinaryComparasionPredicateContext).right = _x
		}

	case 4:
		localctx = NewBetweenPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(244)

			var _x = p.expressionAtom(0)

			localctx.(*BetweenPredicateContext).left = _x
		}
		p.SetState(246)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(245)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(248)
			p.Match(SqlParserBETWEEN)
		}
		{
			p.SetState(249)

			var _x = p.expressionAtom(0)

			localctx.(*BetweenPredicateContext).from = _x
		}
		{
			p.SetState(250)
			p.Match(SqlParserAND)
		}
		{
			p.SetState(251)

			var _x = p.expressionAtom(0)

			localctx.(*BetweenPredicateContext).to = _x
		}

	case 5:
		localctx = NewLikePredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(253)

			var _x = p.expressionAtom(0)

			localctx.(*LikePredicateContext).left = _x
		}
		p.SetState(255)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(254)
				p.Match(SqlParserNOT)
			}

		}
		{
			p.SetState(257)
			p.Match(SqlParserLIKE)
		}
		{
			p.SetState(258)

			var _x = p.expressionAtom(0)

			localctx.(*LikePredicateContext).likeStr = _x
		}
		p.SetState(261)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 30, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(259)
				p.Match(SqlParserESCAPE)
			}
			{
				p.SetState(260)
				p.Match(SqlParserSTRING_LITERAL)
			}

		}

	case 6:
		localctx = NewRegexpPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(263)

			var _x = p.expressionAtom(0)

			localctx.(*RegexpPredicateContext).left = _x
		}
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)

		if _la == SqlParserNOT {
			{
				p.SetState(264)
				p.Match(SqlParserNOT)
			}

		}
		p.SetState(267)

		var _lt = p.GetTokenStream().LT(1)

		localctx.(*RegexpPredicateContext).regex = _lt

		_la = p.GetTokenStream().LA(1)

		if !(_la == SqlParserREGEXP || _la == SqlParserRLIKE) {
			var _ri = p.GetErrorHandler().RecoverInline(p)

			localctx.(*RegexpPredicateContext).regex = _ri
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
		{
			p.SetState(268)

			var _x = p.expressionAtom(0)

			localctx.(*RegexpPredicateContext).regStr = _x
		}

	case 7:
		localctx = NewExpressionAtomPredicateContext(p, localctx)
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(270)
			p.expressionAtom(0)
		}

	}

	return localctx
}

// IExpressionAtomContext is an interface to support dynamic dispatch.
type IExpressionAtomContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsExpressionAtomContext differentiates from other interfaces.
	IsExpressionAtomContext()
}

type ExpressionAtomContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionAtomContext() *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_expressionAtom
	return p
}

func (*ExpressionAtomContext) IsExpressionAtomContext() {}

func NewExpressionAtomContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionAtomContext {
	var p = new(ExpressionAtomContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_expressionAtom

	return p
}

func (s *ExpressionAtomContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionAtomContext) CopyFrom(ctx *ExpressionAtomContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *ExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionAtomContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type UnaryExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewUnaryExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *UnaryExpressionAtomContext {
	var p = new(UnaryExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *UnaryExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionAtomContext) UnaryOperator() IUnaryOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IUnaryOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IUnaryOperatorContext)
}

func (s *UnaryExpressionAtomContext) ExpressionAtom() IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *UnaryExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUnaryExpressionAtom(s)
	}
}

func (s *UnaryExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUnaryExpressionAtom(s)
	}
}

func (s *UnaryExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUnaryExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ColumnNameExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewColumnNameExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ColumnNameExpressionAtomContext {
	var p = new(ColumnNameExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *ColumnNameExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ColumnNameExpressionAtomContext) ColumnName() IColumnNameContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IColumnNameContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IColumnNameContext)
}

func (s *ColumnNameExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterColumnNameExpressionAtom(s)
	}
}

func (s *ColumnNameExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitColumnNameExpressionAtom(s)
	}
}

func (s *ColumnNameExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitColumnNameExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type ConstantExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewConstantExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ConstantExpressionAtomContext {
	var p = new(ConstantExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *ConstantExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstantExpressionAtomContext) Constant() IConstantContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConstantContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConstantContext)
}

func (s *ConstantExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitConstantExpressionAtom(s)
	}
}

func (s *ConstantExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitConstantExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type FunctionCallExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewFunctionCallExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *FunctionCallExpressionAtomContext {
	var p = new(FunctionCallExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *FunctionCallExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallExpressionAtomContext) FunctionCall() IFunctionCallContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFunctionCallContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *FunctionCallExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterFunctionCallExpressionAtom(s)
	}
}

func (s *FunctionCallExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitFunctionCallExpressionAtom(s)
	}
}

func (s *FunctionCallExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitFunctionCallExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type BracketedExpressionAtomContext struct {
	*ExpressionAtomContext
}

func NewBracketedExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BracketedExpressionAtomContext {
	var p = new(BracketedExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *BracketedExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BracketedExpressionAtomContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BracketedExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterBracketedExpressionAtom(s)
	}
}

func (s *BracketedExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitBracketedExpressionAtom(s)
	}
}

func (s *BracketedExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitBracketedExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

type MathExpressionAtomContext struct {
	*ExpressionAtomContext
	left  IExpressionAtomContext
	right IExpressionAtomContext
}

func NewMathExpressionAtomContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MathExpressionAtomContext {
	var p = new(MathExpressionAtomContext)

	p.ExpressionAtomContext = NewEmptyExpressionAtomContext()
	p.parser = parser
	p.CopyFrom(ctx.(*ExpressionAtomContext))

	return p
}

func (s *MathExpressionAtomContext) GetLeft() IExpressionAtomContext { return s.left }

func (s *MathExpressionAtomContext) GetRight() IExpressionAtomContext { return s.right }

func (s *MathExpressionAtomContext) SetLeft(v IExpressionAtomContext) { s.left = v }

func (s *MathExpressionAtomContext) SetRight(v IExpressionAtomContext) { s.right = v }

func (s *MathExpressionAtomContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathExpressionAtomContext) MathOperator() IMathOperatorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMathOperatorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMathOperatorContext)
}

func (s *MathExpressionAtomContext) AllExpressionAtom() []IExpressionAtomContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem())
	var tst = make([]IExpressionAtomContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionAtomContext)
		}
	}

	return tst
}

func (s *MathExpressionAtomContext) ExpressionAtom(i int) IExpressionAtomContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionAtomContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionAtomContext)
}

func (s *MathExpressionAtomContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterMathExpressionAtom(s)
	}
}

func (s *MathExpressionAtomContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitMathExpressionAtom(s)
	}
}

func (s *MathExpressionAtomContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitMathExpressionAtom(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ExpressionAtom() (localctx IExpressionAtomContext) {
	return p.expressionAtom(0)
}

func (p *SqlParser) expressionAtom(_p int) (localctx IExpressionAtomContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpressionAtomContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpressionAtomContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 42
	p.EnterRecursionRule(localctx, 42, SqlParserRULE_expressionAtom, _p)

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
	p.SetState(284)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 33, p.GetParserRuleContext()) {
	case 1:
		localctx = NewConstantExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(274)
			p.Constant()
		}

	case 2:
		localctx = NewColumnNameExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(275)
			p.ColumnName()
		}

	case 3:
		localctx = NewFunctionCallExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(276)
			p.FunctionCall()
		}

	case 4:
		localctx = NewUnaryExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(277)
			p.UnaryOperator()
		}
		{
			p.SetState(278)
			p.expressionAtom(3)
		}

	case 5:
		localctx = NewBracketedExpressionAtomContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(280)
			p.Match(SqlParserLR_BRACKET)
		}
		{
			p.SetState(281)
			p.expression(0)
		}
		{
			p.SetState(282)
			p.Match(SqlParserRR_BRACKET)
		}

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(292)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewMathExpressionAtomContext(p, NewExpressionAtomContext(p, _parentctx, _parentState))
			localctx.(*MathExpressionAtomContext).left = _prevctx

			p.PushNewRecursionContext(localctx, _startState, SqlParserRULE_expressionAtom)
			p.SetState(286)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(287)
				p.MathOperator()
			}
			{
				p.SetState(288)

				var _x = p.expressionAtom(2)

				localctx.(*MathExpressionAtomContext).right = _x
			}

		}
		p.SetState(294)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 34, p.GetParserRuleContext())
	}

	return localctx
}

// IUnaryOperatorContext is an interface to support dynamic dispatch.
type IUnaryOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsUnaryOperatorContext differentiates from other interfaces.
	IsUnaryOperatorContext()
}

type UnaryOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryOperatorContext() *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_unaryOperator
	return p
}

func (*UnaryOperatorContext) IsUnaryOperatorContext() {}

func NewUnaryOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryOperatorContext {
	var p = new(UnaryOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_unaryOperator

	return p
}

func (s *UnaryOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryOperatorContext) NOT() antlr.TerminalNode {
	return s.GetToken(SqlParserNOT, 0)
}

func (s *UnaryOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitUnaryOperator(s)
	}
}

func (s *UnaryOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitUnaryOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) UnaryOperator() (localctx IUnaryOperatorContext) {
	localctx = NewUnaryOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SqlParserRULE_unaryOperator)
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
	p.SetState(295)
	_la = p.GetTokenStream().LA(1)

	if !(_la == SqlParserNOT || (((_la-197)&-(0x1f+1)) == 0 && ((1<<uint((_la-197)))&((1<<(SqlParserPLUS-197))|(1<<(SqlParserMINUS-197))|(1<<(SqlParserEXCLAMATION_SYMBOL-197))|(1<<(SqlParserBIT_NOT_OP-197)))) != 0)) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

// IComparisonOperatorContext is an interface to support dynamic dispatch.
type IComparisonOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsComparisonOperatorContext differentiates from other interfaces.
	IsComparisonOperatorContext()
}

type ComparisonOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyComparisonOperatorContext() *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_comparisonOperator
	return p
}

func (*ComparisonOperatorContext) IsComparisonOperatorContext() {}

func NewComparisonOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ComparisonOperatorContext {
	var p = new(ComparisonOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_comparisonOperator

	return p
}

func (s *ComparisonOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *ComparisonOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ComparisonOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ComparisonOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitComparisonOperator(s)
	}
}

func (s *ComparisonOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitComparisonOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) ComparisonOperator() (localctx IComparisonOperatorContext) {
	localctx = NewComparisonOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SqlParserRULE_comparisonOperator)

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

	p.SetState(306)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 35, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(297)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(298)
			p.Match(SqlParserGREATER_SYMBOL)
		}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(299)
			p.Match(SqlParserLESS_SYMBOL)
		}

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(300)
			p.Match(SqlParserLESS_SYMBOL)
		}
		{
			p.SetState(301)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(302)
			p.Match(SqlParserGREATER_SYMBOL)
		}
		{
			p.SetState(303)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(304)
			p.Match(SqlParserEXCLAMATION_SYMBOL)
		}
		{
			p.SetState(305)
			p.Match(SqlParserEQUAL_SYMBOL)
		}

	}

	return localctx
}

// ILogicalOperatorContext is an interface to support dynamic dispatch.
type ILogicalOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsLogicalOperatorContext differentiates from other interfaces.
	IsLogicalOperatorContext()
}

type LogicalOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLogicalOperatorContext() *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_logicalOperator
	return p
}

func (*LogicalOperatorContext) IsLogicalOperatorContext() {}

func NewLogicalOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LogicalOperatorContext {
	var p = new(LogicalOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_logicalOperator

	return p
}

func (s *LogicalOperatorContext) GetParser() antlr.Parser { return s.parser }

func (s *LogicalOperatorContext) AND() antlr.TerminalNode {
	return s.GetToken(SqlParserAND, 0)
}

func (s *LogicalOperatorContext) OR() antlr.TerminalNode {
	return s.GetToken(SqlParserOR, 0)
}

func (s *LogicalOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LogicalOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LogicalOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitLogicalOperator(s)
	}
}

func (s *LogicalOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitLogicalOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) LogicalOperator() (localctx ILogicalOperatorContext) {
	localctx = NewLogicalOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SqlParserRULE_logicalOperator)

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

	p.SetState(314)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case SqlParserAND:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(308)
			p.Match(SqlParserAND)
		}

	case SqlParserBIT_AND_OP:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(309)
			p.Match(SqlParserBIT_AND_OP)
		}
		{
			p.SetState(310)
			p.Match(SqlParserBIT_AND_OP)
		}

	case SqlParserOR:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(311)
			p.Match(SqlParserOR)
		}

	case SqlParserBIT_OR_OP:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(312)
			p.Match(SqlParserBIT_OR_OP)
		}
		{
			p.SetState(313)
			p.Match(SqlParserBIT_OR_OP)
		}

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IMathOperatorContext is an interface to support dynamic dispatch.
type IMathOperatorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// IsMathOperatorContext differentiates from other interfaces.
	IsMathOperatorContext()
}

type MathOperatorContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMathOperatorContext() *MathOperatorContext {
	var p = new(MathOperatorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = SqlParserRULE_mathOperator
	return p
}

func (*MathOperatorContext) IsMathOperatorContext() {}

func NewMathOperatorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MathOperatorContext {
	var p = new(MathOperatorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = SqlParserRULE_mathOperator

	return p
}

func (s *MathOperatorContext) GetParser() antlr.Parser { return s.parser }
func (s *MathOperatorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MathOperatorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MathOperatorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.EnterMathOperator(s)
	}
}

func (s *MathOperatorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SqlParserListener); ok {
		listenerT.ExitMathOperator(s)
	}
}

func (s *MathOperatorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SqlParserVisitor:
		return t.VisitMathOperator(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SqlParser) MathOperator() (localctx IMathOperatorContext) {
	localctx = NewMathOperatorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SqlParserRULE_mathOperator)
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
	p.SetState(316)
	_la = p.GetTokenStream().LA(1)

	if !(((_la-194)&-(0x1f+1)) == 0 && ((1<<uint((_la-194)))&((1<<(SqlParserSTAR-194))|(1<<(SqlParserDIVIDE-194))|(1<<(SqlParserMODULE-194))|(1<<(SqlParserPLUS-194))|(1<<(SqlParserMINUS-194)))) != 0) {
		p.GetErrorHandler().RecoverInline(p)
	} else {
		p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}

	return localctx
}

func (p *SqlParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 19:
		var t *ExpressionContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionContext)
		}
		return p.Expression_Sempred(t, predIndex)

	case 21:
		var t *ExpressionAtomContext = nil
		if localctx != nil {
			t = localctx.(*ExpressionAtomContext)
		}
		return p.ExpressionAtom_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *SqlParser) Expression_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *SqlParser) ExpressionAtom_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
