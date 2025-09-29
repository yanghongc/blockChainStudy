// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

/**
    创建一个名为Voting的合约，包含以下功能：
    一个mapping来存储候选人的得票数
    一个vote函数，允许用户投票给某个候选人
    一个getVotes函数，返回某个候选人的得票数
    一个resetVotes函数，重置所有候选人的得票数
*/
contract Voting {
    //一个mapping来存储候选人的得票数
    mapping (string => uint256) private votes;
    mapping (string => bool) private isVote;
    //设置候选人
    string[] private votesKey; 
    function setVotes(string memory candidata)public {
        votes[candidata] = 0;
        isVote[candidata] = true;
        votesKey.push(candidata);
    }
    //一个vote函数，允许用户投票给某个候选人
    function vote(string memory candidata) public {
        require(isVote[candidata],"candidata is not exist!");
        votes[candidata] += 1;
    }
    //一个getVotes函数，返回某个候选人的得票数
    function getVotes(string memory candidata) public view returns (uint256){
        require(isVote[candidata],"candidata is not exist!");
        return votes[candidata];
    }
    //一个resetVotes函数，重置所有候选人的得票数
    function resetVotes()public {
        for (uint i = 0; i < votesKey.length; i++){
            votes[votesKey[i]] = 0;
        }
    }
}