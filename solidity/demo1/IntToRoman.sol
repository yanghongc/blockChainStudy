// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract IntToRoman {
    // 主函数：把正整数 (1..3999) 转为罗马数字
    function intToRoman(uint256 num) public pure returns (string memory) {
        require(num >= 1 && num <= 3999, "num out of range (1..3999)");

        // 值与对应罗马字母（从大到小）
        uint16[13] memory vals = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1];
        string[13] memory syms = [
            "M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"
        ];

        bytes memory out = new bytes(0);

        for (uint i = 0; i < vals.length; i++) {
            uint16 v = vals[i];
            string memory s = syms[i];
            while (num >= v) {
                num -= v;
                out = abi.encodePacked(out, s);
            }
            if (num == 0) {
                break;
            }
        }
        return string(out);
    }
}