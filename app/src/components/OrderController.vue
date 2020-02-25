<template>
    <div class="hello">
        <div class="border-t">
            <nav class="navbar">
                <div class="navbar-menu">
                    <div class="navbar-start">
                        <div class="field is-grouped">
                            <p class="control">
                                <button class="button is-danger">
                                    Back
                                </button>
                            </p>
                            <p class="control">
                                <button class="button is-success">
                                    Next
                                </button>
                            </p>
                        </div>
                    </div>

                    <div class="navbar-end">
                        <div class="navbar-item">
                            <p class="total-count">
                                Total: $DCR 0.000100
                            </p>
                        </div>
                    </div>
                </div>
            </nav>
        </div>
    </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  }
}
</script>

<style>
.total-count {
    color: #000;
}

.border-t {
    padding-top: 15px;
    margin-top: 15px;
    border: 0 solid #dae1e7;
    border-top-width: 1px;
}
</style>
