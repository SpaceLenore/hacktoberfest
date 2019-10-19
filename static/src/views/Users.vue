<template>
    <div class="component">
        <div class="header">
            <h1>Hacktoberfest Hackathon</h1>
        </div>
        <div class="addbox" v-if="showUserBox">
            <span class="close-addbox" @click="showUserBox=false">[x]</span>
            <h2 class="addbox-title">Add user</h2>
            <input class="addbox-input" type="text" placeholder="GitHub username" v-model="user" @keyup.enter="addUsername">
        </div>
        <div v-else @click="showUserBox=true" class="addUserLink">Add users</div>
        <div class="users">
            <div v-for="{username, count} in users" v-bind:key="username"> {{username}}: {{count}}</div>
        </div>
    </div>
</template>

<script>

export default {
    data() {
        return {
            api: "http://193.11.184.114:5000/api/",
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
            console.log("List updated");
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
    html,
    body {
        width: 100%;
        margin: 0;
        padding: 0;
        background: #152347;        
    }

    .component {
        width: 100%;
        color: #fff;
    }

    .header h1 {
        font-size: 2.5rem;
        font-style: italic;
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
        border: none;
        padding: 1rem;
        font-size: 1.2rem;
    }
</style>