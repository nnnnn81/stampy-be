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

### POST  /check-email
登録済みメールアドレスの確認

req
```
{
  "email": "hoge@example.com",
}
```
res（200）
```
{
  "doesUserExist": true
}
```


## これ以下はすべてtoken必要
### GET  /user
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
### PUT  /user
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

### PUT  /user/pwd
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
    "startDate": "2024-02-22 00:00:00",
    "endDate": "2024-02-22 00:00:00",
    "isStampy": false,
    "backgroundUrl": "",
}
```

res（201）
```
{
    title: `title`,
    createdBy: {
      "id": "0",
      username: "username",
      email: "createUser@gmail.com",
      avatarUrl: "",
    },
    joinedUser: {
      "id": "1",
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
        "id": "0",
        "username": "username",
        "email": "createUser@gmail.com",
        "avatarUrl": "",
      },
      "joinedUser": {
        "id": "1",
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
          "cardId": "",
        }
      ],
      "backgroundUrl": "https://source.unsplash.com/ZkOt0N7rP4s",
    },
    {
      "id": "",
      "title": `title`,
      "createdBy": {
        "id": "0",
        "username": "username",
        "email": "createUser@gmail.com",
        "avatarUrl": "",
      },
      "joinedUser": {
        "id": "1",
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
  "title": `title`,
  "currentDay": 0,
  "isCompleted": false,
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
  "cardId": "",
}
```

res（201）
```
{
  "stampId": "",
  "stamp": "🌟",
  "message": "おめでとう！",
  "nthDay": 1,
  "stampedBy": {
    "id": "1",
    "username": "山田太郎",
    "email": "email.com",
    "avatarUrl": "",
  },
  "cardId": "",
}
```

### POST  /notice
レターの作成

req
```
{
  type: "notification",
  title: "",
  stamp: "🌟",
  content: "",
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
  hrefPrefix: "", // letterのとき"/letter"になってほしい、みたいなはなし
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
  listType: "text",　// 4つある、用途としては通知のメタ情報（押したらダイアログがでるとか、ただのテキストとか）
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