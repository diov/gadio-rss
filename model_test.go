package main

import (
	"encoding/json"
	"testing"
)

const resp = `
{
  "data": [
    {
      "id": "140635",
      "type": "radios",
      "attributes": {
        "title": "忙碌起来！GadioNews08.28",
        "desc": "你觉得《使命召唤：先锋》看上去如何？",
        "thumb": "79a2bbb2-27e8-4adb-bac0-3f5bac57b45a.jpg",
        "published-at": "2021-08-28T23:00:00.000+08:00",
        "duration": 4710
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "45"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5679"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140923",
      "type": "radios",
      "attributes": {
        "title": "模界村 Vol.10 四驱车的今与昔",
        "desc": "你小时候玩过吗？",
        "thumb": "3e830b9f-dd35-4e46-b714-54b8d8717d26.jpg",
        "published-at": "2021-08-27T23:00:00.000+08:00",
        "duration": 4036
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "45"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5663"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140633",
      "type": "radios",
      "attributes": {
        "title": "诱拐案件有隐情，红眼鬼女露狰狞 流行之神 第三话",
        "desc": "人生苦短几人看破，无心之举多少过错",
        "thumb": "d49914c5-244c-4e55-9014-7ec539eb4062.jpg",
        "published-at": "2021-08-27T23:00:00.000+08:00",
        "duration": 4487
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "53"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5661"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140838",
      "type": "radios",
      "attributes": {
        "title": "免费试听集：《燃烧的银河》Episode 1",
        "desc": "战锤·荷鲁斯之乱系列有声书",
        "thumb": "f09558aa-59eb-4125-86da-ff2fe2f718d9.jpg",
        "published-at": "2021-08-27T22:00:00.000+08:00",
        "duration": 2730
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "64"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5607"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140632",
      "type": "radios",
      "attributes": {
        "title": "真人快打故事 卷一 | 倒转时光雷电见前世 地狱归来蝎子欲复仇",
        "desc": "「九代之后的编剧可真是太不容易了」",
        "thumb": "c7d6205e-56a6-470a-9507-4e70d521d13e.jpg",
        "published-at": "2021-08-26T23:00:00.000+08:00",
        "duration": 4041
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "53"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5641"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140631",
      "type": "radios",
      "attributes": {
        "title": "《机核人间观察》复盘：牺牲他人，或是放弃自己",
        "desc": "跑团能让你成为更好的自己么？",
        "thumb": "1aaba57d-d683-4071-a424-751d0db0c84e.jpg",
        "published-at": "2021-08-25T23:00:00.000+08:00",
        "duration": 4385
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "53"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5606"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140628",
      "type": "radios",
      "attributes": {
        "title": "芯片吃紧，主机显卡全靠抢，这些情况到底是怎么造成的？",
        "desc": "这些问题影响了你的消费计划吗？",
        "thumb": "ca7d2423-1c40-4931-8120-d750eb9db230.jpg",
        "published-at": "2021-08-24T23:02:00.000+08:00",
        "duration": 3482
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "12"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5629"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140378",
      "type": "radios",
      "attributes": {
        "title": "细节丰富！GadioNews08.21",
        "desc": "兑现抽奖",
        "thumb": "a2ad88b9-9afe-43ff-bc81-e1286a9e46c5.jpg",
        "published-at": "2021-08-21T23:00:00.000+08:00",
        "duration": 4151
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "45"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5630"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140376",
      "type": "radios",
      "attributes": {
        "title": "中元夜话2021",
        "desc": "如果害怕你就拍拍手",
        "thumb": "07d3a7cf-7752-4c73-b686-52636b74ee2a.jpg",
        "published-at": "2021-08-21T23:00:00.000+08:00",
        "duration": 6469
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "13"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5631"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    },
    {
      "id": "140377",
      "type": "radios",
      "attributes": {
        "title": "二就是二一是一，梦里和谁过七夕 天天ACG 8.20",
        "desc": "说了说EVA，这期也有幻想话题",
        "thumb": "b50c70ff-4c87-4e84-a69a-88454a17fffe.jpg",
        "published-at": "2021-08-20T23:00:00.000+08:00",
        "duration": 4511
      },
      "relationships": {
        "category": {
          "data": {
            "type": "categories",
            "id": "45"
          }
        },
        "media": {
          "data": {
            "type": "medias",
            "id": "5627"
          }
        }
      },
      "meta": {
        "vote-flag": null,
        "vote-id": null,
        "bookmark-id": null
      }
    }
  ],
  "included": [
    {
      "id": "45",
      "type": "categories",
      "attributes": {
        "name": "Gadio News",
        "desc": "",
        "logo": "50d51a9a-846a-4f72-9284-547a637dbdb1.png",
        "background": "de09daa7-5fbd-48c1-98c6-f6326b8da07b.jpg",
        "created-at": "2016-09-29T18:00:31.000+08:00",
        "updated-at": "2020-06-12T16:42:50.000+08:00",
        "subscriptions-count": 11734,
        "scope": "original"
      },
      "relationships": {
        "originals": {},
        "articles": {},
        "videos": {},
        "radios": {},
        "topics": {},
        "latest-video": {}
      },
      "meta": {
        "subscription-id": null,
        "subscription-weight": null
      }
    },
    {
      "id": "53",
      "type": "categories",
      "attributes": {
        "name": "Gadio Story",
        "desc": "",
        "logo": "1e3bcc5f-923f-4188-84c2-2f97cc7a3a38.png",
        "background": "83c591f3-f3a4-4326-a6a8-57bda3500bae.png",
        "created-at": "2018-01-03T16:17:57.000+08:00",
        "updated-at": "2018-01-03T18:25:54.000+08:00",
        "subscriptions-count": 21272,
        "scope": "original"
      },
      "relationships": {
        "originals": {},
        "articles": {},
        "videos": {},
        "radios": {},
        "topics": {},
        "latest-video": {}
      },
      "meta": {
        "subscription-id": null,
        "subscription-weight": null
      }
    },
    {
      "id": "64",
      "type": "categories",
      "attributes": {
        "name": "机核有声书",
        "desc": "",
        "logo": "790796b8-d249-40d2-9269-351902be0d3e.png",
        "background": "91c82d73-ecc2-44c5-8bc2-cbbfc60ec63f.jpg",
        "created-at": "2020-03-15T14:09:47.000+08:00",
        "updated-at": "2021-01-04T17:59:54.000+08:00",
        "subscriptions-count": 5774,
        "scope": "original"
      },
      "relationships": {
        "originals": {},
        "articles": {},
        "videos": {},
        "radios": {},
        "topics": {},
        "latest-video": {}
      },
      "meta": {
        "subscription-id": null,
        "subscription-weight": null
      }
    },
    {
      "id": "12",
      "type": "categories",
      "attributes": {
        "name": "Gadio Pro",
        "desc": "",
        "logo": "51848f93-bd9e-4bb9-ab69-6c5aad92736b.png",
        "background": "8c99662a-9cc9-43be-9053-1a4e374d70d4.jpg",
        "created-at": "2015-02-03T19:21:27.000+08:00",
        "updated-at": "2019-01-23T14:14:36.000+08:00",
        "subscriptions-count": 48447,
        "scope": "original"
      },
      "relationships": {
        "originals": {},
        "articles": {},
        "videos": {},
        "radios": {},
        "topics": {},
        "latest-video": {}
      },
      "meta": {
        "subscription-id": null,
        "subscription-weight": null
      }
    },
    {
      "id": "13",
      "type": "categories",
      "attributes": {
        "name": "Gadio Life",
        "desc": "",
        "logo": "d3968990-b211-4af3-8b01-a4afb2bcd199.png",
        "background": "01bf526b-d6f6-45cc-9f50-e343051d679b.jpg",
        "created-at": "2015-02-03T19:21:27.000+08:00",
        "updated-at": "2019-01-24T15:58:52.000+08:00",
        "subscriptions-count": 22087,
        "scope": "original"
      },
      "relationships": {
        "originals": {},
        "articles": {},
        "videos": {},
        "radios": {},
        "topics": {},
        "latest-video": {}
      },
      "meta": {
        "subscription-id": null,
        "subscription-weight": null
      }
    },
    {
      "id": "5679",
      "type": "medias",
      "attributes": {
        "audio": "3bbf7102-6563-4165-a6c4-9919357ca419.mp3",
        "duration": 4710,
        "title": "新闻0828",
        "original-src": null,
        "created-at": "2021-08-27T16:44:01.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5663",
      "type": "medias",
      "attributes": {
        "audio": "e0bde72b-22fe-42c6-a556-a69b841610fa.mp3",
        "duration": 4036,
        "title": "模界村10 四驱车",
        "original-src": null,
        "created-at": "2021-08-25T15:18:53.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5661",
      "type": "medias",
      "attributes": {
        "audio": "191b497d-9d73-4bc2-8bd3-b5618c2c7126.mp3",
        "duration": 4487,
        "title": "流行之神 第三期",
        "original-src": null,
        "created-at": "2021-08-25T05:17:19.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5607",
      "type": "medias",
      "attributes": {
        "audio": "78b2a587-20b2-4b7f-8d7a-a26e8a69ffda.mp3",
        "duration": 2730,
        "title": "燃烧的银河-1（免费试听集）8.27 22:00更新",
        "original-src": null,
        "created-at": "2021-08-18T12:54:34.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5641",
      "type": "medias",
      "attributes": {
        "audio": "d304feab-fc63-43e9-a6c1-a3e55de09d9b.mp3",
        "duration": 4041,
        "title": "戏说真人快打 02 期",
        "original-src": null,
        "created-at": "2021-08-23T22:28:53.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5606",
      "type": "medias",
      "attributes": {
        "audio": "4467953f-d6da-4fbd-9ce7-717eb500a8df.mp3",
        "duration": 4385,
        "title": "机核人间观察复盘",
        "original-src": null,
        "created-at": "2021-08-17T19:10:14.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5629",
      "type": "medias",
      "attributes": {
        "audio": "f8d25a70-0c6b-4c82-8591-4727e6394a2b.mp3",
        "duration": 3482,
        "title": "芯片屏幕",
        "original-src": null,
        "created-at": "2021-08-20T09:15:51.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5630",
      "type": "medias",
      "attributes": {
        "audio": "fb606608-9f02-413e-9778-3f85ba83af9e.mp3",
        "duration": 4151,
        "title": "新闻0821",
        "original-src": null,
        "created-at": "2021-08-20T19:15:11.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5631",
      "type": "medias",
      "attributes": {
        "audio": "aebe194f-cf56-4413-a340-78171e758a1f.mp3",
        "duration": 6469,
        "title": "中元夜话2021",
        "original-src": null,
        "created-at": "2021-08-21T01:36:01.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    },
    {
      "id": "5627",
      "type": "medias",
      "attributes": {
        "audio": "edb8e7ae-2edb-493d-a49a-3f0653f0b2c0.mp3",
        "duration": 4511,
        "title": "天天ACG 8月第二期",
        "original-src": null,
        "created-at": "2021-08-19T16:06:33.000+08:00",
        "media-type": "audio",
        "process-state": "success",
        "playlist": null
      }
    }
  ],
  "meta": {
    "record-count": 30000
  }
}
`

func TestUnMarshall(t *testing.T) {
	var response Response
	err := json.Unmarshal([]byte(resp), &response)
	if nil != err {
		t.Error(err)
		return
	}
	for _, datum := range response.Data {
		t.Log(datum)
	}
}

func TestValidateUrl(t *testing.T) {
	str := "f09558aa-59eb-4125-86da-ff2fe2f718d9.jpg"
	valid := isUrlValid(str)
	t.Log(valid)
}
