<template>
    <div class="component">
        <div class="header">
            <img src="@/assets/Hacktoberfest_19.png" alt="Hacktoberfest logo" height="150px">
            <h1>Hackathon Leaderboard</h1>
        </div>
        <div class="addbox" v-if="showUserBox">
            <span class="close-addbox" @click="showUserBox=false">[x]</span>
            <h2 class="addbox-title">Add user</h2>
            <input class="addbox-input" type="text" placeholder="GitHub username" v-model="user" @keyup.enter="addUsername">
        </div>
        <div v-else>
            <a @click="showUserBox=true" class="addUserLink">Add users</a>
        </div>
        <div class="users" v-if="users.length > 0">
            <div class="user" v-for="{username, count} in users" v-bind:key="username">
                <span class="username">{{username}}</span>
                <span class="count"> {{count}}</span>
            </div>
        </div>
    </div>
</template>

<script>

export default {
    data() {
        return {
            api: process.env.API_ENDPOINT,
            users: [],
            user: null,
            showUserBox: true
        }
    },
    methods: {
        addUsername() {
            this.users.push({"username": this.user, "count": 0})
            this.user = null;
            this.updateUsers();
        },

        updateUsers() {
            for (let i = 0; i < this.users.length; i++) {
                fetch(this.api + this.users[i].username)
                .then((response) => {
                    return response.json()
                })
                .then((data) => {
                    this.users[i].count = data.count;
                })
            }
            // TODO: SORT LIST BY MOST PR
        }
    },
    created() {
        setInterval(() => {
            if (this.users.length > 0) {
                this.updateUsers();
            }
        }, 60000)
    }
}
</script>

<style>
    @import url('https://fonts.googleapis.com/css?family=Space+Mono&display=swap');

    html,
    body {
        width: 100%;
        margin: 0;
        padding: 0;
        background: #152347;    
        font-family: 'Space Mono', monospace;    
    }

    .component {
        width: 100%;
        color: #fff;
    }

    .header h1 {
        margin: 0;
        font-size: 2.5rem;
        font-weight: 300;
    }

    .addUserLink {
        cursor: pointer;
        font-size: 1.1rem;
    }

    .addbox {
        position: relative;
        margin: 2rem auto;
        padding: 2rem;
        width: 40%;
        border: 1px solid #8d1272;
        box-shadow: 0 0 8px #ff00aa;
    }

    .close-addbox {
        position: absolute;
        top: 0.5rem;
        right: 1rem;
        cursor: pointer;
    }

    .addbox-title {
        margin: 0 auto 1rem auto;
    }

    .addbox-input {
        background-color: #39486e;
        padding: 1rem;
        font-size: 1.2rem;
        border: 1px solid #39486e;;
    }

    .addbox-input:focus {
        border-bottom: 1px solid #fff922;
    }

    .addUserLink {
        color: #ff00aa;
        margin: 1rem 0;
    }

    .addUserLink:hover {
        text-decoration: underline;
    }

    .users {
        width: 80%;
        margin: 2rem auto;
        display: flex;
        flex-direction: column;
        border-bottom: 1px solid rgba(255, 255, 255, 0.4);
        border-top: 1px solid rgba(255, 255, 255, 0.4);
    }

    .user:nth-child(odd) {
        background: #39486e
    }

    .user {
        display: flex;
        justify-content: space-between;
        font-size: 1.5rem;
        padding: 0.8rem 0;
    }

    .username {
        padding: 0 1rem;
    }

    .count {
        padding: 0 1rem;
    }
</style>