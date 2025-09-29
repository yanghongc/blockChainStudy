// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInt {
    mapping ( string => uint ) public romanValues;

    constructor (){
        romanValues["I"] = 1;
        romanValues["V"] = 5;
        romanValues["X"] = 10;
        romanValues["L"] = 50;
        romanValues["C"] = 100;
        romanValues["D"] = 500;
        romanValues["M"] = 1000;
    }
    function romanToInt(string memory s) public view returns (uint) {
        bytes memory str = bytes(s);
        uint sum = 0;
        //获取第一个数
        uint preNum = romanValues[string(abi.encodePacked(str[0]))]; 
        //循环比较
        for (uint i = 1; i < str.length; i++) {
            //获取当前数
            uint curNum = romanValues[string(abi.encodePacked(str[i]))];
            //如果前一个数比当前数小，则当前数减去前一个数
            if (preNum < curNum) {
                sum -= preNum;
            } else {
                //否则直接加上前一个数
                sum += preNum;
            }
            //更新前一个数
            preNum = curNum;
       }

       sum += preNum;

       return sum;
        
    }

}