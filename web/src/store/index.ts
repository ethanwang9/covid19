import useAppStore from "./modules/app";
import useAdminStore from "./modules/admin";

const useStore = () => ({
    app: useAppStore(),
    admin: useAdminStore(),
});

export default useStore;