// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
    反转字符串 (Reverse String)
    题目描述：反转一个字符串。输入 "abcde"，输出 "edcba"
*/
contract ReveStr {

    function reveStr(string memory str) public pure returns (string memory){
        bytes memory bytesStr = bytes(str);
        uint length = bytesStr.length; // 获取字符串长度
        bytes memory reversedBytes = new bytes(length); // 创建一个新的字节数组
        for (uint i = 0; i < length; i++) {
            reversedBytes[i] = bytesStr[length - 1 - i]; // 反转字符串
        }
        return string(reversedBytes);
    }
    
}