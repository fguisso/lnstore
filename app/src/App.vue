<template>
    <div id="app">
        <section class="hero is-primary is-fullheight">
            <div class="hero-head">
                <nav class="navbar">
                    <div class="container">
                        <div class="navbar-brand">
                            <a class="navbar-item" href="#">
                                LNstore
                            </a>
                        </div>
                        <div class="navbar-menu">
                            <div class="navbar-end">
                                <span class="navbar-item">
                                    <a class="button is-white is-outlined" href="#">
                                        Github
                                    </a>
                                </span>
                                <span class="navbar-item">
                                    <router-link class="button is-white is-outlined" to="/about">About</router-link>
                                </span>
                                <span class="navbar-item">
                                    <router-link class="button is-white is-outlined" to="/">Home</router-link>
                                </span>
                            </div>
                        </div>
                    </div>
                </nav>
            </div>
            <div class="hero-body">
                <div class="container ">
                    <div class="column is-10 is-offset-1 has-text-centered">
                        <h1 class="title">
                            A thunder of nutrients!
                        </h1>
                        <h2 class="subtitle">
                            Order a bowl of açai and pay instantly.
                        </h2>
                    </div>
                    <div class="column is-10 is-offset-1">
                        <div class="box">
                            <router-view/>
                            <OrderController/>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    </div>
</template>

<script>
import OrderController from '@/components/OrderController.vue'

export default {
  name: 'App',
  components: {
    OrderController
  }
}
</script>
