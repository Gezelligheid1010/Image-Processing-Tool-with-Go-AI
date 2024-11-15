<template>
  <div class="content">
    <div class="center">
      <!-- 搜索框和按钮 -->
      <div class="search-bar">
        <input type="text" v-model="searchQuery" placeholder="搜索分组..." />
        <!-- 将图标放入一个按钮容器中 -->
        <button @click="searchCategory" class="search-icon-button">
          <i class="search-icon el-icon-search"></i>
        </button>
      </div>

      <!-- 作品展示区 -->
      <div class="work-grid">
        <!-- 新建作品按钮始终放在第一个位置 -->
        <div class="work-item new-work-item" @click="createNewWork">
          <div class="new-work">+</div>
        </div>

        <!-- 显示作品项和空占位符 -->
        <div
            v-for="(work, index) in displayedWorks"
            :key="work.category_id || `placeholder-${index}`"
            class="work-item"
            :class="{ 'placeholder-item': !work.cover_url }"
            @click="work.category_id ? viewCategoryDetails(work) : null"
        >
          <img v-if="work.cover_url" :src="work.cover_url" alt="封面图" class="work-cover" />
          <div v-if="work.category_name" class="work-title">{{ work.category_name }}</div>
        </div>
      </div>

      <!-- 分页控件 -->
      <div class="pagination-block">
        <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page="pageNumber"
            :page-sizes="[6, 12, 18, 24]"
            :page-size="pageSize"
            layout="total, sizes, prev, pager, next, jumper"
            :total="filteredWorks.length">
        </el-pagination>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "Category",
  data() {
    return {
      works: [],
      filteredWorks: [],
      searchQuery: "",
      pageNumber: 1,
      pageSize: 6,
    };
  },
  computed: {
    paginatedWorks() {
      const start = (this.pageNumber - 1) * this.pageSize;
      return this.filteredWorks.slice(start, start + this.pageSize);
    },
    displayedWorks() {
      const works = this.paginatedWorks;
      const emptySlots = Math.max(0, 8 - works.length);
      return works.concat(Array.from({ length: emptySlots }, (_, index) => ({ placeholderId: index })));
    }
  },
  methods: {
    async fetchWorks() {

      try {
        const response = await this.$axios({
          method: "get",
          url: "/category",
          params: {
            page: this.pageNumber,
            size: this.pageSize,
            // 其他必要的查询参数可以在此添加
          }
        });

        console.log("response:", response)
        if (response.code === 1000) {
          this.works = response.data;  // 将数据赋值给 `works`
          this.filteredWorks = this.works;  // 初始化过滤后的数据
          // this.pageTotal = response.data.length % this.pageSize;  // 设置分页总数
        } else {
          console.log("获取作品数据失败1:", response.message);
        }
      } catch (error) {
        console.error("获取作品数据失败2:", error);
      }
    },

    filterWorks() {
      // console.log("this.filteredWorks:",this.filteredWorks)
      // console.log("this.searchQuery:",this.searchQuery.toLowerCase())
      this.filteredWorks = this.works.filter(work =>
          work.category_name.includes(this.searchQuery)
      );
      this.pageNumber = 1; // 重置到第一页
    },
    createNewWork() {
      this.$router.push({ name: "NewCategory" });
    },
    viewCategoryDetails(work) {
      console.log("查看作品详情", work);
      this.$router.push({ name: "CategoryDetail" , params: { category_id: work.category_id }});
    },
    handleCurrentChange(val) {
      this.pageNumber = val;
    },
    handleSizeChange(val) {
      this.pageSize = val;
    },

    searchCategory() {
      this.filterWorks();
    },


  },
  mounted() {
    this.fetchWorks();
  }
};
</script>
<!--<style scoped>-->

<style scoped lang="less">
.content {
  max-width: 100%;
  min-height: 600px;
  box-sizing: border-box;
  display: flex;
  flex-direction: row;
  justify-content: center;
  margin: 0 auto;
  padding: 20px 24px;
  margin-top: 48px;

  .center {
    width: 640px;
    padding-bottom: 10px;
    margin: 0 24px;

    .pagination-block {
      background: #fff;
      padding: 8px;
    }
  }
}


.search-bar {
  display: flex;
  align-items: center;
  margin-bottom: 20px;

  input[type="text"] {
    padding: 10px;
    font-size: 16px;
    width: 100%;
    border-radius: 4px 0 0 4px; /* 左侧圆角 */
    border: 1px solid #ccc;
    box-sizing: border-box;
  }

  .search-icon-button {
    background-color: #007bff;
    border: none;
    padding: 10px;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
    border-radius: 0 4px 4px 0; /* 右侧圆角 */
    transition: background-color 0.2s, transform 0.1s;
    margin-left: -1px; /* 去除按钮和输入框之间的空隙 */
  }

  .search-icon {
    font-size: 20px;
    color: white;
  }

  .search-icon-button:hover {
    background-color: #0056b3;
  }

  .search-icon-button:active {
    background-color: #004085;
    transform: scale(0.95);
  }
}

.work-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(150px, 1fr));
  gap: 16px;
}
.work-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;

  .work-cover {
    width: 100%;
    aspect-ratio: 1 / 1; /* 强制为正方形 */
    object-fit: cover;
    border-radius: 8px;
  }

  .work-title {
    margin-top: 8px;
    font-size: 14px;
    font-weight: bold;
    text-align: center;
    color: #333;
  }

}


//.new-work-item {
//  border: none; /* 移除新建作品按钮的边框 */
//}

.new-work-item {
  border: 2px dashed #aaa; /* 虚线边框 */
  display: flex;
  justify-content: center;
  align-items: center;
  width: 150px; /* 设置宽度以确保为正方形 */
  height: 150px; /* 设置高度以确保为正方形 */
  cursor: pointer;
  transition: border-color 0.2s;
}

.new-work-item:hover {
  border-color: #333; /* 悬停效果 */
}

.new-work {
  font-size: 36px;
  color: #aaa;
  line-height: 1; /* 确保加号垂直居中 */
}

//.work-cover {
//  width: 100%;
//  height: 100%;
//  object-fit: cover;
//  border-radius: 8px;
//}

.placeholder-item {
  border: none; /* 移除其他空位的虚线 */
  background-color: #f0f0f0;
  border-radius: 8px;
}

.pagination-block {
  margin-top: 20px;
  display: flex;
  justify-content: center;
}
</style>
