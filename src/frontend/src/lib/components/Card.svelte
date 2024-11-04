<!-- title date format -->
<script>
  export let title;
  export let startDate;
  export let endDate;
  export let format;
  export let href;
</script>

<a href="{href}" class="card">
  <div class="title">{title}</div>
  <div class="bottom-row">
    <!-- <img src="https://img.icons8.com/ios-filled/50/000000/calendar--v1.png" alt="Calendar Icon" class="icon"> -->
    <span class="format">{startDate}</span>
    <span class="format">{endDate}</span>
    <span class="format">{format}</span>
  </div>
</a>

<style>
    .card {
      width: 200px;
      border: 1px solid #ccc;
      padding: 16px;
      margin: 16px 0;
      border-radius: 8px;
      display: flex;
      flex-direction: column;
      align-items: center;
      font-family: Arial, sans-serif;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
    }

    .title {
      font-size: 18px;
      font-weight: bold;
      margin-bottom: 16px;
    }

    .bottom-row {
      display: flex;
      justify-content: space-between;
      width: 100%;
      align-items: center;
    }

    /* .icon {
      width: 24px;
      height: 24px;
      margin-right: 8px;
    } */

    .format {
      font-size: 14px;
    }

    .card:hover {
      background-color: #f9f9f9;
    }
</style>
