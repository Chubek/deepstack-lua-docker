#!/bin/bash


cd ~/DeepHoldem/ACPCServer
./dealer testMatch holdem.nolimit.2p.reverse_blinds.game 1000 0 Alice Bob > /home/logs/log_deaer.log 
cd Source
th th Player/manual_player.lua 2984 > /home/logs/log_manual.log &  th Player/deepstack.lua 2985 > /home/logs/log_deepstack.log