<template>
  <div class="uk-overflow-auto">
    <table class="uk-table uk-table-divider">
      <thead>
        <tr>
          <th>Action</th>
          <th>Price</th>
          <th>Volume</th>
          <th>Condition ID</th>
          <th>Time</th>
        </tr>
      </thead>
      <tbody>
        <tr class="buy-row">
          <td>Buy</td>
          <td>123 Р</td>
          <td>10</td>
          <td>11</td>
          <td>23 Dec 2021</td>
        </tr>
        <tr class="sell-row">
          <td>Sell</td>
          <td>123 Р</td>
          <td>10</td>
          <td>11</td>
          <td>23 Dec 2021</td>
        </tr>
        <tr class="buy-row">
          <td>Buy</td>
          <td>123 Р</td>
          <td>10</td>
          <td>11</td>
          <td>23 Dec 2021</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style>
.buy-row {
  background-color: #e8faac;
}

.sell-row {
  background-color: #f2c9d5;
}
</style>
