// Generated from Lambda.g4 by ANTLR 4.7.

package lambda

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 39, 407,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3,
	10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13,
	3, 14, 3, 14, 3, 15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 18, 3, 18, 3,
	18, 3, 19, 3, 19, 3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 22, 3, 22,
	3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 25, 3, 25, 3, 25, 3, 26, 3,
	26, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 7, 28, 192, 10, 28, 12, 28,
	14, 28, 195, 11, 28, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 5,
	29, 204, 10, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30,
	3, 30, 3, 30, 5, 30, 216, 10, 30, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 32, 3, 32, 3, 32, 3, 32, 5, 32, 228, 10, 32, 3, 33, 3, 33, 3, 33,
	5, 33, 233, 10, 33, 3, 34, 3, 34, 3, 34, 5, 34, 238, 10, 34, 3, 35, 3,
	35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 5, 35, 249, 10, 35,
	3, 36, 3, 36, 7, 36, 253, 10, 36, 12, 36, 14, 36, 256, 11, 36, 3, 37, 3,
	37, 7, 37, 260, 10, 37, 12, 37, 14, 37, 263, 11, 37, 3, 38, 3, 38, 3, 38,
	6, 38, 268, 10, 38, 13, 38, 14, 38, 269, 3, 39, 3, 39, 3, 39, 5, 39, 275,
	10, 39, 3, 39, 5, 39, 278, 10, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3,
	39, 5, 39, 286, 10, 39, 5, 39, 288, 10, 39, 3, 40, 6, 40, 291, 10, 40,
	13, 40, 14, 40, 292, 3, 41, 3, 41, 5, 41, 297, 10, 41, 3, 41, 3, 41, 3,
	42, 3, 42, 5, 42, 303, 10, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 5, 43,
	310, 10, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3, 44, 5, 44, 318, 10,
	44, 3, 45, 3, 45, 5, 45, 322, 10, 45, 3, 46, 3, 46, 3, 46, 3, 46, 3, 46,
	3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3, 48, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49,
	3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 5, 51, 359,
	10, 51, 3, 52, 3, 52, 3, 52, 7, 52, 364, 10, 52, 12, 52, 14, 52, 367, 11,
	52, 3, 52, 3, 52, 3, 53, 3, 53, 3, 53, 7, 53, 374, 10, 53, 12, 53, 14,
	53, 377, 11, 53, 3, 53, 3, 53, 3, 54, 3, 54, 5, 54, 383, 10, 54, 3, 55,
	3, 55, 3, 56, 3, 56, 3, 57, 3, 57, 3, 58, 3, 58, 3, 59, 3, 59, 3, 60, 5,
	60, 396, 10, 60, 3, 61, 5, 61, 399, 10, 61, 3, 62, 6, 62, 402, 10, 62,
	13, 62, 14, 62, 403, 3, 62, 3, 62, 2, 2, 63, 3, 3, 5, 4, 7, 5, 9, 6, 11,
	7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25, 14, 27, 15, 29, 16,
	31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43, 23, 45, 24, 47, 25,
	49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 2, 61, 2, 63, 2, 65, 2, 67,
	31, 69, 32, 71, 2, 73, 2, 75, 2, 77, 33, 79, 2, 81, 2, 83, 34, 85, 35,
	87, 2, 89, 2, 91, 2, 93, 2, 95, 36, 97, 37, 99, 2, 101, 38, 103, 2, 105,
	2, 107, 2, 109, 2, 111, 2, 113, 2, 115, 2, 117, 2, 119, 2, 121, 2, 123,
	39, 3, 2, 17, 6, 2, 45, 45, 47, 47, 96, 96, 126, 126, 5, 2, 39, 39, 44,
	44, 49, 49, 7, 2, 35, 35, 40, 40, 44, 45, 47, 47, 96, 96, 3, 2, 51, 59,
	4, 2, 90, 90, 122, 122, 4, 2, 71, 71, 103, 103, 4, 2, 45, 45, 47, 47, 11,
	2, 36, 36, 41, 41, 94, 94, 99, 100, 104, 104, 112, 112, 116, 116, 118,
	118, 120, 120, 3, 2, 50, 59, 3, 2, 50, 57, 5, 2, 50, 59, 67, 72, 99, 104,
	3, 2, 12, 12, 22, 2, 50, 59, 1634, 1643, 1778, 1787, 2408, 2417, 2536,
	2545, 2664, 2673, 2792, 2801, 2920, 2929, 3049, 3057, 3176, 3185, 3304,
	3313, 3432, 3441, 3666, 3675, 3794, 3803, 3874, 3883, 4162, 4171, 4971,
	4979, 6114, 6123, 6162, 6171, 65298, 65307, 260, 2, 67, 92, 99, 124, 172,
	172, 183, 183, 188, 188, 194, 216, 218, 248, 250, 545, 548, 565, 594, 687,
	690, 698, 701, 707, 722, 723, 738, 742, 752, 752, 892, 892, 904, 904, 906,
	908, 910, 910, 912, 931, 933, 976, 978, 985, 988, 1013, 1026, 1155, 1166,
	1222, 1225, 1226, 1229, 1230, 1234, 1271, 1274, 1275, 1331, 1368, 1371,
	1371, 1379, 1417, 1490, 1516, 1522, 1524, 1571, 1596, 1602, 1612, 1651,
	1749, 1751, 1751, 1767, 1768, 1788, 1790, 1810, 1810, 1812, 1838, 1922,
	1959, 2311, 2363, 2367, 2367, 2386, 2386, 2394, 2403, 2439, 2446, 2449,
	2450, 2453, 2474, 2476, 2482, 2484, 2484, 2488, 2491, 2526, 2527, 2529,
	2531, 2546, 2547, 2567, 2572, 2577, 2578, 2581, 2602, 2604, 2610, 2612,
	2613, 2615, 2616, 2618, 2619, 2651, 2654, 2656, 2656, 2676, 2678, 2695,
	2701, 2703, 2703, 2705, 2707, 2709, 2730, 2732, 2738, 2740, 2741, 2743,
	2747, 2751, 2751, 2770, 2770, 2786, 2786, 2823, 2830, 2833, 2834, 2837,
	2858, 2860, 2866, 2868, 2869, 2872, 2875, 2879, 2879, 2910, 2911, 2913,
	2915, 2951, 2956, 2960, 2962, 2964, 2967, 2971, 2972, 2974, 2974, 2976,
	2977, 2981, 2982, 2986, 2988, 2992, 2999, 3001, 3003, 3079, 3086, 3088,
	3090, 3092, 3114, 3116, 3125, 3127, 3131, 3170, 3171, 3207, 3214, 3216,
	3218, 3220, 3242, 3244, 3253, 3255, 3259, 3296, 3296, 3298, 3299, 3335,
	3342, 3344, 3346, 3348, 3370, 3372, 3387, 3426, 3427, 3463, 3480, 3484,
	3507, 3509, 3517, 3519, 3519, 3522, 3528, 3587, 3634, 3636, 3637, 3650,
	3656, 3715, 3716, 3718, 3718, 3721, 3722, 3724, 3724, 3727, 3727, 3734,
	3737, 3739, 3745, 3747, 3749, 3751, 3751, 3753, 3753, 3756, 3757, 3759,
	3762, 3764, 3765, 3775, 3782, 3784, 3784, 3806, 3807, 3842, 3842, 3906,
	3948, 3978, 3981, 4098, 4131, 4133, 4137, 4139, 4140, 4178, 4183, 4258,
	4295, 4306, 4344, 4354, 4443, 4449, 4516, 4522, 4603, 4610, 4616, 4618,
	4680, 4682, 4682, 4684, 4687, 4690, 4696, 4698, 4698, 4700, 4703, 4706,
	4744, 4746, 4746, 4748, 4751, 4754, 4784, 4786, 4786, 4788, 4791, 4794,
	4800, 4802, 4802, 4804, 4807, 4810, 4816, 4818, 4824, 4826, 4848, 4850,
	4880, 4882, 4882, 4884, 4887, 4890, 4896, 4898, 4936, 4938, 4956, 5026,
	5110, 5123, 5752, 5763, 5788, 5794, 5868, 6018, 6069, 6178, 6265, 6274,
	6314, 7682, 7837, 7842, 7931, 7938, 7959, 7962, 7967, 7970, 8007, 8010,
	8015, 8018, 8025, 8027, 8027, 8029, 8029, 8031, 8031, 8033, 8063, 8066,
	8118, 8120, 8126, 8128, 8128, 8132, 8134, 8136, 8142, 8146, 8149, 8152,
	8157, 8162, 8174, 8180, 8182, 8184, 8190, 8321, 8321, 8452, 8452, 8457,
	8457, 8460, 8469, 8471, 8471, 8475, 8479, 8486, 8486, 8488, 8488, 8490,
	8490, 8492, 8495, 8497, 8499, 8501, 8507, 8546, 8581, 12295, 12297, 12323,
	12331, 12339, 12343, 12346, 12348, 12355, 12438, 12447, 12448, 12451, 12540,
	12542, 12544, 12551, 12590, 12595, 12688, 12706, 12729, 13314, 13314, 19895,
	19895, 19970, 19970, 40871, 40871, 40962, 42126, 44034, 44034, 55205, 55205,
	63746, 64047, 64258, 64264, 64277, 64281, 64287, 64287, 64289, 64298, 64300,
	64312, 64314, 64318, 64320, 64320, 64322, 64323, 64325, 64326, 64328, 64435,
	64469, 64831, 64850, 64913, 64916, 64969, 65010, 65021, 65138, 65140, 65142,
	65142, 65144, 65278, 65315, 65340, 65347, 65372, 65384, 65472, 65476, 65481,
	65484, 65489, 65492, 65497, 65500, 65502, 5, 2, 11, 12, 15, 15, 34, 34,
	2, 424, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3,
	2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17,
	3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2,
	25, 3, 2, 2, 2, 2, 27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2,
	2, 33, 3, 2, 2, 2, 2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2,
	2, 2, 41, 3, 2, 2, 2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2,
	2, 2, 2, 49, 3, 2, 2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3,
	2, 2, 2, 2, 57, 3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 77,
	3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 95, 3, 2, 2, 2, 2,
	97, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 123, 3, 2, 2, 2, 3, 125, 3, 2, 2,
	2, 5, 128, 3, 2, 2, 2, 7, 130, 3, 2, 2, 2, 9, 132, 3, 2, 2, 2, 11, 134,
	3, 2, 2, 2, 13, 136, 3, 2, 2, 2, 15, 138, 3, 2, 2, 2, 17, 140, 3, 2, 2,
	2, 19, 142, 3, 2, 2, 2, 21, 145, 3, 2, 2, 2, 23, 148, 3, 2, 2, 2, 25, 150,
	3, 2, 2, 2, 27, 153, 3, 2, 2, 2, 29, 155, 3, 2, 2, 2, 31, 157, 3, 2, 2,
	2, 33, 159, 3, 2, 2, 2, 35, 161, 3, 2, 2, 2, 37, 164, 3, 2, 2, 2, 39, 167,
	3, 2, 2, 2, 41, 169, 3, 2, 2, 2, 43, 172, 3, 2, 2, 2, 45, 174, 3, 2, 2,
	2, 47, 177, 3, 2, 2, 2, 49, 180, 3, 2, 2, 2, 51, 183, 3, 2, 2, 2, 53, 185,
	3, 2, 2, 2, 55, 188, 3, 2, 2, 2, 57, 203, 3, 2, 2, 2, 59, 215, 3, 2, 2,
	2, 61, 217, 3, 2, 2, 2, 63, 227, 3, 2, 2, 2, 65, 232, 3, 2, 2, 2, 67, 237,
	3, 2, 2, 2, 69, 248, 3, 2, 2, 2, 71, 250, 3, 2, 2, 2, 73, 257, 3, 2, 2,
	2, 75, 264, 3, 2, 2, 2, 77, 287, 3, 2, 2, 2, 79, 290, 3, 2, 2, 2, 81, 294,
	3, 2, 2, 2, 83, 302, 3, 2, 2, 2, 85, 306, 3, 2, 2, 2, 87, 317, 3, 2, 2,
	2, 89, 321, 3, 2, 2, 2, 91, 323, 3, 2, 2, 2, 93, 328, 3, 2, 2, 2, 95, 333,
	3, 2, 2, 2, 97, 341, 3, 2, 2, 2, 99, 353, 3, 2, 2, 2, 101, 358, 3, 2, 2,
	2, 103, 360, 3, 2, 2, 2, 105, 370, 3, 2, 2, 2, 107, 382, 3, 2, 2, 2, 109,
	384, 3, 2, 2, 2, 111, 386, 3, 2, 2, 2, 113, 388, 3, 2, 2, 2, 115, 390,
	3, 2, 2, 2, 117, 392, 3, 2, 2, 2, 119, 395, 3, 2, 2, 2, 121, 398, 3, 2,
	2, 2, 123, 401, 3, 2, 2, 2, 125, 126, 7, 63, 2, 2, 126, 127, 7, 64, 2,
	2, 127, 4, 3, 2, 2, 2, 128, 129, 7, 42, 2, 2, 129, 6, 3, 2, 2, 2, 130,
	131, 7, 46, 2, 2, 131, 8, 3, 2, 2, 2, 132, 133, 7, 43, 2, 2, 133, 10, 3,
	2, 2, 2, 134, 135, 7, 48, 2, 2, 135, 12, 3, 2, 2, 2, 136, 137, 7, 44, 2,
	2, 137, 14, 3, 2, 2, 2, 138, 139, 7, 49, 2, 2, 139, 16, 3, 2, 2, 2, 140,
	141, 7, 39, 2, 2, 141, 18, 3, 2, 2, 2, 142, 143, 7, 62, 2, 2, 143, 144,
	7, 62, 2, 2, 144, 20, 3, 2, 2, 2, 145, 146, 7, 64, 2, 2, 146, 147, 7, 64,
	2, 2, 147, 22, 3, 2, 2, 2, 148, 149, 7, 40, 2, 2, 149, 24, 3, 2, 2, 2,
	150, 151, 7, 40, 2, 2, 151, 152, 7, 96, 2, 2, 152, 26, 3, 2, 2, 2, 153,
	154, 7, 45, 2, 2, 154, 28, 3, 2, 2, 2, 155, 156, 7, 47, 2, 2, 156, 30,
	3, 2, 2, 2, 157, 158, 7, 126, 2, 2, 158, 32, 3, 2, 2, 2, 159, 160, 7, 96,
	2, 2, 160, 34, 3, 2, 2, 2, 161, 162, 7, 63, 2, 2, 162, 163, 7, 63, 2, 2,
	163, 36, 3, 2, 2, 2, 164, 165, 7, 35, 2, 2, 165, 166, 7, 63, 2, 2, 166,
	38, 3, 2, 2, 2, 167, 168, 7, 62, 2, 2, 168, 40, 3, 2, 2, 2, 169, 170, 7,
	62, 2, 2, 170, 171, 7, 63, 2, 2, 171, 42, 3, 2, 2, 2, 172, 173, 7, 64,
	2, 2, 173, 44, 3, 2, 2, 2, 174, 175, 7, 64, 2, 2, 175, 176, 7, 63, 2, 2,
	176, 46, 3, 2, 2, 2, 177, 178, 7, 40, 2, 2, 178, 179, 7, 40, 2, 2, 179,
	48, 3, 2, 2, 2, 180, 181, 7, 126, 2, 2, 181, 182, 7, 126, 2, 2, 182, 50,
	3, 2, 2, 2, 183, 184, 7, 35, 2, 2, 184, 52, 3, 2, 2, 2, 185, 186, 7, 62,
	2, 2, 186, 187, 7, 47, 2, 2, 187, 54, 3, 2, 2, 2, 188, 193, 5, 107, 54,
	2, 189, 192, 5, 107, 54, 2, 190, 192, 5, 119, 60, 2, 191, 189, 3, 2, 2,
	2, 191, 190, 3, 2, 2, 2, 192, 195, 3, 2, 2, 2, 193, 191, 3, 2, 2, 2, 193,
	194, 3, 2, 2, 2, 194, 56, 3, 2, 2, 2, 195, 193, 3, 2, 2, 2, 196, 197, 7,
	126, 2, 2, 197, 204, 7, 126, 2, 2, 198, 199, 7, 40, 2, 2, 199, 204, 7,
	40, 2, 2, 200, 204, 5, 59, 30, 2, 201, 204, 5, 61, 31, 2, 202, 204, 5,
	63, 32, 2, 203, 196, 3, 2, 2, 2, 203, 198, 3, 2, 2, 2, 203, 200, 3, 2,
	2, 2, 203, 201, 3, 2, 2, 2, 203, 202, 3, 2, 2, 2, 204, 58, 3, 2, 2, 2,
	205, 206, 7, 63, 2, 2, 206, 216, 7, 63, 2, 2, 207, 208, 7, 35, 2, 2, 208,
	216, 7, 63, 2, 2, 209, 216, 7, 62, 2, 2, 210, 211, 7, 62, 2, 2, 211, 216,
	7, 63, 2, 2, 212, 216, 7, 64, 2, 2, 213, 214, 7, 64, 2, 2, 214, 216, 7,
	63, 2, 2, 215, 205, 3, 2, 2, 2, 215, 207, 3, 2, 2, 2, 215, 209, 3, 2, 2,
	2, 215, 210, 3, 2, 2, 2, 215, 212, 3, 2, 2, 2, 215, 213, 3, 2, 2, 2, 216,
	60, 3, 2, 2, 2, 217, 218, 9, 2, 2, 2, 218, 62, 3, 2, 2, 2, 219, 228, 9,
	3, 2, 2, 220, 221, 7, 62, 2, 2, 221, 228, 7, 62, 2, 2, 222, 223, 7, 64,
	2, 2, 223, 228, 7, 64, 2, 2, 224, 228, 7, 40, 2, 2, 225, 226, 7, 40, 2,
	2, 226, 228, 7, 96, 2, 2, 227, 219, 3, 2, 2, 2, 227, 220, 3, 2, 2, 2, 227,
	222, 3, 2, 2, 2, 227, 224, 3, 2, 2, 2, 227, 225, 3, 2, 2, 2, 228, 64, 3,
	2, 2, 2, 229, 233, 9, 4, 2, 2, 230, 231, 7, 62, 2, 2, 231, 233, 7, 47,
	2, 2, 232, 229, 3, 2, 2, 2, 232, 230, 3, 2, 2, 2, 233, 66, 3, 2, 2, 2,
	234, 238, 5, 71, 36, 2, 235, 238, 5, 73, 37, 2, 236, 238, 5, 75, 38, 2,
	237, 234, 3, 2, 2, 2, 237, 235, 3, 2, 2, 2, 237, 236, 3, 2, 2, 2, 238,
	68, 3, 2, 2, 2, 239, 240, 7, 118, 2, 2, 240, 241, 7, 116, 2, 2, 241, 242,
	7, 119, 2, 2, 242, 249, 7, 103, 2, 2, 243, 244, 7, 104, 2, 2, 244, 245,
	7, 99, 2, 2, 245, 246, 7, 110, 2, 2, 246, 247, 7, 117, 2, 2, 247, 249,
	7, 103, 2, 2, 248, 239, 3, 2, 2, 2, 248, 243, 3, 2, 2, 2, 249, 70, 3, 2,
	2, 2, 250, 254, 9, 5, 2, 2, 251, 253, 5, 109, 55, 2, 252, 251, 3, 2, 2,
	2, 253, 256, 3, 2, 2, 2, 254, 252, 3, 2, 2, 2, 254, 255, 3, 2, 2, 2, 255,
	72, 3, 2, 2, 2, 256, 254, 3, 2, 2, 2, 257, 261, 7, 50, 2, 2, 258, 260,
	5, 111, 56, 2, 259, 258, 3, 2, 2, 2, 260, 263, 3, 2, 2, 2, 261, 259, 3,
	2, 2, 2, 261, 262, 3, 2, 2, 2, 262, 74, 3, 2, 2, 2, 263, 261, 3, 2, 2,
	2, 264, 265, 7, 50, 2, 2, 265, 267, 9, 6, 2, 2, 266, 268, 5, 113, 57, 2,
	267, 266, 3, 2, 2, 2, 268, 269, 3, 2, 2, 2, 269, 267, 3, 2, 2, 2, 269,
	270, 3, 2, 2, 2, 270, 76, 3, 2, 2, 2, 271, 272, 5, 79, 40, 2, 272, 274,
	7, 48, 2, 2, 273, 275, 5, 79, 40, 2, 274, 273, 3, 2, 2, 2, 274, 275, 3,
	2, 2, 2, 275, 277, 3, 2, 2, 2, 276, 278, 5, 81, 41, 2, 277, 276, 3, 2,
	2, 2, 277, 278, 3, 2, 2, 2, 278, 288, 3, 2, 2, 2, 279, 280, 5, 79, 40,
	2, 280, 281, 5, 81, 41, 2, 281, 288, 3, 2, 2, 2, 282, 283, 7, 48, 2, 2,
	283, 285, 5, 79, 40, 2, 284, 286, 5, 81, 41, 2, 285, 284, 3, 2, 2, 2, 285,
	286, 3, 2, 2, 2, 286, 288, 3, 2, 2, 2, 287, 271, 3, 2, 2, 2, 287, 279,
	3, 2, 2, 2, 287, 282, 3, 2, 2, 2, 288, 78, 3, 2, 2, 2, 289, 291, 5, 109,
	55, 2, 290, 289, 3, 2, 2, 2, 291, 292, 3, 2, 2, 2, 292, 290, 3, 2, 2, 2,
	292, 293, 3, 2, 2, 2, 293, 80, 3, 2, 2, 2, 294, 296, 9, 7, 2, 2, 295, 297,
	9, 8, 2, 2, 296, 295, 3, 2, 2, 2, 296, 297, 3, 2, 2, 2, 297, 298, 3, 2,
	2, 2, 298, 299, 5, 79, 40, 2, 299, 82, 3, 2, 2, 2, 300, 303, 5, 79, 40,
	2, 301, 303, 5, 77, 39, 2, 302, 300, 3, 2, 2, 2, 302, 301, 3, 2, 2, 2,
	303, 304, 3, 2, 2, 2, 304, 305, 7, 107, 2, 2, 305, 84, 3, 2, 2, 2, 306,
	309, 7, 41, 2, 2, 307, 310, 5, 87, 44, 2, 308, 310, 5, 89, 45, 2, 309,
	307, 3, 2, 2, 2, 309, 308, 3, 2, 2, 2, 310, 311, 3, 2, 2, 2, 311, 312,
	7, 41, 2, 2, 312, 86, 3, 2, 2, 2, 313, 318, 5, 117, 59, 2, 314, 318, 5,
	95, 48, 2, 315, 318, 5, 97, 49, 2, 316, 318, 5, 99, 50, 2, 317, 313, 3,
	2, 2, 2, 317, 314, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 317, 316, 3, 2, 2,
	2, 318, 88, 3, 2, 2, 2, 319, 322, 5, 91, 46, 2, 320, 322, 5, 93, 47, 2,
	321, 319, 3, 2, 2, 2, 321, 320, 3, 2, 2, 2, 322, 90, 3, 2, 2, 2, 323, 324,
	7, 94, 2, 2, 324, 325, 5, 111, 56, 2, 325, 326, 5, 111, 56, 2, 326, 327,
	5, 111, 56, 2, 327, 92, 3, 2, 2, 2, 328, 329, 7, 94, 2, 2, 329, 330, 7,
	122, 2, 2, 330, 331, 5, 113, 57, 2, 331, 332, 5, 113, 57, 2, 332, 94, 3,
	2, 2, 2, 333, 334, 7, 94, 2, 2, 334, 335, 7, 119, 2, 2, 335, 336, 3, 2,
	2, 2, 336, 337, 5, 113, 57, 2, 337, 338, 5, 113, 57, 2, 338, 339, 5, 113,
	57, 2, 339, 340, 5, 113, 57, 2, 340, 96, 3, 2, 2, 2, 341, 342, 7, 94, 2,
	2, 342, 343, 7, 87, 2, 2, 343, 344, 3, 2, 2, 2, 344, 345, 5, 113, 57, 2,
	345, 346, 5, 113, 57, 2, 346, 347, 5, 113, 57, 2, 347, 348, 5, 113, 57,
	2, 348, 349, 5, 113, 57, 2, 349, 350, 5, 113, 57, 2, 350, 351, 5, 113,
	57, 2, 351, 352, 5, 113, 57, 2, 352, 98, 3, 2, 2, 2, 353, 354, 7, 94, 2,
	2, 354, 355, 9, 9, 2, 2, 355, 100, 3, 2, 2, 2, 356, 359, 5, 103, 52, 2,
	357, 359, 5, 105, 53, 2, 358, 356, 3, 2, 2, 2, 358, 357, 3, 2, 2, 2, 359,
	102, 3, 2, 2, 2, 360, 365, 7, 98, 2, 2, 361, 364, 5, 117, 59, 2, 362, 364,
	5, 115, 58, 2, 363, 361, 3, 2, 2, 2, 363, 362, 3, 2, 2, 2, 364, 367, 3,
	2, 2, 2, 365, 363, 3, 2, 2, 2, 365, 366, 3, 2, 2, 2, 366, 368, 3, 2, 2,
	2, 367, 365, 3, 2, 2, 2, 368, 369, 7, 98, 2, 2, 369, 104, 3, 2, 2, 2, 370,
	375, 7, 36, 2, 2, 371, 374, 5, 87, 44, 2, 372, 374, 5, 89, 45, 2, 373,
	371, 3, 2, 2, 2, 373, 372, 3, 2, 2, 2, 374, 377, 3, 2, 2, 2, 375, 373,
	3, 2, 2, 2, 375, 376, 3, 2, 2, 2, 376, 378, 3, 2, 2, 2, 377, 375, 3, 2,
	2, 2, 378, 379, 7, 36, 2, 2, 379, 106, 3, 2, 2, 2, 380, 383, 5, 121, 61,
	2, 381, 383, 7, 97, 2, 2, 382, 380, 3, 2, 2, 2, 382, 381, 3, 2, 2, 2, 383,
	108, 3, 2, 2, 2, 384, 385, 9, 10, 2, 2, 385, 110, 3, 2, 2, 2, 386, 387,
	9, 11, 2, 2, 387, 112, 3, 2, 2, 2, 388, 389, 9, 12, 2, 2, 389, 114, 3,
	2, 2, 2, 390, 391, 9, 13, 2, 2, 391, 116, 3, 2, 2, 2, 392, 393, 10, 13,
	2, 2, 393, 118, 3, 2, 2, 2, 394, 396, 9, 14, 2, 2, 395, 394, 3, 2, 2, 2,
	396, 120, 3, 2, 2, 2, 397, 399, 9, 15, 2, 2, 398, 397, 3, 2, 2, 2, 399,
	122, 3, 2, 2, 2, 400, 402, 9, 16, 2, 2, 401, 400, 3, 2, 2, 2, 402, 403,
	3, 2, 2, 2, 403, 401, 3, 2, 2, 2, 403, 404, 3, 2, 2, 2, 404, 405, 3, 2,
	2, 2, 405, 406, 8, 62, 2, 2, 406, 124, 3, 2, 2, 2, 33, 2, 191, 193, 203,
	215, 227, 232, 237, 248, 254, 261, 269, 274, 277, 285, 287, 292, 296, 302,
	309, 317, 321, 358, 363, 365, 373, 375, 382, 395, 398, 403, 3, 8, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'=>'", "'('", "','", "')'", "'.'", "'*'", "'/'", "'%'", "'<<'", "'>>'",
	"'&'", "'&^'", "'+'", "'-'", "'|'", "'^'", "'=='", "'!='", "'<'", "'<='",
	"'>'", "'>='", "'&&'", "'||'", "'!'", "'<-'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
	"", "", "", "", "", "", "", "", "", "IDENTIFIER", "BINARY_OP", "INT_LIT",
	"BOOL_LIT", "FLOAT_LIT", "IMAGINARY_LIT", "RUNE_LIT", "LITTLE_U_VALUE",
	"BIG_U_VALUE", "STRING_LIT", "WS",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "T__11", "T__12", "T__13", "T__14", "T__15", "T__16",
	"T__17", "T__18", "T__19", "T__20", "T__21", "T__22", "T__23", "T__24",
	"T__25", "IDENTIFIER", "BINARY_OP", "REL_OP", "ADD_OP", "MUL_OP", "UNARY_OP",
	"INT_LIT", "BOOL_LIT", "DECIMAL_LIT", "OCTAL_LIT", "HEX_LIT", "FLOAT_LIT",
	"DECIMALS", "EXPONENT", "IMAGINARY_LIT", "RUNE_LIT", "UNICODE_VALUE", "BYTE_VALUE",
	"OCTAL_BYTE_VALUE", "HEX_BYTE_VALUE", "LITTLE_U_VALUE", "BIG_U_VALUE",
	"ESCAPED_CHAR", "STRING_LIT", "RAW_STRING_LIT", "INTERPRETED_STRING_LIT",
	"LETTER", "DECIMAL_DIGIT", "OCTAL_DIGIT", "HEX_DIGIT", "NEWLINE", "UNICODE_CHAR",
	"UNICODE_DIGIT", "UNICODE_LETTER", "WS",
}

type LambdaLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewLambdaLexer(input antlr.CharStream) *LambdaLexer {

	l := new(LambdaLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Lambda.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// LambdaLexer tokens.
const (
	LambdaLexerT__0           = 1
	LambdaLexerT__1           = 2
	LambdaLexerT__2           = 3
	LambdaLexerT__3           = 4
	LambdaLexerT__4           = 5
	LambdaLexerT__5           = 6
	LambdaLexerT__6           = 7
	LambdaLexerT__7           = 8
	LambdaLexerT__8           = 9
	LambdaLexerT__9           = 10
	LambdaLexerT__10          = 11
	LambdaLexerT__11          = 12
	LambdaLexerT__12          = 13
	LambdaLexerT__13          = 14
	LambdaLexerT__14          = 15
	LambdaLexerT__15          = 16
	LambdaLexerT__16          = 17
	LambdaLexerT__17          = 18
	LambdaLexerT__18          = 19
	LambdaLexerT__19          = 20
	LambdaLexerT__20          = 21
	LambdaLexerT__21          = 22
	LambdaLexerT__22          = 23
	LambdaLexerT__23          = 24
	LambdaLexerT__24          = 25
	LambdaLexerT__25          = 26
	LambdaLexerIDENTIFIER     = 27
	LambdaLexerBINARY_OP      = 28
	LambdaLexerINT_LIT        = 29
	LambdaLexerBOOL_LIT       = 30
	LambdaLexerFLOAT_LIT      = 31
	LambdaLexerIMAGINARY_LIT  = 32
	LambdaLexerRUNE_LIT       = 33
	LambdaLexerLITTLE_U_VALUE = 34
	LambdaLexerBIG_U_VALUE    = 35
	LambdaLexerSTRING_LIT     = 36
	LambdaLexerWS             = 37
)
