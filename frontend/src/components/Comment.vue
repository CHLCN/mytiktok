<template>
  <from-bottom-dialog
    :page-id="pageId"
    :modelValue="modelValue"
    @update:modelValue="(e) => $emit('update:modelValue', e)"
    @cancel="cancel"
    :show-heng-gang="false"
    maskMode="light"
    :height="height"
    tag="comment"
    mode="white"
  >
    <template v-slot:header>
      <div class="title">
        <dy-back mode="dark" img="close" direction="right" style="opacity: 0" />
        <span v-if="comments">{{ comments.length }}条评论</span>
        <dy-back
          mode="dark"
          img="close"
          direction="right"
          @click.stop="cancel"
        />
      </div>
    </template>
    <div class="comment">
      <div class="wrapper" >
        <div class="items">
          <div class="item" v-for="item in comments">
            <div class="main">
              <div class="content">
                <img :src="item.user.avatar" alt="" class="head-image" />
                <div class="comment-container">
                  <div class="name">{{ item.user.nickname }}</div>
                  <div class="detail">{{ item.content }}</div>
                  <div class="time-wrapper">
                    <div class="left">
                      <div class="time">
                        {{ $time(item.create_date) }} · 安徽
                      </div>
                      <div class="reply-text">回复</div>
                    </div>
                    <div class="love" @click="loved(item)">
                      <img
                        v-show="item.isLoved"
                        src="../assets/img/icon/components/like-red-small.png"
                        alt=""
                        class="love-image"
                      />
                      <img
                        v-show="!item.isLoved"
                        src="../assets/img/icon/components/like-gray-small.png"
                        alt=""
                        class="love-image"
                      />
                      <span>{{ formatNumber(item.loveNum) }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <no-more />
      </div>

      <transition name="fade">
        <Mask v-if="isCall" mode="lightgray" @click="isCall = false" />
      </transition>
      <div class="input-toolbar">
        <transition name="fade">
          <div class="call-friend" v-if="isCall">
            <div
              class="friend"
              v-for="item in friends.user_list"
              @click="toggleCall(item)"
            >
              <img
                :style="item.select ? 'opacity: .5;' : ''"
                class="avatar"
                :src="$imgPreview(item.avatar)"
                alt=""
              />
              <span>{{ item.nickname }}</span>
              <img
                v-if="item.select"
                class="checked"
                src="../assets/img/icon/components/check/check-red-share.png"
              />
            </div>
          </div>
        </transition>

        <div class="toolbar">
          <div class="input-wrapper">
            <AutoInput v-model="comment" class="myinput"></AutoInput>
            <div class="right">
              <img
                src="../assets/img/icon/message/call.png"
                @click="isCall = !isCall"
              />
              <img src="../assets/img/icon/message/emoji-black.png" />
            </div>
          </div>
          <img
            v-if="comment"
            src="../assets/img/icon/message/up.png"
            @click="send"
          />
        </div>
      </div>
      <ConfirmDialog
        title="私信给"
        ok-text="发送"
        v-model:visible="showPrivateChat"
      >
        <Search mode="light" v-model="test" :isShowSearchIcon="false" />
      </ConfirmDialog>
    </div>
  </from-bottom-dialog>
</template>

<script>
import AutoInput from "./AutoInput";
import ConfirmDialog from "./dialog/ConfirmDialog";
import { mapState } from "vuex";
import FromBottomDialog from "./dialog/FromBottomDialog";
import Loading from "./Loading";
import Search from "./Search";
import request from "../utils/request";
import bus, { EVENT_KEY } from "@/utils/bus";

export default {
  name: "Comment",
  components: {
    AutoInput,
    ConfirmDialog,
    FromBottomDialog,
    Loading,
    Search,
  },
  props: {
    modelValue: false,
    videoId: {
      type: String,
      default: null,
    },
    pageId: {
      type: String,
      default: "home-index",
    },
    height: {
      type: String,
      default: "70vh",
    },
  },
  computed: {
    // ...mapState(["friends"]),
  },
  watch: {
    modelValue(newVale) {
      if (newVale) {
        this.getData();
      } else {
        this.comments = [];
      }
    },
  },
  data() {
    return {
      friends:{},
      comment: "",
      test: "",
      comments: [],
      options: [
        { id: 1, name: "私信回复" },
        { id: 2, name: "复制" },
        { id: 3, name: "搜一搜" },
        { id: 4, name: "举报" },
      ],
      selectRow: {},
      showPrivateChat: false,
      isInput: false,
      isCall: false,
      videoId: 0,
    };
  },
  mounted() {
    bus.on(EVENT_KEY.OPEN_COMMENTS, this.getID);
    this.getData();
  },
  methods: {
    getID(itemId) {
      this.videoId = itemId;
    },
    async send() {
      let res = await request.post(
          "/comment/action/",
          {
          },
          {
            params: {
              user_id: this.$store.state.user_id,
              video_id: this.videoId,
              action_type:1,
              comment_text:this.comment,
            }
          }
      );
      console.log(res);
      let updateres = await request.get(
          "/comment/list/",
          {
            params: {
              user_id: this.$store.state.user_id,
              video_id: this.videoId,
            },
          },
          {}
      );
      this.comments = updateres.data.comment_list;
    },
    async getData() {
      // await this.$sleep(500);
      let res = await request.get(
        "/comment/list/",
        {
          params: {
            user_id: this.$store.state.user_id,
            video_id: this.videoId,
          },
        },
        {}
      );
      this.comments = res.data.comment_list;
      let friendres = await request.get(
          "/relation/friend/list/",
          {
            params: {
              user_id: this.$store.state.user_id,
            },
          },
          {}
      );
      console.log(friendres)
      this.friends = friendres.data
    },
    cancel() {
      this.$emit("update:modelValue", false);
      this.$emit("close");
    },
    async toggleCall(item) {
      item.select = !item.select;
      let nickname = item.nickname;
      if (this.comment.includes("@" + nickname)) {
        this.comment = this.comment.replace(`@${nickname} `, "");

      } else {
        this.comment += `@${nickname} `;
      }
    },
    loved(row) {
      if (row.isLoved) {
        row.loveNum--;
      } else {
        row.loveNum++;
      }
      row.isLoved = !row.isLoved;
    },
    showOptions(row) {
      this.$showSelectDialog(this.options, (e) => {
        if (e.id === 1) {
          this.selectRow = row;
          this.showPrivateChat = true;
        }
      });
    },
    // showComment() {
    //     this.isCommenting = !this.isCommenting;
    //     console.log(666)
    // }
    call() {
      console.log(this.commit);
    },
  },
};
</script>

<style lang="less" scoped>
@import "../assets/less/index";

.title {
  z-index: 2;
  position: fixed;
  left: 0;
  right: 0;
  height: 40rem;
  padding: 0 15rem;
  background: white;
  display: flex;
  align-items: center;
  justify-content: space-between;
  border-radius: 10rem 10rem 0 0;

  span {
    font-size: 12rem;
    text-align: center;
  }

  img {
    width: 13rem;
    height: 13rem;
  }
}

.comment {
  width: 100vw;
  height: v-bind(height);
  background: #fff;
  z-index: 5;
  border-radius: 10rem 10rem 0 0;

  .wrapper {
    width: 100%;
    position: relative;
    padding-top: 40rem;
    padding-bottom: 60rem;
  }

  .items {
    width: 100%;

    .item {
      width: 100%;

      .main {
        width: 100%;
        padding: 5rem 0;
        display: flex;

        &:active {
          background: #53535321;
        }

        .head-image {
          margin-left: 15rem;
          margin-right: 10rem;
          width: 35rem;
          height: 35rem;
          border-radius: 50%;
        }
      }

      .replies {
        padding-left: 55rem;

        .reply {
          padding: 5rem 0 5rem 5rem;
          display: flex;

          &:active {
            background: #53535321;
          }

          .head-image {
            margin-right: 10rem;
            width: 20rem;
            height: 20rem;
            border-radius: 50%;
          }
        }

        .more {
          font-size: 12rem;
          margin: 5rem;
          display: flex;
          align-items: center;
          color: @second-text-color;

          .gang {
            background: #d5d5d5;
            width: 20rem;
            margin-right: 10rem;
            height: 1px;
          }

          .arrow {
            margin-left: 5rem;
            margin-top: 8rem;
            width: 0;
            height: 0;
            border: 5rem solid transparent;
            border-top: 6rem solid @second-text-color;
          }
        }
      }

      .content {
        width: 100%;
        display: flex;
        font-size: 14rem;

        .comment-container {
          flex: 1;
          margin-right: 20rem;

          .name {
            color: @second-text-color;
            margin-bottom: 5rem;
            display: flex;
            align-items: center;

            .reply-user {
              margin-left: 5rem;
              width: 0;
              height: 0;
              border: 5rem solid transparent;
              border-left: 6rem solid gray;
            }
          }

          .detail {
            margin-bottom: 5rem;
          }

          .time-wrapper {
            display: flex;
            align-items: center;
            justify-content: space-between;
            font-size: 12rem;

            .left {
              display: flex;

              .time {
                color: #c4c3c3;
                margin-right: 10rem;
              }

              .reply-text {
                color: @second-text-color;
              }
            }

            .love {
              color: @second-text-color;
              display: flex;
              align-items: center;

              .love-image {
                width: 18rem;
                border-radius: 50%;
              }

              span {
                font-size: 10rem;
                word-break: keep-all;
              }
            }
          }
        }
      }
    }
  }

  @normal-bg-color: rgb(35, 38, 47);
  @chat-bg-color: rgb(105, 143, 244);

  .input-toolbar {
    border-radius: 10rem 10rem 0 0;
    background: white;
    position: fixed;
    width: 100vw;
    bottom: 0;
    z-index: 3;

    @space-width: 18rem;
    @icon-width: 48rem;

    .call-friend {
      padding-top: 30rem;
      overflow-x: scroll;
      display: flex;
      padding-right: @space-width;

      .friend {
        width: @icon-width;
        position: relative;
        margin-left: @space-width;
        margin-bottom: @space-width;
        font-size: 10rem;
        display: flex;
        flex-direction: column;
        align-items: center;

        .avatar {
          width: @icon-width;
          height: @icon-width;
          border-radius: 50%;
        }

        span {
          margin-top: 5rem;
          text-align: center;
          width: @icon-width;
          white-space: nowrap;
          overflow: hidden;
          text-overflow: ellipsis;
        }

        .checked {
          position: absolute;
          top: @icon-width - 1.5;
          right: -2px;
          width: 20rem;
          height: 20rem;
          border-radius: 50%;
        }
      }
    }

    .toolbar {
      @icon-width: 25rem;
      display: flex;
      align-items: center;
      padding: 10rem 15rem;
      border-top: 1px solid #e2e1e1;

      .input-wrapper {
        flex: 1;
        display: flex;
        align-items: center;
        justify-content: space-between;
        box-sizing: border-box;
        padding: 5rem 10rem;
        background: #eee;
        border-radius: 20rem;

        .right {
          display: flex;
          align-items: center;
        }

        .auto-input {
          width: calc(100vw - 180rem);
        }
      }

      img {
        width: @icon-width;
        height: @icon-width;
        border-radius: 50%;
        margin-left: 15rem;
      }
    }
  }
}

.comment-enter-active,
.comment-leave-active {
  transition: all 0.15s ease;
}

.comment-enter-from,
.comment-leave-to {
  transform: translateY(60vh);
}
</style>
