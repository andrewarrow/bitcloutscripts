
# sample clout.gv file
https://gist.github.com/andrewarrow/7a0b1144c0b4806c521c9094933ea3c5

Only about 35k lines, not the full data.

![image1](https://i.imgur.com/5ZqYYjn.png)

![image2](https://i.imgur.com/9MI81J5.png)

Commands tried:

```
   dot -Tsvg clout.gv > output.svg
   sfdp -x -Tsvg clout.gv > clout.svg
   sfdp -Tsvg clout.gv > clout.svg
   sfdp -x -Goverlap=scale -Tsvg clout.gv > clout.svg
   sfdp -x -Goverlap=scale -Tpng clout.gv > clout.png
   sfdp -x -Tpng clout.gv > clout.png
```
