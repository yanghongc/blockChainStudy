// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSort {

    function mergeSort(uint256[] memory a, uint256[] memory b ) public pure returns(uint256[] memory){
        uint ia = 0;
        uint ib = 0;
        uint ic = 0;
        uint256[] memory c = new uint256[](a.length + b.length);
        //比较a，b数组，将较小者放入到c数组中
        while (ia < a.length && ib < b.length) {
            if (a[ia] < b[ib]) {
                c[ic] = a[ia];
                ia++;
            } else {
                c[ic] = b[ib];
                ib++;
            }
            ic++;
        }
        //如果a中有剩余，将其全放到c中
        while (ia < a.length) {
            c[ic] = a[ia];
            ia++;
            ic++;
        }
        //如果b中有剩余，将其全放到c中
        while (ib < b.length) {
            c[ic] = a[ib];
            ib++;
            ic++;
        }

        return c;
    }
}