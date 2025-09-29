// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

//在一个有序数组中查找目标值。
contract SelTarget {
    function selTarget(uint256[] memory a,uint256 t) public pure returns (uint){
        uint256 left = 0;
        uint256 right = a.length-1;
        while(left<=right){
            uint256 mid = (left+right)/2;
            if(a[mid]==t){
                return mid;
            }else if(a[mid]>t){
                right = mid-1;
            }else{
                left = mid+1;
            }
        }
        return 0;
    }
}