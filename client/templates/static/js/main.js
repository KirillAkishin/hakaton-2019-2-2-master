const app = new Vue({
  el: '#app',
  data: {
    // auth: !!localStorage.token,
    auth: true,
    username: '',
    pass: '',
    error: null,
    tabs: ['Открыть позицию', 'Мои позиции', 'Цены'],
    activeTab: 0,
    amount: 0,
    price: 0,
    ticker: '',
    options: null,
    requestHeaders: [
      { name: 'Тикер', key: 'ticker' },
      { name: 'Количество', key: 'amount' },
      { name: 'Тип', key: 'type' },
      { name: 'Статус', key: 'status' },
      { name: 'Отменить', key: 'cancel' },
    ],
    requests: null,
    activeTickerTab: 0,
    priceHeaders: [
      { name: 'Время', key: 'time' },
      { name: 'Открытые', key: 'open' },
      { name: 'Верхние', key: 'high' },
      { name: 'Нижние', key: 'low' },
      { name: 'Закрытые', key: 'close' },
      { name: 'Количество', key: 'amount' },
    ],
    activeInterval: null,
  },
  created() {
    axios.defaults.baseURL = '';
    this.options = ['SPFB.RTS', 'IMOEX', 'USD000UTSTOM'];
    this.requests = [
      {
        ticker: 'SPFB.RTS',
        amount: 100,
        type: 'SELL',
        status: 'частично исполнена',
      },
      {
        ticker: 'IMOEX',
        amount: 100,
        type: 'SELL',
        status: 'ожидает исполнения',
      },
    ];
  },
  methods: {
    blur(e) {
      e.target.classList.add('input-blurred');
    },
    prevent(e) {
      if (e.keyCode >= 187 && e.keyCode <= 189) {
        e.preventDefault();
      }
    },
    addToken(type) {
      axios.post('/api/v1/register', {
        username: this.username,
        pass: this.pass,
      })
        .then(() => {
          console.log(response);
          this.auth = true;
        })
        .catch((error) => {
          console.log(error.response)
        });
      axios.defaults.headers.common.Authorization = type;
      console.log(axios.defaults.headers.common);
    },
    deleteToken() {
      this.auth = false;
      delete axios.defaults.headers.common.Authorization;
      console.log(axios.defaults.headers.common);
    },
    go(type) {
      Object.keys(this.$refs).forEach((ref) => {
        this.$refs[ref].classList.add('input-blurred');
      })
      if (this.username && this.pass) {
        this.addToken(type);
      } else {
        this.error = 'Введите имя пользователя и пароль';
        setTimeout(() => { this.error = null }, 2000);
      }
    },
    logout() {
      this.deleteToken();
    },
    deal(type) {
      console.log(type);
    },
    tickers(option) {
      console.log('start', option);
      let id = setInterval(() => { console.log('update', option); }, 3000);
      setTimeout(() => { clearInterval(id); console.log('stop', id); }, 10000);
      if (option === this.options[0]) {
        return [
          {
            time: '12:01',
            open: 3,
            high: 4,
            low: 1,
            close: 2,
            amount: 100,
          },
          {
            time: '12:02',
            open: 2,
            high: 4,
            low: 1,
            close: 2,
            amount: 100,
          },
          {
            time: '12:03',
            open: 3,
            high: 4,
            low: 1,
            close: 2,
            amount: 100,
          },
        ];
      } else {
        return [];
      }
    }
  },
});