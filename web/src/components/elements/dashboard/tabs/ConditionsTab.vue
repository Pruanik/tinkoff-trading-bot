<template>
  <div class="uk-overflow-auto">
    <table class="uk-table uk-table-divider">
      <thead>
        <tr>
          <th>ID</th>
          <th>Selling price</th>
          <th>Take profit</th>
          <th>Stope lost</th>
          <th>Direct</th>
          <th>Status</th>
          <th>Price delta</th>
          <th>Setting ID</th>
          <th>Created at</th>
        </tr>
      </thead>
      <tbody>
        <tr>
          <td>123</td>
          <td>10 Р</td>
          <td>25 Р</td>
          <td>8 Р</td>
          <td>Up</td>
          <td>ACTIVE</td>
          <td>-</td>
          <td>1</td>
          <td>23 Dec 2021</td>
        </tr>
        <tr>
          <td>123</td>
          <td>10 Р</td>
          <td>25 Р</td>
          <td>8 Р</td>
          <td>Up</td>
          <td>ACTIVE</td>
          <td>-</td>
          <td>1</td>
          <td>23 Dec 2021</td>
        </tr>
        <tr>
          <td>123</td>
          <td>10 Р</td>
          <td>25 Р</td>
          <td>8 Р</td>
          <td>Up</td>
          <td>ACTIVE</td>
          <td>-</td>
          <td>1</td>
          <td>23 Dec 2021</td>
        </tr>
        <tr>
          <td>123</td>
          <td>10 Р</td>
          <td>25 Р</td>
          <td>8 Р</td>
          <td>Up</td>
          <td>ACTIVE</td>
          <td>-</td>
          <td>1</td>
          <td>23 Dec 2021</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
export default {
  mounted() {
    console.log("mounted conditions");
  },
  unmounted() {
    console.log("unmounted conditions");
  },
};
</script>
