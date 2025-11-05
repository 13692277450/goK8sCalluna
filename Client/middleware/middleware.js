let express = require('express');
let app = express();

app.get('/getData', function (req, res) {
 const jsonData =[{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York1',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston1',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago1',
    score: 33
 },
{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York2',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston2',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago2',
    score: 33
 },
{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York3',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston3',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago4',
    score: 33
 },
{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York4',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston4',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago4',
    score: 33
 },
{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York5',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston5',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago5',
    score: 33
 },
{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York6',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston6',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago6',
    score: 33
 },
{
    ids: 'no1',
    name: 'John',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'New York7',
    score: 11
 },
 {
    ids: 'no2',
    name: 'Frank',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Boston7',
    score: 22
 },
 {
    ids: 'no3',
    name: 'Tiffany',
    img: 'https://bpic.588ku.com/element_origin_min_pic/23/07/11/d32dabe266d10da8b21bd640a2e9b611.jpg!r650',
    key: 'Chicago7',
    score: 33
 },

]
  res.send(jsonData);
});

app.listen(5678, function () {
  console.log('Example app listening on port 5678!');
});