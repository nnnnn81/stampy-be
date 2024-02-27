## テーブル
- users

- stampcards

- stamps

- notice

## エンドポイント(フロントから叩く)

### /signup

`/signup`  
新規ユーザー作成

req
```
{
  "username": "username",
  "email": "hoge@example.com",
  "hashedPass": "hashedPass",
}
```
res
```
{
  "username": "",
  "email": "",
  "token": "",
}
```

### /login
`/login`  
ログイン


req
```
{
  "email": "hoge@example.com",
  "hashedPass": "hashedPass",
}
```
res
```
{
  "token": "",
}
```

## これ以下はすべてtoken必要
### /user
`/user`  
ユーザー情報取得

req
```
なし
```
res
```
{
  "username": "",
  "email": "",
  "avatarUrl": "",
}
```
`/user`  
ユーザー情報の更新(パスワード以外)

req
```
{
  "username": "",
  "email": "",
  "avatarUrl": "",
}
```
res
```
{
  "username": "",
  "email": "",
  "avatarUrl": "",
}
```

`/user/pwd`  
ユーザー情報の更新

req
```
{
  "oldPass": "",
  "newPass": "",
}
```
res
```
なし
```

### /stampcard
`/stampcard`  
スタンプカードの作成

req
```
{
    "title": "title",
    "createdBy": "createUser@gmail.com",
    "joinedUser": "joinedUser@gmail.com",
    "backgroundUrl": "",
}
```

res
```
{
    title: `title`,
    createdBy: {
      username: "username",
      email: "createUser@gmail.com",
      avatarUrl: "",
    },
    joinedUser: {
      username: "username",
      email: "joinedUser@gmail.com",
      avatarUrl: "",
    },
    createdAt: "2024-02-22 00:00:00",
    updatedAt: "2024-02-22 00:00:00",
    currentDay: 0,
    isCompleted: 0,
    isDeleted: false,
    stampNodes: [],
    backgroundUrl: "https://source.unsplash.com/ZkOt0N7rP4s",
}
```
`/stampcards`  
スタンプカード取得

req
```
なし
```

res
```
{
  "cards": [
    {
      "id": "",
      "title": `title`,
      "createdBy": {
        "username": "username",
        "email": "createUser@gmail.com",
        "avatarUrl": "",
      },
      "joinedUser": {
        "username": "username",
        "email": "joinedUser@gmail.com",
        "avatarUrl": "",
      },
      "createdAt": "2024-02-22 00:00:00",
      "updatedAt": "2024-02-22 00:00:00",
      "currentDay": 0,
      "isCompleted": 0,
      "isDeleted": false,
      "stampNodes": [
        {
          "id": "",
          "stamp": "🌟",
          "message": "おめでとう！",
          "stamped": true,
          "stampedAt": "2024-02-22 00:00:00",
          "nthDay": i + 1,
          "stampedBy": {
            "id": "1",
            "username": "山田太郎",
            "email": "email.com",
            "avatarUrl": "",
          },
          "x": 0,
          "y": 0,
          "cardId": "",
        },
        {
          "id": "",
          "stamp": "🌟",
          "message": "おめでとう！",
          "stamped": true,
          "stampedAt": "2024-02-22 00:00:00",
          "nthDay": i + 1,
          "stampedBy": {
            "id": "1",
            "username": "山田太郎",
            "email": "email.com",
            "avatarUrl": "",
          },
          "x": 0,
          "y": 0,
          "cardId": "",
        }
      ],
      "backgroundUrl": "https://source.unsplash.com/ZkOt0N7rP4s",
    },
    {
      "id": "",
      "title": `title`,
      "createdBy": {
        "username": "username",
        "email": "createUser@gmail.com",
        "avatarUrl": "",
      },
      "joinedUser": {
        "username": "username",
        "email": "joinedUser@gmail.com",
        "avatarUrl": "",
      },
      "createdAt": "2024-02-22 00:00:00",
      "updatedAt": "2024-02-22 00:00:00",
      "currentDay": 0,
      "isCompleted": 0,
      "isDeleted": false,
      "stampNodes": [],
      "backgroundUrl": "https://source.unsplash.com/ZkOt0N7rP4s",
    }
  ]
}
```

`/stampcard/:id`  
スタンプカードの更新

req
```
{
  "id": "",
  "title": `title`,
  "currentDay": 0,
  "isCompleted": 0,
  "isDeleted": false,
  "stampNodes": [],
  "backgroundUrl": "https://source.unsplash.com/ZkOt0N7rP4s",
}
```

res　204 No Content
```
なし
```
### /stamp
`/stamp`  
スタンプの作成（取得はスタンプカードと同時に、更新、削除処理はなし）

req
```
{
  "stamp": "🌟",
  "message": "おめでとう！",
  "nthDay": 1,
  "stampedBy": {
    "username": "山田太郎",
    "email": "email.com",
    "avatarUrl": "",
  },
  "x": 0,
  "y": 0,
  "cardId": "",
}
```

res
```
{
  "stamp": "🌟",
  "message": "おめでとう！",
  "nthDay": 1,
  "stampedBy": {
    "username": "山田太郎",
    "email": "email.com",
    "avatarUrl": "",
  },
  "x": 0,
  "y": 0,
  "cardId": "",
}
```

### /letter
`/notice`  
通知/レターの作成

req
```
{
  type: "notification",
  title: "",
  stamp: "🌟",
  content: "",
  hrefPrefix: "",
  read: true,
  createdAt: "createdAt",
  sendAt: "sendAt",
  listType: "text",
}
```

res 201
```
{
  type: "notification",
  id: "1",
  title: "",
  stamp: "🌟",
  content: "",
  hrefPrefix: "/letter",
  sender: {
    id: "1",
    username: "username",
    email: "email",
    avatarUrl: "",
  },
  receiver: {
    id: "1",
    username: "username",
    email: "email",
    avatarUrl: "",
  },
  read: true,
  createdAt: "createdAt",
  sendAt: "sendAt",
  listType: "text",
}
```
- `/notice`  
通知/レターの取得

req
```
なし
```

res
```
{
  notice: [
    {
      type: "notification",
      id: "1",
      title: "",
      stamp: "🌟",
      content: "",
      hrefPrefix: "/letter",
      sender: {
        id: "1",
        username: "username",
        email: "email",
        avatarUrl: "",
      },
      receiver: {
        id: "1",
        username: "username",
        email: "email",
        avatarUrl: "",
      },
      read: true,
      createdAt: "createdAt",
      sendAt: "sendAt",
      listType: "text",
    },
    {
      type: "notification",
      id: "1",
      title: "",
      stamp: "🌟",
      content: "",
      hrefPrefix: "/letter",
      sender: {
        id: "1",
        username: "username",
        email: "email",
        avatarUrl: "",
      },
      receiver: {
        id: "1",
        username: "username",
        email: "email",
        avatarUrl: "",
      },
      read: true,
      createdAt: "createdAt",
      sendAt: "sendAt",
      listType: "text",
    }
  ]
}
```




## バックエンドで考えること

### /user
- パスワード更新だけわけたい

### /auth
- `/auth`  
認証用, emailとpasswordを受け取ってjwtを返す

### /stampcard


### /stamp
- `/stamp`  
スタンプの作成（取得はスタンプカードと同時に、更新、削除処理はなし）

### /letter
- `/letter`  
レターの作成
- `/letter/:id`  
レターの詳細取得

## その他
