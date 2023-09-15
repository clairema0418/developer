import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import testData from '../config/testData.json';

export const fetchProductData = createAsyncThunk('api/fetchProductData', async () => {
  return testData.productData;
});

export const fetchCustomerReviews = createAsyncThunk('api/fetchCustomerReviews', async () => {
  return testData.customerReviewData;
});

export const fetchCaseStudies = createAsyncThunk('api/fetchCaseStudies', async () => {
  return testData.caseStudyData;
});

export const fetchNewsUpdates = createAsyncThunk('api/fetchNewsUpdates', async () => {
  return testData.newsUpdateData;
});

export const fetchCompanyEvents = createAsyncThunk('api/fetchCompanyEvents', async () => {
  return testData.companyEventData;
});

const apiSlice = createSlice({
  name: 'api',
  initialState: {
    productData: [],
    customerReviews: [],
    caseStudies: [],
    newsUpdates: [],
    companyEvents: [],
    status: 'idle',
    error: null,
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchProductData.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchProductData.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.productData = action.payload;
      })
      .addCase(fetchProductData.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      })
      .addCase(fetchCustomerReviews.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCustomerReviews.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.customerReviews = action.payload;
      })
      .addCase(fetchCustomerReviews.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      })
      .addCase(fetchCaseStudies.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCaseStudies.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.caseStudies = action.payload;
      })
      .addCase(fetchCaseStudies.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      })
      .addCase(fetchNewsUpdates.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchNewsUpdates.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.newsUpdates = action.payload;
      })
      .addCase(fetchNewsUpdates.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      })
      .addCase(fetchCompanyEvents.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchCompanyEvents.fulfilled, (state, action) => {
        state.status = 'succeeded';
        state.companyEvents = action.payload;
      })
      .addCase(fetchCompanyEvents.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.error.message;
      });
  },
});

export default apiSlice.reducer;