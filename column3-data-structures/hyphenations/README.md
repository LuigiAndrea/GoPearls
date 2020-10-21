#### 3.7.5

This problem deals with a small part of the problem of hyphenating English words. The following list of rules describes some legal hyphenations of words that end in the letter “c”:
et-ic al-is-tic s-tic p-tic -lyt-ic ot-ic an-tic n-tic c-tic at-ic h-nic n-ic m-ic l-lic b-lic -clic l-ic h-ic f-ic d-ic -bic a-ic -mac i-ac
The rules must be applied in the above order; thus the hyphenations “eth-nic” (which is caught by the rule “h-nic”) and “clinic” (which fails that test and falls through to “n-ic”). How would you represent such rules in a function that is given a word and must return suffix hyphenations?