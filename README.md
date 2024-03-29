## テーブル
- users

- stampcards

- stamps

- notice

## API

### ☑️✅POST  /signup
新規ユーザー作成

req
```
{
  "username": "username",
  "email": "hoge@example.com",
  "password": "hashedPass",
}
```
res（201）
```
{
  "token": "",
}
```

### ☑️✅POST  /login
ログイン

req
```
{
  "email": "hoge@example.com",
  "password": "hashedPass",
}
```
res（200）
```
{
  "token": "",
}
```

### ☑️✅POST  /check-email
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
### ☑️✅GET  /user
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
### ☑️✅PUT  /user
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
### ☑️✅PUT  /user/pwd
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
res（403）
```
なし
```

### ☑️✅POST  /stampcard
スタンプカードの作成

req
```
{
    "title": "title",
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
  "id": ""
}
```
### GET  /stampcard?query={keyword}
スタンプカード一覧取得withクエリ絞り込み(要相談)

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

###  ☑️✅GET  /stampcard/:id
スタンプカードの取得

req
```
なし
```

res（200）
```
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

### PUT /stampcard/delete/:id
スタンプカードの削除

req
```
なし
```

res（204）
```
なし
```


### PUT  /stamp
スタンプの作成（取得はスタンプカードと同時に、更新、削除処理はなし）

req
```
{
  "stamp": "🌟",
  "message": "おめでとう！",
  "nthDay": 1,
  "cardId": 1,
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
通知の作成
isStampy == trueならスタンプと通知（スタンプが届きました）  
それ以外なら通知（スタンプを要求されています）←もうちょい良い言い方あるかも

req
```
{
  StampId: 1
}
```

res（201）
```
{
  type: "notification",
  id: "1",
  title: "",
  stamp: "🌟",
  message: "",
  hrefPrefix: "",
  sender: 2,
  receiver: 1,
  read: false,
  createdAt: "2024-02-22 00:00:00",
  listType: "",
}
```
### POST  /letter
レターの作成(同時に通知も作成)


req
```
{
  StampId: 1
}
```

res（201）
```
{
  type: "letter",
  id: "1",
  title: "titleへの完走レター",
  stamp: "🌟",
  message: "",
  hrefPrefix: "/letter",
  sender: 2,
  receiver: 1,
  read: false,
  createdAt: "2024-02-22 00:00:00",
  listType: "link",
}
```

### GET  /notice
通知の一覧取得withクエリ絞り込み(要相談)

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
    title: "テキストアイテム",
    stamp: "🌟",
    message:
      "テキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキスト",
    hrefPrefix: "/letter",
    sender: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1548142813-c348350df52b?&w=150&h=150&dpr=2&q=80",
    },
    receiver: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1531384441138-2736e62e0919?&w=100&h=100&dpr=2&q=80",
    },
    read: true,
    createdAt: "createdAt",
    sendAt: "sendAt",
    listType: "text",
  },
  {
    type: "notification",
    id: "2",
    title: "最終日ダイアログ",
    message:
      "テキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキスト",
    stamp: "stamp",
    currentDay: 10,
    isLastDay: true,
    hrefPrefix: "hrefPrefix",
    sender: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1548142813-c348350df52b?&w=150&h=150&dpr=2&q=80",
    },
    receiver: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1531384441138-2736e62e0919?&w=100&h=100&dpr=2&q=80",
    },
    read: false,
    createdAt: "createdAt",
    sendAt: "sendAt",
    listType: "sender-dialog",
  }
  ]
}
```
### GET  /notice/:id
通知取得

req
```
なし
```

res（200）
```
  {
    type: "notifcation",
    id: "2",
    title: "「カードタイトル」への完走レター",
    message:
      "テキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキスト",
    stamp: "🌟",
    hrefPrefix: "",
    sender: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1548142813-c348350df52b?&w=150&h=150&dpr=2&q=80",
    },
    receiver: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1531384441138-2736e62e0919?&w=100&h=100&dpr=2&q=80",
    },
    read: false,
    createdAt: "createdAt",
    listType: "",
  }

```



### GET  /letters
レターの一覧取得withクエリ絞り込み(要相談)

req
```
なし
```

res（200）
```
{
  letters: [
  {
    type: "letter",
    id: "1",
    title: "「カードタイトル」への完走レター",
    stamp: "🌟",
    message:
      "テキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキスト",
    hrefPrefix: "/letter",
    sender: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1548142813-c348350df52b?&w=150&h=150&dpr=2&q=80",
    },
    receiver: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1531384441138-2736e62e0919?&w=100&h=100&dpr=2&q=80",
    },
    read: true,
    createdAt: "createdAt",
    sendAt: "sendAt",
    listType: "link",
  },
  {
    type: "letter",
    id: "2",
    title: "「カードタイトル」への完走レター",
    message:
      "テキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキスト",
    stamp: "🌟",
    hrefPrefix: "/letter",
    sender: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1548142813-c348350df52b?&w=150&h=150&dpr=2&q=80",
    },
    receiver: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1531384441138-2736e62e0919?&w=100&h=100&dpr=2&q=80",
    },
    read: false,
    createdAt: "createdAt",
    sendAt: "sendAt",
    listType: "link",
  }
]
}
```
### GET  /letters/:id
レター取得

req
```
なし
```

res（200）
```
  {
    type: "letter",
    id: "2",
    title: "「カードタイトル」への完走レター",
    message:
      "テキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキストテキスト",
    stamp: "🌟",
    hrefPrefix: "/letter",
    sender: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1548142813-c348350df52b?&w=150&h=150&dpr=2&q=80",
    },
    receiver: {
      id: "1",
      username: "username",
      email: "email",
      avatarUrl:
        "https://images.unsplash.com/photo-1531384441138-2736e62e0919?&w=100&h=100&dpr=2&q=80",
    },
    read: false,
    createdAt: "createdAt",
    sendAt: "sendAt",
    listType: "link",
  }

```


### PUT  /notice/read/:id
通知の既読

req
```
なし
```

res
```
なし
```

### GET  /user/total
ユーザーがこれまでに受け取ったスタンプ、チャレンジしたカード、完了したカード、受け取ったレター、stampyを通してつながった人の数を返す
req
```
なし
```

res
```
{
  receivedStamp: 1, // isStampedがtrueであるstampのcount
  challengeCard: 1, // createdUserがuserIdと一致するcardのcount
  completedCard: 1, // ↑+isCompletedがtrueであるcardのcount
  receivedLetter: 1, // ReceiverがuserIdと一致するletterのcount
  people: 1, // できたらやる
}

```