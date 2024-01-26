<template>
  <div class="admin-users">

   <table>
     <thead>
     <tr>
       <th>аватар</th>
       <th>имя</th>
       <th>тип аккаунта</th>
       <th>дата регистрации</th>
       <th>действия</th>
     </tr>
     </thead>

     <tbody>
     <tr v-for="user in list">

       <td>
         <img :src="user.image" alt="">
       </td>
       <td>{{user.first_name}} {{user.last_name}}</td>
       <td>{{user.type}}</td>
       <td>{{user.created_at}}</td>
       <td>
         -
       </td>
     </tr>
     </tbody>
   </table>
  </div>
</template>

<script>

import {$adminApi} from "@/api/api";

export default {
  name: "UsersComponent",
  data(){

    return {
      list: []
    }
  },

  created() {
    this.loadUsers()
  },
  methods: {
    async loadUsers(){

      const resp = await $adminApi.get("/admin/api/users")

      this.list = resp.data

      console.log(resp.data)

    }
  }
}
</script>

<style>

.admin-users td img {
  width: 40px;
  border-radius: 20px;
}
</style>