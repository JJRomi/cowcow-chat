# cowcow-chat


# TODO
[] subscribe channel (create channel)
[] publish channel (key : number -> sorting )
[] publish data save (use queue?)
[] publish data send
[] publish data receive
[] subscribe channel last key save 
    - 이건 나중에 하기 
        -> 이걸 하게되면 client connection이 될 때 바로 last key를 전송해서 마지막 key 다음 data로 저장이 되어야함

- topic: channel
- message
- producer: publish topic 
- consumer: subscribe topic - offset 위치 기억 (fail over 신뢰도 향상)

- create topic "topicName"
    add topic list
- subscribe topic "topicName"
    add consumer
- publish topic "message"
    add topic message
    call consumer list
    send message
    receive message
    save consumer last key