query {
  VoteList{
    id
    name
    details
    logoUrl
    score
  }
}
# ใช้สำหรับดูผลโหวต


mutation {
  CreateVote(input:{name:"ก้าวคนละก้าว",details:"รักหี",logoUrl:"https://google.com/favicon.ico"}){
    success
  }
  # ใช้สำหรับสร้าง vote ให้เลือก
  UpdateVote(input:{name:"ก้าวคนละก้าว",details:"รักหี",logoUrl:"https://google.com/favicon.ico"}){
    success
  }
  # ใช้สำหรับอัพเดตโหวต
}

mutation {
  DeleteVote(input:{id:2}){
    success
  }
  # ใช้สำหรับลบโหวต
}

mutation {
  voteSelect(input:{id:1,ID_CARD_NUMBER:1234,STUDENT_NUMBER:1234}){
    success
  }
  # ใช้สำหรับเลือกโหวต
}

# ws
ใช้สำหรับดูผลคะแนนแบบ realtime
go run ws/ws.go
localhost:8080/ws

#กุโสดโปรจีบ