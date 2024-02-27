## テーブル
- users

- stampcards

- stamps

- notice

## API

### POST  /signup
新規ユーザー作成

req
```
{
  "username": "username",
  "email": "hoge@example.com",
  "hashedPassword": "hashedPass",
}
```
res（201）
```
{
  "username": "",
  "email": "",
  "token": "",
}
```

### POST  /login
ログイン

req
```
{
  "email": "hoge@example.com",
  "hashedPassword": "hashedPass",
}
```
res（200）
```
{
  "token": "",
}
```

## これ以下はすべてtoken必要
### GET  /user/:id
ユーザー情報取得

req
```
なし
```
res（200）
```
{
  "id": "",
  "username": "",
  "email": "",
  "avatarUrl": "",
}
```
### PUT  /user/:id
ユーザー情報の更新(パスワード以外)  
req
```
{
  "username": "",
  "email": "",
  "avatarUrl": "",
}
```
res（204）
```
なし
```

### PUT  /user/:id/pwd
ユーザー情報の更新  
req
```
{
  "oldPass": "",
  "newPass": "",
}
```
res（204）
```
なし
```

### POST  /stampcard
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

res（201）
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
### GET  /stampcard
スタンプカード取得

req
```
なし
```

res（200）
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

### PUT  /stampcard/:id
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

res（204）
```
なし
```
### POST  /stamp
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

res（201）
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

### POST  /notice
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

res（201）
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
### GET  /notice
通知/レターの一覧取得

req
```
なし
```

res（200）
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
認証用, emailとpasswordを受け取ってjwtを返す、jwtにuseridを含める

## その他
